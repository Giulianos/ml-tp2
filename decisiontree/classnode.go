package decisiontree

import "fmt"

type ClassNode struct {
	class string
	id    uint64
}

func NewClassNode(class string) ClassNode {
	return ClassNode{class: class, id: generateId()}
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
