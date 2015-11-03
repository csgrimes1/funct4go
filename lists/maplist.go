package lists

import (
)

func newMapList(target ListNode, mapper func(interface{}) interface{}) ListNode {
	return mapList{ target: target, mapper: mapper }
}

type mapList struct {
	target ListNode
	mapper func(interface{}) interface{}
}

func (node mapList) Done() bool {
	return node.target.Done()
}

func (node mapList) Value() interface{} {
	return node.mapper(node.target.Value())
}

func (node mapList) Tail() ListNode {
	return newMapList( node.target.Tail(), node.mapper )
}
