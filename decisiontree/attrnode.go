package decisiontree

type AttrNode struct {
	attr     string
	children []*Node
}

func (n AttrNode) Type() NodeType {
	return ATTR
}

func (n AttrNode) IsLeaf() bool {
	return false
}

func (n AttrNode) Children() []*Node {
	return n.children
}

func (n AttrNode) Tag() string {
	return n.attr
}
