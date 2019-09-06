package decisiontree

type ClassNode struct {
	class string
}

func (n ClassNode) Type() NodeType {
	return CLASS
}

func (n ClassNode) IsLeaf() bool {
	return true
}

func (n ClassNode) Children() []*Node {
	return nil
}

func (n ClassNode) Tag() string {
	return n.class
}
