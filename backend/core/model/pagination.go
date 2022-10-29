package model

type IPagination interface {
	GetPageIndex() int
	GetPageSize() int
}

type IIdentity interface {
	GetId() interface{}
}
type UriIdentityInt struct {
	Id int `uri:"id"`
}

func (m *UriIdentityInt) GetId() interface{} {
	return m.Id
}

type IdentityString struct {
	Id string `form:"id"`
}

func (m *IdentityString) GetId() interface{} {
	return m.Id
}

type Pagination struct {
	PageIndex int `form:"pageIndex"`
	PageSize  int `form:"pageSize"`
}

func (m *Pagination) GetPageIndex() int {
	if m.PageIndex <= 0 {
		m.PageIndex = 1
	}
	return m.PageIndex
}

func (m *Pagination) GetPageSize() int {
	if m.PageSize <= 0 {
		m.PageSize = 10
	}
	return m.PageSize
}
