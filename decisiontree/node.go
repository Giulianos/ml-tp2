package decisiontree

type NodeType uint16

const (
	ATTR NodeType = iota
	VALUE
	CLASS
)

func (nt NodeType) String() string {
	switch nt {
	case ATTR:
		return "Attribute"
	case VALUE:
		return "Value"
	case CLASS:
		return "Class"
	}
	return "Undefined Node Type"
}

type Node interface {
	Type() NodeType
	IsLeaf() bool
	Children() []*Node
	Tag() string
}
