package decisiontree

type ValNode struct {
	val      string
	children []*Node
}

func (n ValNode) Type() NodeType {
	return VALUE
}

func (n ValNode) IsLeaf() bool {
	return false
}

func (n ValNode) Children() []*Node {
	return n.children
}

func (n ValNode) Tag() string {
	return n.val
}
