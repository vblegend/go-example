package restful

type Response struct {
	// 数据集
	TraceId string      `json:"traceId,omitempty"`
	Code    int         `json:"code"`
	Msg     string      `json:"msg,omitempty"`
	Data    interface{} `json:"data"`
}

type Page struct {
	Count     int         `json:"count"`
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
	List      interface{} `json:"list"`
}
