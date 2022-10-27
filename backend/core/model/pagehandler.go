package model

import (
	"backend/core/restful"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 列表查询处理器
func ListHander(resultTyped IModel) gin.HandlerFunc {
	t := reflect.ValueOf(resultTyped)
	sliceType := reflect.SliceOf(t.Type())
	makeModel := func() interface{} {
		emptySlice := reflect.MakeSlice(sliceType, 0, 0)
		return emptySlice.Interface()
	}
	return func(c *gin.Context) {
		tx := c.MustGet("db").(*gorm.DB)
		sclice := makeModel()
		var count int64
		err := tx.Table(resultTyped.TableName()).Find(sclice).Count(&count).Error
		if err != nil {
			restful.Error(c, http.StatusInternalServerError, err, "")
		} else {
			restful.OK(c, sclice, "OK")
		}
	}
}

// 页查询处理器
func PageHander(query IPagination, resultTyped IModel) gin.HandlerFunc {
	t := reflect.TypeOf(resultTyped)
	sliceType := reflect.SliceOf(t)
	makeModel := func() interface{} {
		return reflect.MakeSlice(sliceType, 0, 0).Interface()
	}
	return func(c *gin.Context) {
		tx := c.MustGet("db").(*gorm.DB)
		sclice := makeModel()
		var count int64
		err := tx.Table(resultTyped.TableName()).Find(sclice).Count(&count).Error
		if err != nil {
			restful.Error(c, http.StatusInternalServerError, err, "")
		} else {
			restful.PageOK(c, sclice, int(count), query.GetPageIndex(), query.GetPageSize(), "OK")
		}
	}
}

// 对象创建处理器
func CreateHander(resultTyped IModel, succeedCallback func(object interface{})) gin.HandlerFunc {
	t := reflect.TypeOf(resultTyped)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	makeModel := func() interface{} {
		return reflect.New(t).Interface()
	}
	return func(c *gin.Context) {
		tx := c.MustGet("db").(*gorm.DB)
		model := makeModel()
		err := c.Bind(model)
		if err != nil {
			restful.Error(c, http.StatusBadRequest, err, "")
			return
		}
		err = tx.Table(resultTyped.TableName()).Create(model).Error
		if err != nil {
			restful.Error(c, http.StatusInternalServerError, err, "")
		} else {
			if succeedCallback != nil {
				succeedCallback(model)
			}
			restful.OK(c, model, "OK")
		}
	}
}

// 对象更新处理器
func UpdateHander(resultTyped IModel, succeedCallback func(object interface{})) gin.HandlerFunc {
	t := reflect.TypeOf(resultTyped)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	makeModel := func() interface{} {
		return reflect.New(t).Interface()
	}
	return func(c *gin.Context) {
		tx := c.MustGet("db").(*gorm.DB)
		model := makeModel()
		err := c.Bind(model)
		if err != nil {
			restful.Error(c, http.StatusBadRequest, err, "")
			return
		}
		err = tx.Table(resultTyped.TableName()).Updates(model).Error
		if err != nil {
			restful.Error(c, http.StatusInternalServerError, err, "")
		} else {
			if succeedCallback != nil {
				succeedCallback(model)
			}
			restful.OK(c, model, "OK")
		}
	}
}

// 对象删除处理器
func DeleteHander(query IIdentity, resultTyped IModel, succeedCallback func(id interface{})) gin.HandlerFunc {
	return func(c *gin.Context) {
		tx := c.MustGet("db").(*gorm.DB)
		err := tx.Table(resultTyped.TableName()).Delete(query.GetId()).Error
		if err != nil {
			restful.Error(c, http.StatusInternalServerError, err, "")
		} else {
			if succeedCallback != nil {
				succeedCallback(query.GetId())
			}
			restful.OK(c, gin.H{}, "OK")
		}
	}
}

// ID查询处理器
func IndexHander(query IIdentity, resultTyped IModel) gin.HandlerFunc {
	t := reflect.TypeOf(resultTyped)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	makeModel := func() interface{} {
		return reflect.New(t).Interface()
	}
	return func(c *gin.Context) {
		tx := c.MustGet("db").(*gorm.DB)
		model := makeModel()
		err := c.Bind(query)
		if err != nil {
			restful.Error(c, http.StatusBadRequest, err, "")
			return
		}
		err = tx.Table(resultTyped.TableName()).Where(query.GetId()).First(&model).Error
		if err != nil {
			restful.Error(c, http.StatusInternalServerError, err, "")
		} else {
			restful.OK(c, model, "OK")
		}
	}
}

// 多条件查询处理器
func SearchHander(query IPagination, resultTyped IModel) gin.HandlerFunc {
	t := reflect.ValueOf(resultTyped)
	sliceType := reflect.SliceOf(t.Type())
	makeModel := func() interface{} {
		emptySlice := reflect.MakeSlice(sliceType, 0, 0)
		return emptySlice.Interface()
	}
	return func(c *gin.Context) {
		tx := c.MustGet("db").(*gorm.DB)
		sclice := makeModel()
		var count int64
		err := tx.Table(resultTyped.TableName()).Find(sclice).Count(&count).Error
		if err != nil {
			restful.Error(c, http.StatusInternalServerError, err, "")
		} else {
			restful.PageOK(c, sclice, int(count), query.GetPageIndex(), query.GetPageSize(), "OK")
		}
	}
}

// 对象创建处理器
func IdHander(query IIdentity, succeedCallback func(object interface{})) gin.HandlerFunc {
	t := reflect.TypeOf(query)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	makeModel := func() interface{} {
		return reflect.New(t).Interface()
	}
	return func(c *gin.Context) {
		model := makeModel()
		err := c.Bind(model)
		if err != nil {
			restful.Error(c, http.StatusBadRequest, err, "")
			return
		}
		if succeedCallback != nil {
			succeedCallback(model)
		}
	}
}
