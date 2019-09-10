package decisiontree

import "log"

func (dt DecisionTree) Classify(example Example) string {
	return recClassify(dt.tree, example)
}

func recClassify(node Node, example Example) string {
	if node.IsLeaf() && node.Type() == CLASS {
		return node.Tag()
	}

	if node.Type() != ATTR {
		log.Fatal("malformed tree")
	}

	var nextNode *Node
	for _, child := range node.Children() {
		if example[node.Tag()] == child.Tag() {
			nextNode = &(child.Children()[0])
		}
	}

	if nextNode == nil {
		log.Print("couldn't classify example")
		return ""
	}

	return recClassify(*nextNode, example)
}
