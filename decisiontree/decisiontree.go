package decisiontree

import "fmt"

type Example map[string]string

type DecisionTree struct {
	predAttr string
	domain   map[string]map[string]string
	tree     Node
}

func NewDecisionTree(examples []Example, predictedAttribute string) DecisionTree {
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
	ret.tree = ret.buildTree(examples)

	return ret
}

func (dt DecisionTree) String() string {
	return fmt.Sprintf("digraph G {\n%s}", dt.tree.DotString())
}
