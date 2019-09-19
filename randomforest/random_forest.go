package randomforest

import (
	"math/rand"

	"github.com/Giulianos/ml-decision-tree/classifier"

	"github.com/Giulianos/ml-decision-tree/decisiontree"
)

type RandomForest struct {
	predAttr    string
	classifiers []decisiontree.DecisionTree
	classes     []string
	rng         *rand.Rand
}

func New(predAttr string, quantity int, seed int64) RandomForest {
	ret := RandomForest{
		classifiers: make([]decisiontree.DecisionTree, quantity),
		rng:         rand.New(rand.NewSource(seed)),
	}

	for i := 0; i < quantity; i++ {
		ret.classifiers = append(ret.classifiers, decisiontree.NewDecisionTree(predAttr))
	}

	return ret
}

func (b *RandomForest) Fit(examples []classifier.Example) error {
	for _, decTree := range b.classifiers {
		err := decTree.Fit(examples)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *RandomForest) saveClasses(examples []classifier.Example) {
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

func (b RandomForest) GetClasses() []string {
	return b.classes
}

func (b RandomForest) GetPredictableAttribute() string {
	return b.predAttr
}
