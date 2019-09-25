package randomforest

import (
	"math/rand"

	"github.com/Giulianos/ml-tp2/classifier"

	"github.com/Giulianos/ml-tp2/decisiontree"
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
		decTree := decisiontree.New(predAttr)
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

func (rf *RandomForest) SetMaxNodeCount(count int) {
	for _, decTree := range rf.classifiers {
		decTree.SetMaxNodeCount(count)
	}
}

func (rf *RandomForest) Fit(examples []classifier.Example) error {
	for i := 0; i < rf.quantity; i++ {
		err := rf.classifiers[i].Fit(rf.getBootstrapSample(examples))
		if err != nil {
			return err
		}
	}

	rf.saveClasses(examples)

	return nil
}

func (rf *RandomForest) saveClasses(examples []classifier.Example) {
	classesSet := make(map[string]int)
	for _, example := range examples {
		class := example[rf.predAttr]
		classesSet[class]++
	}

	rf.classes = make([]string, len(classesSet))
	for class := range classesSet {
		rf.classes = append(rf.classes, class)
	}
}

func (rf RandomForest) GetClasses() []string {
	return rf.classes
}

func (rf RandomForest) GetPredictableAttribute() string {
	return rf.predAttr
}

func (rf RandomForest) GetAvgNodeCount() float64 {
	var nodeCount float64

	for _, dt := range rf.classifiers {
		nodeCount += float64(dt.GetNodeCount())
	}

	return nodeCount / float64(len(rf.classifiers))
}
