package lists

import (
)

func newGeneratorList(initialValue interface{}, generator func(interface{}) interface{}) ListNode {
	return generatorList{
		value: generator(initialValue),
		generator: generator,
	}
}

type generatorList struct {
	value interface {}
	generator func(interface{}) interface{}
}

func (node generatorList) Done() bool {
	return false
}

func (node generatorList) Value() interface{} {
	return node.value
}

func (node generatorList) Tail() ListNode {
	return newGeneratorList(node.value, node.generator)
}
