package decisiontree

func (dt DecisionTree) buildTree(examples []Example) (Node, error) {
	sameClass := dt.isDataPure(examples)

	// All examples are classified with the same class
	if sameClass != nil {
		return NewClassNode(*sameClass), nil
	}

	// Otherwise find discriminant attribute and build the node
	discrAttr := dt.discriminantAttribute(examples)
	attrNode := NewAttrNode(*discrAttr)

	// Get val subtrees for each attribute value
	valNodes, err := dt.buildValSubTrees(examples, *discrAttr)
	if err != nil {
		return nil, err
	}

	// Add val subtrees to the attr node
	for _, node := range valNodes {
		e := attrNode.AddChild(node)
		if e != nil {
			return nil, e
		}
	}

	return attrNode, nil
}

func (dt DecisionTree) buildValSubTrees(examples []Example, attr string) ([]ValNode, error) {
	splittedData := dt.splitData(examples, attr)
	nodes := make([]ValNode, len(splittedData))

	var idx int
	for val, examples := range splittedData {
		node := NewValNode(val)
		subtree, _ := dt.buildTree(examples)
		err := node.AddChild(subtree)
		if err != nil {
			return nil, err
		}
		nodes[idx] = node
		idx++
	}

	return nodes, nil
}

func (dt DecisionTree) splitData(examples []Example, attr string) map[string][]Example {
	splittedData := map[string][]Example{}

	for _, example := range examples {
		value := example[attr]
		_, ok := splittedData[value]
		if !ok {
			splittedData[value] = []Example{}
		}
		splittedData[value] = append(splittedData[value], example)
	}

	return splittedData
}

func (dt DecisionTree) discriminantAttribute(examples []Example) *string {
	var discrAttr string
	var discrAttrGain float64

	for attr := range dt.domain {
		if attr == dt.predAttr {
			continue
		}
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
