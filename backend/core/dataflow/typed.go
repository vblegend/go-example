package dataflow

type IHandleNode interface {
	Init(id string, name string, typed string, handler DataHandler) IHandleNode
	Next(origin FlowData, current FlowData)
	Inputs() INodeCollection
	Outputs() INodeCollection
	GetName() string
	SetName(string)
	GetId() string
	SetId(string)
}

type FlowData interface{}
type DataHandler func(origin FlowData, current FlowData) FlowData

type IFlowEngine interface {
	AppendNode(node IHandleNode)
	InputData()
}

type INodeCollection interface {
	Add(node IHandleNode)
	Remove(node IHandleNode)
	Clear()
	Contains(node IHandleNode) bool
	Count() int
	IndexOf(node IHandleNode) int
	Nodes() []IHandleNode
}
