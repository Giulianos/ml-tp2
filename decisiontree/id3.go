package decisiontree

import (
	"log"
	"math"

	"github.com/Giulianos/ml-decision-tree/classifier"
)

func (dt *DecisionTree) buildTree(examples []classifier.Example) (Node, error) {

	sameClass := dt.isDataPure(examples)

	// All examples are classified with the same class
	if sameClass != nil {
		dt.nodeCount++
		return NewClassNode(*sameClass), nil
	}

	// Cannot split further put mode class
	if dt.splitsCount == dt.maxSplits {
		modeClass := dt.modeClass(examples)
		return NewClassNode(*modeClass), nil
	}

	// Otherwise find discriminant attribute and build the node
	discrAttr := dt.discriminantAttribute(examples)
	attrNode := NewAttrNode(*discrAttr)
	dt.nodeCount++
	dt.splitsCount++

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

func (dt *DecisionTree) buildValSubTrees(examples []classifier.Example, attr string) ([]ValNode, error) {
	splittedData := dt.splitData(examples, attr)
	nodes := make([]ValNode, len(splittedData))

	modeClass := dt.modeClass(examples)

	var idx int
	for val, valexamples := range splittedData {
		node := NewValNode(val)
		dt.nodeCount++
		var subtree Node
		if len(valexamples) != 0 {
			subtree, _ = dt.buildTree(valexamples)
		} else {
			subtree = NewClassNode(*modeClass)
			dt.nodeCount++
		}
		err := node.AddChild(subtree)
		if err != nil {
			return nil, err
		}
		nodes[idx] = node
		idx++
	}

	return nodes, nil
}

func (dt DecisionTree) splitData(examples []classifier.Example, attr string) map[string][]classifier.Example {
	splittedData := map[string][]classifier.Example{}
	log.Printf("Splitted at %s", attr)

	for _, example := range examples {
		value := example[attr]
		_, ok := splittedData[value]
		if !ok {
			splittedData[value] = []classifier.Example{}
		}
		splittedData[value] = append(splittedData[value], example)
	}

	return splittedData
}

func (dt DecisionTree) discriminantAttribute(examples []classifier.Example) *string {
	var discrAttr string
	var discrAttrGain float64

	if dt.gainFunction == GINI {
		discrAttrGain = math.MaxFloat64
	}

	for attr := range dt.domain {
		if attr == dt.predAttr {
			continue
		}
		if dt.gainFunction == SHANNON_ENTROPY {
			attrGain := dt.gain(examples, attr)
			if attrGain > discrAttrGain {
				discrAttr = attr
				discrAttrGain = attrGain
			}
		}
		if dt.gainFunction == GINI {
			attrGini := dt.giniIndex(examples, attr)
			log.Printf("Gini(%s)=%f", attr, attrGini)
			if attrGini < discrAttrGain {
				discrAttr = attr
				discrAttrGain = attrGini
			}
		}
	}

	return &discrAttr
}

func (dt DecisionTree) isDataPure(examples []classifier.Example) *string {
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

func (dt DecisionTree) modeClass(examples []classifier.Example) *string {
	var counts = make(map[string]int)

	for _, example := range examples {
		counts[example[dt.predAttr]]++
	}

	var modeClass string
	var modeClassCount int

	for class, count := range counts {
		if count > modeClassCount {
			modeClass = class
			modeClassCount = count
		}
	}

	return &modeClass
}
