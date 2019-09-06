package decisiontree

type AttrNode struct {
	attr     string
	children []*ValNode
}

type ValNode struct {
	val      string
	children []*AttrNode
	class    *ClassNode
}

type ClassNode struct {
	class string
}
