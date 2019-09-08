package decisiontree

import "fmt"

type Example map[string]string

type DecisionTree struct {
	predAttr string
	domain   map[string]map[string]string
	tree     Node
	Built    bool
}

func NewDecisionTree(examples []Example, predictedAttribute string) (DecisionTree, error) {
	ret := DecisionTree{predAttr: predictedAttribute}
	ret.domain = make(map[string]map[string]string)

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

func (dt DecisionTree) String() string {
	return fmt.Sprintf("digraph G {\n%s}", dt.tree.DotString())
}
