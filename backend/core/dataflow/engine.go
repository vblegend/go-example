package dataflow

type FlowEngine struct {
	nodes []IHandleNode
}

func (e *FlowEngine) AppendNode(node IHandleNode) {

}

func (e *FlowEngine) InputData() {

}

func NewEngine() IFlowEngine {
	return &FlowEngine{}
}
