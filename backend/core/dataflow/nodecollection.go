package dataflow

type NodeCollection struct {
	nodes []IHandleNode
}

func NewNodeCollection() INodeCollection {
	return &NodeCollection{
		nodes: make([]IHandleNode, 0),
	}
}

func (e *NodeCollection) Add(node IHandleNode) {
	e.nodes = append(e.nodes, node)
}

func (e *NodeCollection) Remove(node IHandleNode) {
	index := e.IndexOf(node)
	if index > -1 {
		e.nodes = append(e.nodes[:index], e.nodes[index+1:]...)
	}
}

func (e *NodeCollection) IndexOf(node IHandleNode) int {
	for i := 0; i < len(e.nodes); i++ {
		if e.nodes[i] == node {
			return i
		}
	}
	return -1
}

func (e *NodeCollection) Clear() {
	e.nodes = make([]IHandleNode, 0)
}

func (e *NodeCollection) Contains(node IHandleNode) bool {
	return e.IndexOf(node) > -1
}

func (e *NodeCollection) Count() int {
	return len(e.nodes)
}

func (e *NodeCollection) Nodes() []IHandleNode {
	return e.nodes
}
