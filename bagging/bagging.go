package bagging

import (
	"math/rand"

	"github.com/Giulianos/ml-decision-tree/classifier"

	"github.com/Giulianos/ml-decision-tree/decisiontree"
)

type Bagging struct {
	predAttr    string
	classifiers []decisiontree.DecisionTree
	classes     []string
	rng         *rand.Rand
}

func NewBagging(predAttr string, quantity int, seed int64) Bagging {
	ret := Bagging{
		classifiers: make([]decisiontree.DecisionTree, quantity),
		rng:         rand.New(rand.NewSource(seed)),
	}

	for i := 0; i < quantity; i++ {
		ret.classifiers = append(ret.classifiers, decisiontree.NewDecisionTree(predAttr))
	}

	return ret
}

func (b *Bagging) Fit(examples []classifier.Example) error {
	for _, decTree := range b.classifiers {
		err := decTree.Fit(examples)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Bagging) saveClasses(examples []classifier.Example) {
	classesSet := make(map[string]int)
	for _, example := range examples {
		class := example[b.predAttr]
		classesSet[class]++
	}

	b.classes = make([]string, len(classesSet))
	for class := range classesSet {
		b.classes = append(b.classes, class)
	}
}

func (b Bagging) GetClasses() []string {
	return b.classes
}
