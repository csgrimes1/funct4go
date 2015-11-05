package lists


import (
)

func maybeAdvance(target ListNode, predicate func(interface{}) bool) (bool, ListNode) {
	if target.Done() {
		return false, emptyNode()
	} else if predicate(target.Value()) {
		return true, target
	}
	//Tail recurse.
	return maybeAdvance(target.Tail(), predicate)
}

func newFilterNode(target ListNode, predicate func(interface{}) bool) ListNode {
	/*
		When creating a new ListNode, advance past any nodes that don't
		match the predicate.
	 */
	ok, next := maybeAdvance(target, predicate)
	if !ok {
		return emptyNode()
	}
	return filterList{ target: next, predicate: predicate }
}

type filterList struct {
	target ListNode
	predicate func(interface{}) bool
}

func (node filterList) Done() bool {
	return node.target.Done()
}

func (node filterList) Value() interface{} {
	return node.target.Value()
}

func (node filterList) Tail() ListNode {
	return newFilterNode( node.target.Tail(), node.predicate)
}
