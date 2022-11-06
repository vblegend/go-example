package restful

import (
	"fmt"
	"net/http"
	"reflect"
	"server/sugar/model"
	"server/sugar/utils"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HandlerActionCallBack 执行处理函数，model为数据库模型对象
type HandlerActionCallBack func(model model.IModel) error

// HandlerExeFunc 一个执行方法处理器
type HandlerExeFunc func(call HandlerActionCallBack, api *Api)

// HandlerQueryCallBack 查询参数的调用函数，query 为 查询参数，model为数据库模型对象
type HandlerQueryCallBack func(query interface{}, model model.IModel) error

// HandlerQueryExecFunc 一个查询并执行的方法处理器
type HandlerQueryExecFunc func(call HandlerQueryCallBack, api *Api)

func execQueryAfter(origin interface{}) {
	switch reflect.TypeOf(origin).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(origin)
		for i := 0; i < s.Len(); i++ {
			lpModel := s.Index(i).Addr().Interface()
			if ia, ok := lpModel.(model.IModelQueryAfter); ok {
				ia.OnQueryAfter()
			}
		}
	default:
		if _, ok := origin.(model.IModel); ok {
			if qa, ok := origin.(model.IModelQueryAfter); ok {
				qa.OnQueryAfter()
			}
		} else {
			panic("invalid data type with model.IModel")
		}
	}
}

// ListHander 列表查询处理器
func ListHander(resultTyped model.IModel) gin.HandlerFunc {
	makeSclice := utils.MakeSliceFunc(resultTyped)
	return func(c *gin.Context) {
		tx := c.MustGet("db").(*gorm.DB).WithContext(c)
		sclice := makeSclice()
		err := tx.Table(resultTyped.TableName()).Find(&sclice).Error
		if err != nil {
			Error(c, http.StatusInternalServerError, err)
		} else {
			fs := time.Now()
			execQueryAfter(sclice)
			l := time.Now().Sub(fs)
			fmt.Println(l)
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
			execQueryAfter(sclice)
			PageOK(c, sclice, int(count), params.GetPageIndex(), params.GetPageSize(), "OK")
		}
	}
}

// CreateHander 对象创建处理器，从接口读取 typeModel 并写入库，成功后调用回调函数succeedCallback
func CreateHander(handler HandlerExeFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		api := new(Api).MakeContext(c).MakeOrm()
		callback := func(lpModel model.IModel) error {
			tx := c.MustGet("db").(*gorm.DB).WithContext(c)
			err := c.Bind(lpModel)
			if err != nil {
				Error(c, http.StatusBadRequest, err)
				return err
			}
			if ib, ok := lpModel.(model.IModelInsertBefore); ok {
				if err = ib.OnInsertBefore(); err != nil {
					Error(c, http.StatusBadRequest, err)
					return err
				}
			}
			err = tx.Table(lpModel.TableName()).Create(lpModel).Error
			if ia, ok := lpModel.(model.IModelInsertAfter); ok {
				ia.OnInsertAfter()
			}
			if err != nil {
				Error(c, http.StatusInternalServerError, err)
				return err
			} else {
				OK(c, lpModel, "OK")
				return nil
			}
		}
		handler(callback, api)
	}
}

// UpdateHander 对象更新处理器，从接口读取 typeModel 并写入库，成功后调用回调函数succeedCallback
func UpdateHander(handler HandlerExeFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		api := new(Api).MakeContext(c).MakeOrm()
		callback := func(lpModel model.IModel) error {
			tx := c.MustGet("db").(*gorm.DB).WithContext(c)
			err := c.Bind(lpModel)
			if err != nil {
				Error(c, http.StatusBadRequest, err)
				return err
			}
			if ub, ok := lpModel.(model.IModelUpdateBefore); ok {
				if err = ub.OnUpdateBefore(); err != nil {
					Error(c, http.StatusBadRequest, err)
					return err
				}
			}
			err = tx.Table(lpModel.TableName()).Updates(lpModel).Error
			if ua, ok := lpModel.(model.IModelUpdateAfter); ok {
				ua.OnUpdateAfter()
			}
			if err != nil {
				Error(c, http.StatusInternalServerError, err)
				return err
			} else {
				OK(c, lpModel, "OK")
				return nil
			}
		}
		handler(callback, api)
	}
}

// DeleteHander 对象删除处理器，从tableModel 表内删除符合条件的记录，成功后调用回调函数succeedCallback
func DeleteHander(handler HandlerQueryExecFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		api := new(Api).MakeContext(c).MakeOrm()
		callback := func(query interface{}, lpModel model.IModel) error {
			tx := c.MustGet("db").(*gorm.DB).WithContext(c)
			err := AutoBind(c, query)
			if err != nil {
				Error(c, http.StatusBadRequest, err)
				return err
			}
			if db, ok := lpModel.(model.IModelDeleteBefore); ok {
				if err = db.OnDeleteBefore(); err != nil {
					Error(c, http.StatusBadRequest, err)
					return err
				}
			}
			err = tx.Table(lpModel.TableName()).Delete(query).Error
			if da, ok := lpModel.(model.IModelDeleteAfter); ok {
				da.OnDeleteAfter()
			}
			if err != nil {
				Error(c, http.StatusInternalServerError, err)
				return err
			} else {
				OK(c, gin.H{}, "OK")
				return nil
			}
		}
		handler(callback, api)
	}
}

// WhereFirstHander 单例查询处理器  返回 符合 queryModel 的第一条记录
func WhereFirstHander(queryModel interface{}, resultModel model.IModel) gin.HandlerFunc {
	makeModel := utils.MakeModelFunc(resultModel)
	makeQuery := utils.MakeModelFunc(queryModel)
	return func(c *gin.Context) {
		tx := c.MustGet("db").(*gorm.DB).WithContext(c)
		lpModel := makeModel()
		query := makeQuery()
		err := AutoBind(c, query)
		if err != nil {
			Error(c, http.StatusBadRequest, err)
			return
		}
		err = tx.Table(resultModel.TableName()).Where(query).First(&lpModel).Error
		if err != nil {
			Error(c, http.StatusInternalServerError, err)
		} else {
			if da, ok := lpModel.(model.IModelQueryAfter); ok {
				da.OnQueryAfter()
			}
			OK(c, lpModel, "OK")
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
			execQueryAfter(sclice)
			OK(c, sclice, "OK")
		}
	}
}

// ActionHander 动作处理器
func ActionHander(handler HandlerExeFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		api := new(Api).MakeContext(c).MakeOrm()
		callback := func(lpModel model.IModel) error {
			err := AutoBind(c, lpModel)
			if err != nil {
				Error(c, http.StatusBadRequest, err)
				return err
			}
			OK(c, gin.H{}, "OK")
			return nil
		}
		handler(callback, api)
	}
}
