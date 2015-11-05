package lists

type BasicList interface {
	Done() bool
	Value() (bool, interface{})
	Next() (bool,List)

	node() ListNode
}



type FunctionalList interface {
	Count() int

	Map(mapper func(interface{}) interface{}) List
	MapT(mapperFunc interface{}) List
	Filter(predicate func(interface{}) bool) List
	FilterT(predicate interface{}) List
	Take(length int) List
	TakeWhile( predicate func(value interface{}) bool ) List
	TakeWhileT( predicate interface{} ) List
	Skip(number int) List
	Fold(initialValue interface{}, accumulator func(value interface{}, accumulation interface{}) interface{}) interface{}
	FoldT(initialValue interface{}, accumulator interface{}) interface{}
	FlatMap() List
	String() string
}

type List interface {
	BasicList
	FunctionalList
}