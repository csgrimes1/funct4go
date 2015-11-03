package lists

type Maybe struct {
	List
}

func CreateOptionalResult(value interface{}) Maybe {
	return Maybe{List: List {
		InnerCollection: maybeNode{
			hasValue: true,
			result: value,
		},
	}}
}

func CreateOptionalEmpty() Maybe {
	return Maybe{List: List {
		InnerCollection: emptyNode(),
	},}
}

func CreateBooleanResult(result bool) Maybe {
	return Maybe {List: List{
		InnerCollection: maybeNode{
			hasValue:	result,
			result:		result,
		},
	},}
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
