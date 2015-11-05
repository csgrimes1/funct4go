package lists

import (
)

func newMapNode(target ListNode, mapper func(interface{}) interface{}) ListNode {
	return mapNode{ target: target, mapper: mapper }
}

type mapNode struct {
	target ListNode
	mapper func(interface{}) interface{}
}

func (node mapNode) Done() bool {
	return node.target.Done()
}

func (node mapNode) Value() interface{} {
	return node.mapper(node.target.Value())
}

func (node mapNode) Tail() ListNode {
	return newMapNode( node.target.Tail(), node.mapper )
}
