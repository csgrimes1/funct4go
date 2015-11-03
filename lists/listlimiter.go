package lists


import (
)

func newLimiter(index int, target ListNode, predicate func(index int, value interface{}) bool) ListNode {
	if target.Done() {
		return emptyNode()
	} else if !predicate(index, target.Value()) {
		return emptyNode()
	}
	return limiter{index: index, target: target, predicate: predicate }
}

type limiter struct {
	target ListNode
	predicate func(index int, value interface{}) bool
	index int
}

func (node limiter) Done() bool {
	return node.target.Done()
}

func (node limiter) Value() interface{} {
	return node.target.Value()
}

func (node limiter) Tail() ListNode {
	return newLimiter(node.index + 1, node.target.Tail(), node.predicate)
}
