package lists


type ListNode interface {
	Value() interface{}
	Tail() ListNode
	Done() bool
}

func safeNextNode(node ListNode) ListNode {
	if node.Done() {
		return emptyNode()
	}

	return node.Tail()
}