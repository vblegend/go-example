package flowtest

import "backend/core/dataflow"

type AddNode struct {
	dataflow.BaseNode
}

func (n *AddNode) DataHandler(origin dataflow.FlowData, current dataflow.FlowData) dataflow.FlowData {

	switch t := current.(type) {
	case int:
		{
			return t + 1
		}
	default:
		{
			return t
		}
	}
}

func NewAddNode(id string, name string) dataflow.IHandleNode {
	node := &AddNode{}
	node.Init(id, name, "add.node", node.DataHandler)
	return node
}
