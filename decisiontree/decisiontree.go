package decisiontree

import (
	"fmt"

	"github.com/Giulianos/ml-decision-tree/classifier"
)

type DecisionTree struct {
	predAttr    string
	domain      map[string]map[string]string
	tree        Node
	Built       bool
	nodeCount   int
	maxSplits   int
	splitsCount int
}

func NewDecisionTree(examples []classifier.Example, predictedAttribute string) (DecisionTree, error) {
	ret := DecisionTree{predAttr: predictedAttribute}
	ret.domain = make(map[string]map[string]string)
	ret.maxSplits = 3

	for _, example := range examples {
		for attr, value := range example {
			_, ok := ret.domain[attr]
			if !ok {
				ret.domain[attr] = make(map[string]string)
			}
			ret.domain[attr][value] = "1"
		}
	}

	// Build tree using id3 algorithm
	resetIdGenerator()
	tree, err := ret.buildTree(examples)

	// Save built tree
	ret.tree = tree
	ret.Built = true

	return ret, err
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
