package decisiontree

import (
	"fmt"

	"github.com/Giulianos/ml-decision-tree/classifier"
)

type DecisionTree struct {
	predAttr     string
	domain       map[string]map[string]string
	tree         Node
	Built        bool
	nodeCount    int
	maxSplits    int
	splitsCount  int
	gainFunction GainFunction
}

type GainFunction uint16

const (
	SHANNON_ENTROPY GainFunction = iota
	GINI
)

func (dt *DecisionTree) Fit(examples []classifier.Example) error {
	for _, example := range examples {
		for attr, value := range example {
			_, ok := dt.domain[attr]
			if !ok {
				dt.domain[attr] = make(map[string]string)
			}
			dt.domain[attr][value] = "1"
		}
	}

	// Build tree using id3 algorithm
	resetIdGenerator()
	tree, err := dt.buildTree(examples)

	// Save built tree
	dt.tree = tree
	dt.Built = true

	return err
}

func NewDecisionTree(predictedAttribute string) DecisionTree {
	ret := DecisionTree{predAttr: predictedAttribute}
	ret.domain = make(map[string]map[string]string)
	ret.maxSplits = -1

	return ret
}

func (dt *DecisionTree) SetMaxSplits(splits int) {
	dt.maxSplits = splits
}

func (dt *DecisionTree) SetGainFunction(function GainFunction) {
	dt.gainFunction = function
}

func (dt DecisionTree) GetClasses() []string {
	classes := make([]string, len(dt.domain[dt.predAttr]))
	var idx int
	for class := range dt.domain[dt.predAttr] {
		classes[idx] = class
		idx++
	}

	return classes
}

func (dt DecisionTree) String() string {
	return fmt.Sprintf("digraph G {\n%s}", dt.tree.DotString())
}
