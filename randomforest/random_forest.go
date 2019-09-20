package randomforest

import (
	"math/rand"

	"github.com/Giulianos/ml-decision-tree/classifier"

	"github.com/Giulianos/ml-decision-tree/decisiontree"
)

type RandomForest struct {
	predAttr    string
	quantity    int
	classifiers []*decisiontree.DecisionTree
	classes     []string
	rng         *rand.Rand
}

func New(predAttr string, quantity int, seed int64) RandomForest {
	ret := RandomForest{
		classifiers: make([]*decisiontree.DecisionTree, quantity),
		rng:         rand.New(rand.NewSource(seed)),
		predAttr:    predAttr,
		quantity:    quantity,
	}

	for i := 0; i < ret.quantity; i++ {
		decTree := decisiontree.NewDecisionTree(predAttr)
		ret.classifiers[i] = &decTree
	}

	return ret
}

func (rf *RandomForest) SetMaxSplits(splits int) {
	for _, decTree := range rf.classifiers {
		decTree.SetMaxSplits(splits)
	}
}

func (rf *RandomForest) SetGainFunction(function decisiontree.GainFunction) {
	for _, decTree := range rf.classifiers {
		decTree.SetGainFunction(function)
	}
}

func (rf *RandomForest) SetMinSplitCount(count int) {
	for _, decTree := range rf.classifiers {
		decTree.SetMinSplitCount(count)
	}
}

func (b *RandomForest) Fit(examples []classifier.Example) error {
	for i := 0; i < b.quantity; i++ {
		err := b.classifiers[i].Fit(examples)
		if err != nil {
			return err
		}
	}

	b.saveClasses(examples)

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

func (rf RandomForest) GetAvgNodeCount() float64 {
	var nodeCount float64

	for _, dt := range rf.classifiers {
		nodeCount += float64(dt.GetNodeCount())
	}

	return nodeCount / float64(len(rf.classifiers))
}
