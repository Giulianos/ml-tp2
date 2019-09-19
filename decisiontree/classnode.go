package decisiontree

import "fmt"

type ClassNode struct {
	class string
	id    uint64
	depth *int
}

func NewClassNode(class string) ClassNode {
	var depth int
	return ClassNode{class: class, id: generateId(), depth: &depth}
}

func (n ClassNode) Type() NodeType {
	return CLASS
}

func (n ClassNode) IsLeaf() bool {
	return true
}

func (n ClassNode) Children() []Node {
	return nil
}

func (n ClassNode) Tag() string {
	return n.class
}

func (n ClassNode) Id() uint64 {
	return n.id
}

func (n ClassNode) DotString() string {
	return fmt.Sprintf("%d[label=\"%s\", shape=\"box\", color=\"blue\"]\n", n.id, n.class)
}

func (n ClassNode) AddChild(child Node) error {
	return fmt.Errorf("Can't add children to a class node\n")
}

func (n ClassNode) Depth() int {
	return *n.depth
}

func (n ClassNode) SetDepth(depth int) {
	*n.depth = depth
}
