package lists

type Maybe struct {
	List
}

func CreateOptionalResult(value interface{}) Maybe {
	return Maybe {
		List: newList(newSliceNode([]interface{}{value})),
	}
}

func CreateOptionalEmpty() Maybe {
	return Maybe {
		List: newEmptyList(),
	}
}

func CreateBooleanResult(result bool) Maybe {
	if result {
		return CreateOptionalResult(result)
	}
	return CreateOptionalEmpty()
}

type maybeNode struct {
	hasValue bool
	result interface{}
}

func (node maybeNode) Done() bool {
	return !node.hasValue
}

func (node maybeNode) Value() interface{} {
	return node.result
}

func (node maybeNode) Tail() ListNode {
	return emptyNode()
}
