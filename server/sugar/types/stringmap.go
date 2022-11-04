package types

// SMap 一个 string MAP 对象
type SMap[T string | int | uint | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64 | interface{}] map[string]T

// SIntMap 是一个 map[string]int 对象
type SIntMap = SMap[int]

// SUIntMap 是一个 map[string]uint 对象
type SUIntMap = SMap[uint]

// SSMap 是一个 map[string]string 对象
type SSMap = SMap[string]

func (m SMap[T]) Put(key string, value T) {
	m[key] = value
}

func (m SMap[T]) Get(key string) T {
	return m[key]
}

func (m SMap[T]) Exist(key string) bool {
	_, ok := m[key]
	return ok
}

func (m SMap[T]) Clear() {
	for k := range m {
		delete(m, k)
	}
}

func (m SMap[T]) Keys() Strings {
	keys := Strings{}
	for k := range m {
		keys.Append(k)
	}
	return keys
}

func (m SMap[T]) SortKeys() []string {
	keys := Strings{}
	for k := range m {
		keys.Append(k)
	}
	keys.Sort()
	return keys
}
