package utils

import "reflect"

// CreateObjectFunc 创建对象的方法类型
type CreateObjectFunc func() interface{}

// MakeSliceFunc 提供一个对象，返回一个创建该对象切片的方法
func MakeSliceFunc(obj interface{}) CreateObjectFunc {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	sliceType := reflect.SliceOf(t)
	return func() interface{} {
		return reflect.MakeSlice(sliceType, 0, 0).Interface()
	}
}

// MakeModelFunc 提供一个对象，返回一个创建该对象的方法
func MakeModelFunc(obj interface{}) CreateObjectFunc {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return func() interface{} {
		return reflect.New(t).Interface()
	}
}
