package decisiontree

import (
	"log"

	"github.com/Giulianos/ml-decision-tree/classifier"
)

func (dt DecisionTree) Classify(example classifier.Example) (string, float64) {
	return recClassify(dt.tree, example), 1.0
}

func recClassify(node Node, example classifier.Example) string {
	if node.IsLeaf() && node.Type() == CLASS {
		return node.Tag()
	}

	if node.Type() != ATTR {
		log.Fatal("malformed tree")
	}

	// Here we know that node is an attr node (the cast is safe)
	attrNode, ok := node.(AttrNode)
	if !ok {
		// Something really bad happened
		log.Fatal("malformed tree")
	}

	var nextNode *Node
	for _, child := range node.Children() {
		if example[node.Tag()] == child.Tag() {
			nextNode = &(child.Children()[0])
		}
	}

	if nextNode == nil {
		return attrNode.GetMode()
	}

	return recClassify(*nextNode, example)
}
