package decisiontree

import "fmt"

type AttrNode struct {
	attr     string
	children *[]Node
	id       uint64
	depth    *int
	mode     string
}

func NewAttrNode(attr string) AttrNode {
	var childrenSlice []Node
	var depth int
	return AttrNode{attr: attr, id: generateId(), children: &childrenSlice, depth: &depth}
}

func (n AttrNode) Type() NodeType {
	return ATTR
}

func (n AttrNode) IsLeaf() bool {
	return false
}

func (n AttrNode) Children() []Node {
	return *n.children
}

func (n AttrNode) Tag() string {
	return n.attr
}

func (n AttrNode) Id() uint64 {
	return n.id
}

func (n AttrNode) Depth() int {
	return *n.depth
}

func (n AttrNode) DotString() string {
	nodeAttributes := fmt.Sprintf("%d[label=\"%s\"]\n", n.id, n.attr)
	childrenDotString := ""
	for _, child := range *n.children {
		childrenDotString += fmt.Sprintf("%d -> %d\n", n.id, child.Id())
		childrenDotString += child.DotString()
	}
	return nodeAttributes + childrenDotString
}

func (n AttrNode) AddChild(child Node) error {
	*n.children = append(*n.children, child)

	return nil
}

func (n AttrNode) SetDepth(depth int) {
	*n.depth = depth
}

func (n *AttrNode) SetMode(mode string) {
	n.mode = mode
}

func (n AttrNode) GetMode() string {
	return n.mode
}
