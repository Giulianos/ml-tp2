package decisiontree

import "sync"

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

var nextID uint64 = 0

var mutex = &sync.Mutex{}

func generateId() uint64 {
	mutex.Lock()
	id := nextID
	nextID++
	mutex.Unlock()
	return id
}

func resetIdGenerator() {
	mutex.Lock()
	nextID = 0
	mutex.Unlock()
}

type Node interface {
	Type() NodeType
	IsLeaf() bool
	Children() []Node
	Tag() string
	Id() uint64
	DotString() string
	AddChild(child Node) error
}
