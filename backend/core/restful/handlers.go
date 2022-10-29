package restful

import (
	"backend/core/model"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MakeSliceFunc(obj interface{}) func() interface{} {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	sliceType := reflect.SliceOf(t)
	return func() interface{} {
		return reflect.MakeSlice(sliceType, 0, 0).Interface()
	}
}

func MakeModelFunc(obj interface{}) func() interface{} {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return func() interface{} {
		return reflect.New(t).Interface()
	}
}

// 列表查询处理器
func ListHander(resultTyped model.IModel) gin.HandlerFunc {
	makeSclice := MakeSliceFunc(resultTyped)
	return func(c *gin.Context) {
		tx := c.MustGet("db").(*gorm.DB).WithContext(c)
		sclice := makeSclice()
		err := tx.Table(resultTyped.TableName()).Find(&sclice).Error
		if err != nil {
			Error(c, http.StatusInternalServerError, err, "")
		} else {
			OK(c, sclice, "OK")
		}
	}
}

// 页查询处理器
func WherePageHander(queryObject interface{}, page model.IPagination, resultTyped model.IModel) gin.HandlerFunc {
	var makeQuery func() interface{}
	makeModel := MakeModelFunc(page)
	makeSclice := MakeSliceFunc(resultTyped)
	if queryObject != nil {
		makeQuery = MakeModelFunc(queryObject)
	}
	return func(c *gin.Context) {
		tx := c.MustGet("db").(*gorm.DB).WithContext(c)
		params := makeModel().(model.IPagination)
		err := c.Bind(params)
		if err != nil {
			Error(c, http.StatusBadRequest, err, "")
			return
		}
		sclice := makeSclice()
		var count int64
		offset := params.GetPageIndex()*params.GetPageSize() - params.GetPageSize()
		tx = tx.Table(resultTyped.TableName())
		if makeQuery != nil {
			query := makeQuery()
			err = c.Bind(query)
			if err != nil {
				Error(c, http.StatusBadRequest, err, "")
				return
			}
			tx = tx.Where(query)
		}
		err = tx.Count(&count).Offset(offset).Limit(params.GetPageSize()).Find(&sclice).Error
		if err != nil {
			Error(c, http.StatusInternalServerError, err, "")
		} else {
			PageOK(c, sclice, int(count), params.GetPageIndex(), params.GetPageSize(), "OK")
		}
	}
}

// 对象创建处理器
func CreateHander(resultTyped model.IModel, succeedCallback func(object interface{})) gin.HandlerFunc {
	makeModel := MakeModelFunc(resultTyped)
	return func(c *gin.Context) {
		tx := c.MustGet("db").(*gorm.DB).WithContext(c)
		model := makeModel()
		err := c.Bind(model)
		if err != nil {
			Error(c, http.StatusBadRequest, err, "")
			return
		}
		err = tx.Table(resultTyped.TableName()).Create(model).Error
		if err != nil {
			Error(c, http.StatusInternalServerError, err, "")
		} else {
			if succeedCallback != nil {
				succeedCallback(reflect.ValueOf(model).Elem().Interface())
			}
			OK(c, model, "OK")
		}
	}
}

// 对象更新处理器
func UpdateHander(resultTyped model.IModel, succeedCallback func(object interface{})) gin.HandlerFunc {
	t := reflect.TypeOf(resultTyped)
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
			Error(c, http.StatusBadRequest, err, "")
			return
		}
		err = tx.Table(resultTyped.TableName()).Updates(model).Error
		if err != nil {
			Error(c, http.StatusInternalServerError, err, "")
		} else {
			if succeedCallback != nil {
				succeedCallback(model)
			}
			OK(c, model, "OK")
		}
	}
}

// 对象删除处理器
func DeleteHander(queryObject interface{}, resultTyped model.IModel, succeedCallback func(queryObject interface{})) gin.HandlerFunc {
	makeQuery := MakeModelFunc(queryObject)
	return func(c *gin.Context) {
		query := makeQuery()
		tx := c.MustGet("db").(*gorm.DB).WithContext(c)
		err := AutoBind(c, query)
		if err != nil {
			Error(c, http.StatusBadRequest, err, "")
			return
		}
		err = tx.Table(resultTyped.TableName()).Delete(query).Error
		if err != nil {
			Error(c, http.StatusInternalServerError, err, "")
		} else {
			if succeedCallback != nil {
				succeedCallback(query)
			}
			OK(c, gin.H{}, "OK")
		}
	}
}

// 单例查询处理器
func WhereFirstHander(queryObject interface{}, resultTyped model.IModel) gin.HandlerFunc {
	makeModel := MakeModelFunc(resultTyped)
	makeQuery := MakeModelFunc(queryObject)
	return func(c *gin.Context) {
		tx := c.MustGet("db").(*gorm.DB).WithContext(c)
		model := makeModel()
		query := makeQuery()
		err := AutoBind(c, query)
		if err != nil {
			Error(c, http.StatusBadRequest, err, "")
			return
		}
		err = tx.Table(resultTyped.TableName()).Where(query).First(&model).Error
		if err != nil {
			Error(c, http.StatusInternalServerError, err, "")
		} else {
			OK(c, model, "OK")
		}
	}
}

// 多条件查询处理器
func WhereListHander(queryObject interface{}, resultTyped model.IModel) gin.HandlerFunc {
	makeSclice := MakeSliceFunc(resultTyped)
	makeQuery := MakeModelFunc(queryObject)
	return func(c *gin.Context) {
		tx := c.MustGet("db").(*gorm.DB)
		query := makeQuery()
		err := AutoBind(c, query)
		if err != nil {
			Error(c, http.StatusBadRequest, err, "")
			return
		}
		sclice := makeSclice()
		err = tx.Table(resultTyped.TableName()).Where(query).Find(&sclice).Error
		if err != nil {
			Error(c, http.StatusInternalServerError, err, "")
		} else {
			OK(c, sclice, "OK")
		}
	}
}

// 对象创建处理器
func ActionHander(queryObject interface{}, succeedCallback func(object interface{})) gin.HandlerFunc {
	makeQuery := MakeModelFunc(queryObject)
	return func(c *gin.Context) {
		model := makeQuery()
		err := AutoBind(c, model)
		if err != nil {
			Error(c, http.StatusBadRequest, err, "")
			return
		}
		if succeedCallback != nil {
			succeedCallback(model)
		}
	}
}
