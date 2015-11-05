package lists

import (
)


func firstChildNode(parentNode ListNode) ListNode {
	if parentNode.Done() {
		return emptyNode()
	}

	asList := parentNode.Value().(List)
	if asList.Done() {
		return emptyNode()
	}
	return asList.node()
}


func newChildNodes( parentNode ListNode) ListNode {
	//this=parent1 -> parent2 -> .. -> parentN
	//children1       children2        childrenN
	//append( List[Lists] )
	possibleResult := flatteningList{
		later: safeNextNode(parentNode),
		sooner: firstChildNode(parentNode),
	}
	if !possibleResult.sooner.Done() {
		return possibleResult
	} else if possibleResult.later.Done() {
		return emptyListNode{}
	}
	return newChildNodes(possibleResult.later)
}

type flatteningList struct {
	sooner, later ListNode
}

func (node flatteningList) Done() bool {
	if !node.sooner.Done() {
		return false
	}
	return node.Tail().Done()
}

func (node flatteningList) Value() interface{} {
	return node.sooner.Value()
}

func (node flatteningList) Tail() ListNode {
	if node.sooner.Done() {
	} else if !node.sooner.Tail().Done() {
		return flatteningList{
			later: node.later,
			sooner: node.sooner.Tail(),
		}
	}

	return newChildNodes(node.later)
}
