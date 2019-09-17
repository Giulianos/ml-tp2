package bagging

import (
	"github.com/Giulianos/ml-decision-tree/classifier"
	"github.com/Giulianos/ml-decision-tree/decisiontree"
)

func (b Bagging) buildDTree(examples []classifier.Example) (decisiontree.DecisionTree, error) {
	tree := decisiontree.NewDecisionTree(b.predAttr)
	// TODO: setup any classifier properties here
	// ex.: tree.SetMaxSplits(1)
	err := tree.Fit(examples)
	return tree, err
}
