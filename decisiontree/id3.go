package decisiontree

import "log"

func (dt DecisionTree) buildTree(examples []Example) Node {
	sameClass := dt.isDataPure(examples)

	// All examples are classified with the same class
	if sameClass != nil {
		return NewClassNode(*sameClass)
	}

	// Otherwise find discriminant attribute and build the node
	discrAttr := dt.discriminantAttribute(examples)
	attrNode := NewAttrNode(*discrAttr)

	// TODO: filter examples
	filteredExamples := examples
	valNodes := dt.buildValSubTrees(filteredExamples)

	// Add val subtrees to the attr node
	for _, node := range valNodes {
		e := attrNode.AddChild(node)
		if e != nil {
			log.Fatal(e)
		}
	}

	return attrNode
}

func (dt DecisionTree) buildValSubTrees(examples []Example) []ValNode {
	// TODO: implement method
	return nil
}

func (dt DecisionTree) discriminantAttribute(examples []Example) *string {
	var discrAttr string
	var discrAttrGain float64

	for attr := range dt.domain {
		attrGain := dt.gain(examples, attr)
		if attrGain > discrAttrGain {
			discrAttr = attr
			discrAttrGain = attrGain
		}
	}

	return &discrAttr
}

func (dt DecisionTree) isDataPure(examples []Example) *string {
	var currentClass string
	for idx, example := range examples {
		if idx == 0 {
			currentClass = example[dt.predAttr]
		} else if currentClass != example[dt.predAttr] {
			return nil
		}
	}

	return &currentClass
}
