package flowtest

import (
	"backend/core/dataflow"
	"fmt"
)

type PrintNode struct {
	dataflow.BaseNode
}

func (n *PrintNode) DataHandler(origin dataflow.FlowData, current dataflow.FlowData) dataflow.FlowData {
	fmt.Printf("origin:%v   current:%v", origin, current)
	fmt.Println()
	return current
}

func NewPrintNode(id string, name string) dataflow.IHandleNode {
	node := &PrintNode{}
	node.Init(id, name, "print.node", node.DataHandler)
	return node
}
