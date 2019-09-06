package decisiontree

type Example map[string]string

type DecisionTree struct {
	predAttr string
	domain   map[string][]string
}

func NewDecisionTree(examples []Example, predictedAttribute string) DecisionTree {
	ret := DecisionTree{predAttr: predictedAttribute}
	domain := make(map[string]map[string]string)

	for _, example := range examples {
		for attr, value := range example {
			_, ok := domain[attr]
			if !ok {
				domain[attr] = make(map[string]string)
			}
			domain[attr][value] = "1"
		}
	}

	return ret
}
