package restful

import (
	"net/http"
	"reflect"
	"server/sugar/model"
	"server/sugar/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 列表查询处理器
func ListHander(resultTyped model.IModel) gin.HandlerFunc {
	makeSclice := utils.MakeSliceFunc(resultTyped)
	return func(c *gin.Context) {
		tx := c.MustGet("db").(*gorm.DB).WithContext(c)
		sclice := makeSclice()
		err := tx.Table(resultTyped.TableName()).Find(&sclice).Error
		if err != nil {
			Error(c, http.StatusInternalServerError, err)
		} else {
			OK(c, sclice, "OK")
		}
	}
}

// WherePageHander 页查询处理器  返回查询的指定页的数据
// queryModel 查询参数模型（为nil时不做where过滤），可重用 resultModel
// pageModel 分页参数模型
// resultModel 返回类型模型
func WherePageHander(queryModel interface{}, pageModel model.IPagination, resultModel model.IModel) gin.HandlerFunc {
	var makeQuery func() interface{}
	makeModel := utils.MakeModelFunc(pageModel)
	makeSclice := utils.MakeSliceFunc(resultModel)
	if queryModel != nil {
		makeQuery = utils.MakeModelFunc(queryModel)
	}
	return func(c *gin.Context) {
		tx := c.MustGet("db").(*gorm.DB).WithContext(c)
		params := makeModel().(model.IPagination)
		err := c.Bind(params)
		if err != nil {
			Error(c, http.StatusBadRequest, err)
			return
		}
		sclice := makeSclice()
		var count int64
		offset := params.GetPageIndex()*params.GetPageSize() - params.GetPageSize()
		tx = tx.Table(resultModel.TableName())
		if makeQuery != nil {
			query := makeQuery()
			err = c.Bind(query)
			if err != nil {
				Error(c, http.StatusBadRequest, err)
				return
			}
			tx = tx.Where(query)
		}
		err = tx.Count(&count).Offset(offset).Limit(params.GetPageSize()).Find(&sclice).Error
		if err != nil {
			Error(c, http.StatusInternalServerError, err)
		} else {
			PageOK(c, sclice, int(count), params.GetPageIndex(), params.GetPageSize(), "OK")
		}
	}
}

// CreateHander 对象创建处理器，从接口读取 typeModel 并写入库，成功后调用回调函数succeedCallback
func CreateHander(typeModel model.IModel, succeedCallback func(object interface{})) gin.HandlerFunc {
	makeModel := utils.MakeModelFunc(typeModel)
	return func(c *gin.Context) {
		tx := c.MustGet("db").(*gorm.DB).WithContext(c)
		model := makeModel()
		err := c.Bind(model)
		if err != nil {
			Error(c, http.StatusBadRequest, err)
			return
		}
		err = tx.Table(typeModel.TableName()).Create(model).Error
		if err != nil {
			Error(c, http.StatusInternalServerError, err)
		} else {
			if succeedCallback != nil {
				succeedCallback(reflect.ValueOf(model).Elem().Interface())
			}
			OK(c, model, "OK")
		}
	}
}

// UpdateHander 对象更新处理器，从接口读取 typeModel 并写入库，成功后调用回调函数succeedCallback
func UpdateHander(typeModel model.IModel, succeedCallback func(object interface{})) gin.HandlerFunc {
	t := reflect.TypeOf(typeModel)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	makeModel := func() interface{} {
		return reflect.New(t).Interface()
	}
	return func(c *gin.Context) {
		tx := c.MustGet("db").(*gorm.DB).WithContext(c)
		model := makeModel()
		err := c.Bind(model)
		if err != nil {
			Error(c, http.StatusBadRequest, err)
			return
		}
		err = tx.Table(typeModel.TableName()).Updates(model).Error
		if err != nil {
			Error(c, http.StatusInternalServerError, err)
		} else {
			if succeedCallback != nil {
				succeedCallback(model)
			}
			OK(c, model, "OK")
		}
	}
}

// DeleteHander 对象删除处理器，从tableModel 表内删除符合条件的记录，成功后调用回调函数succeedCallback
func DeleteHander(queryModel interface{}, tableModel model.IModel, succeedCallback func(queryObject interface{})) gin.HandlerFunc {
	makeQuery := utils.MakeModelFunc(queryModel)
	return func(c *gin.Context) {
		query := makeQuery()
		tx := c.MustGet("db").(*gorm.DB).WithContext(c)
		err := AutoBind(c, query)
		if err != nil {
			Error(c, http.StatusBadRequest, err)
			return
		}
		err = tx.Table(tableModel.TableName()).Delete(query).Error
		if err != nil {
			Error(c, http.StatusInternalServerError, err)
		} else {
			if succeedCallback != nil {
				succeedCallback(query)
			}
			OK(c, gin.H{}, "OK")
		}
	}
}

// WhereFirstHander 单例查询处理器  返回 符合 queryModel 的第一条记录
func WhereFirstHander(queryModel interface{}, resultModel model.IModel) gin.HandlerFunc {
	makeModel := utils.MakeModelFunc(resultModel)
	makeQuery := utils.MakeModelFunc(queryModel)
	return func(c *gin.Context) {
		tx := c.MustGet("db").(*gorm.DB).WithContext(c)
		model := makeModel()
		query := makeQuery()
		err := AutoBind(c, query)
		if err != nil {
			Error(c, http.StatusBadRequest, err)
			return
		}
		err = tx.Table(resultModel.TableName()).Where(query).First(&model).Error
		if err != nil {
			Error(c, http.StatusInternalServerError, err)
		} else {
			OK(c, model, "OK")
		}
	}
}

// WhereListHander 列表件查询处理器，根据queryModel 查询指定resultModel对象列表
func WhereListHander(queryModel interface{}, resultModel model.IModel) gin.HandlerFunc {
	makeSclice := utils.MakeSliceFunc(resultModel)
	makeQuery := utils.MakeModelFunc(queryModel)
	return func(c *gin.Context) {
		tx := c.MustGet("db").(*gorm.DB)
		query := makeQuery()
		err := AutoBind(c, query)
		if err != nil {
			Error(c, http.StatusBadRequest, err)
			return
		}
		sclice := makeSclice()
		err = tx.Table(resultModel.TableName()).Where(query).Find(&sclice).Error
		if err != nil {
			Error(c, http.StatusInternalServerError, err)
		} else {
			OK(c, sclice, "OK")
		}
	}
}

// ActionHander 动作处理器
func ActionHander(queryObject interface{}, succeedCallback func(object interface{})) gin.HandlerFunc {
	makeQuery := utils.MakeModelFunc(queryObject)
	return func(c *gin.Context) {
		model := makeQuery()
		err := AutoBind(c, model)
		if err != nil {
			Error(c, http.StatusBadRequest, err)
			return
		}
		if succeedCallback != nil {
			succeedCallback(model)
		}
	}
}
