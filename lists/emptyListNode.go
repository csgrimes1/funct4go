package lists

import (
)

func emptyNode() ListNode {
	return emptyListNode{ }
}

type emptyListNode struct {
}

func (node emptyListNode) Done() bool {
	return true
}

func (node emptyListNode) Value() interface{} {
	return 0
}

func (node emptyListNode) Tail() ListNode {
	return emptyNode()
}
