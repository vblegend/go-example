package types

import (
	"fmt"
	"reflect"
	"sort"
)

type TInt interface {
	int | int8 | int16 | int32 | int64
}

type TUint interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type TFloat interface {
	float32 | float64
}

type TNumber interface {
	TInt | TUint | TFloat
}

// Slice string 切片数组
type Slice[T comparable] []T

// Strings string 切片数组
type Strings = Slice[string]

// Ints int 切片数组
type Ints = Slice[int]

// UInts uint 切片数组
type UInts = Slice[uint]

// Len 返回数组长度
func (s Slice[T]) Len() int {
	return len(s)
}

// IndexOf 返回数组中字符串的位置，未找到返回-1
func (s Slice[T]) IndexOf(str T) int {
	for i := 0; i < len(s); i++ {
		if s[i] == str {
			return i
		}
	}
	return -1
}

// Append 追加一条数据
func (s *Slice[T]) Append(str T) {
	a := append(*s, str)
	*s = a
}

// Clear 清空数组内容
func (s *Slice[T]) Clear() {
	*s = Slice[T]{}
}

// Sort 对数组进行正序排序
func (s *Slice[T]) Sort() {
	sort.Sort(s)
}

func (s *Slice[T]) Less(i, j int) bool {
	origin := *s
	return s.dayu(origin[i], origin[j])
}

func (s *Slice[T]) dayu(i, j any) bool {
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.String {
		return i.(string) < j.(string)
	}
	if v.Kind() >= reflect.Int && v.Kind() <= reflect.Int64 {
		return i.(int) < j.(int)
	}
	if v.Kind() >= reflect.Uint && v.Kind() <= reflect.Uint64 {
		return i.(uint) < j.(uint)
	}
	if v.Kind() == reflect.Float32 || v.Kind() == reflect.Float64 {
		return i.(float64) < j.(float64)
	}
	var ii = fmt.Sprintf("%v", i)
	var jj = fmt.Sprintf("%v", j)
	return ii < jj
}

// Reverse 翻转数组内容
func (s *Slice[T]) Reverse() {
	origin := *s
	blen := len(origin)
	for i := 0; i < blen/2; i++ {
		temp := origin[blen-1-i]
		origin[blen-1-i] = origin[i]
		origin[i] = temp
	}
	*s = origin
}

func (s *Slice[T]) Swap(i, j int) {
	origin := *s
	origin[i], origin[j] = origin[j], origin[i]
	*s = origin
}

// Push 在数组开头放入一条数据
func (s *Slice[T]) Push(str T) {
	a := append(Slice[T]{str}, *s...)
	*s = a
}

// Pop 从数组开头取出一个元素
func (s *Slice[T]) Pop() (r T) {
	if s.Len() == 0 {
		panic("the index range is out of bounds")
	}
	origin := *s
	r = origin[0]
	*s = origin[1:]
	return
}

// Take 从数组取出一条数据
func (s *Slice[T]) Take(index int) T {
	if index < 0 || index >= s.Len() {
		panic("the index range is out of bounds")
	}
	origin := *s
	res := origin[index]
	*s = append(origin[:index], origin[index+1:]...)
	return res
}

// RemoveAt 移除一条指定位置的元素
func (s *Slice[T]) RemoveAt(index int) {
	if index < 0 || index >= s.Len() {
		panic("the index range is out of bounds")
	}
	origin := *s
	*s = append(origin[:index], origin[index+1:]...)
}

// Remove 移除一条数据
func (s *Slice[T]) Remove(str T) {
	origin := *s
	for i := 0; i < len(origin); i++ {
		if origin[i] == str {
			origin.RemoveAt(i)
			return
		}
	}
}

// RemoveAll 移除所有等于str数据
func (s *Slice[T]) RemoveAll(str T) {
	origin := *s
	for i := 0; i < len(origin); i++ {
		if origin[i] == str {
			origin.RemoveAt(i)
			i--
		}
	}
}
