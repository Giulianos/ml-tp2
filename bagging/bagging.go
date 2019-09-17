package bagging

import (
	"math/rand"

	"github.com/Giulianos/ml-decision-tree/decisiontree"
)

type Bagging struct {
	predAttr    string
	classifiers []decisiontree.DecisionTree
	rng         *rand.Rand
}

func NewBagging(predAttr string, quantity int, seed int64) Bagging {
	return Bagging{
		classifiers: make([]decisiontree.DecisionTree, quantity),
		rng:         rand.New(rand.NewSource(seed)),
	}
}

func (b Bagging) GetClasses() []string {
	return []string{}
}
