package randomforest

import (
	"github.com/Giulianos/ml-decision-tree/classifier"
	"github.com/Giulianos/ml-decision-tree/decisiontree"
)

func (rf RandomForest) buildDTree(examples []classifier.Example) (decisiontree.DecisionTree, error) {
	tree := decisiontree.NewDecisionTree(rf.predAttr)
	// TODO: setup any classifier properties here
	// ex.: tree.SetMaxSplits(1)
	err := tree.Fit(examples)
	return tree, err
}
