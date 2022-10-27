package dataflow

type BaseNode struct {
	name    string
	id      string
	typed   string
	handler DataHandler
	input   INodeCollection
	output  INodeCollection
}

func (n *BaseNode) Init(id string, name string, typed string, handler DataHandler) IHandleNode {
	n.id = id
	n.name = name
	n.typed = typed
	n.handler = handler
	n.input = NewNodeCollection()
	n.output = NewNodeCollection()
	return n
}

func (n *BaseNode) Next(origin FlowData, current FlowData) {
	currentData := n.handler(origin, current)
	nodes := n.Outputs().Nodes()
	for i := 0; i < len(nodes); i++ {
		nodes[i].Next(origin, currentData)
	}
}

func (n *BaseNode) Inputs() INodeCollection {
	return n.input
}

func (n *BaseNode) Outputs() INodeCollection {
	return n.output
}

func (n *BaseNode) GetName() string {
	return n.name
}

func (n *BaseNode) SetName(value string) {
	n.name = value
}

func (n *BaseNode) GetId() string {
	return n.id
}

func (n *BaseNode) SetId(value string) {
	n.id = value
}
