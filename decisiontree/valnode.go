package decisiontree

import "fmt"

type ValNode struct {
	val      string
	children *[]Node
	id       uint64
	depth    int
}

func NewValNode(val string) ValNode {
	var childrenSlice []Node
	return ValNode{val: val, id: generateId(), children: &childrenSlice}
}

func (n ValNode) Type() NodeType {
	return VALUE
}

func (n ValNode) IsLeaf() bool {
	return false
}

func (n ValNode) Children() []Node {
	return *n.children
}

func (n ValNode) Tag() string {
	return n.val
}

func (n ValNode) Id() uint64 {
	return n.id
}

func (n ValNode) Depth() int {
	return n.depth
}

func (n ValNode) DotString() string {
	nodeAttributes := fmt.Sprintf("%d[label=\"%s\"]\n", n.id, n.val)
	childrenDotString := ""
	for _, child := range *n.children {
		childrenDotString += fmt.Sprintf("%d -> %d\n", n.id, child.Id())
		childrenDotString += child.DotString()
	}
	return nodeAttributes + childrenDotString
}

func (n ValNode) AddChild(child Node) error {
	*n.children = append(*n.children, child)

	return nil
}
