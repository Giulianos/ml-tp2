package bagging

import (
	"github.com/Giulianos/ml-decision-tree/classifier"
	"github.com/Giulianos/ml-decision-tree/decisiontree"
)

type Bagging struct {
	classifiers []decisiontree.DecisionTree
}

func NewBagging(quantity int) Bagging {
	return Bagging{
		classifiers: make([]decisiontree.DecisionTree, quantity),
	}
}

func (b Bagging) Classify(example classifier.Example) (string, float64) {
	// TODO: implement method
	return "", 0
}

func (b Bagging) GetClasses() []string {
	return []string{}
}
