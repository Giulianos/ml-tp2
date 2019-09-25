package decisiontree

import (
	"fmt"
	"math"

	"github.com/Giulianos/ml-tp2/classifier"
)

type DecisionTree struct {
	predAttr      string
	domain        map[string]map[string]string
	tree          Node
	Built         bool
	nodeCount     int
	maxNodeCount  int
	maxSplits     int
	maxDepth      int
	minSplitCount int
	splitsCount   int
	gainFunction  GainFunction
}

type GainFunction uint16

const (
	SHANNON_ENTROPY GainFunction = iota
	GINI
)

func (dt *DecisionTree) Fit(examples []classifier.Example) error {
	for _, example := range examples {
		for attr, value := range example {
			_, ok := dt.domain[attr]
			if !ok {
				dt.domain[attr] = make(map[string]string)
			}
			dt.domain[attr][value] = "1"
		}
	}

	// Build tree using id3 algorithm
	resetIdGenerator()
	tree, err := dt.buildTree(examples)

	// Save built tree
	dt.tree = tree
	dt.Built = true

	return err
}

func New(predictedAttribute string) DecisionTree {
	ret := DecisionTree{predAttr: predictedAttribute}
	ret.domain = make(map[string]map[string]string)
	ret.maxSplits = math.MaxInt64
	ret.maxDepth = math.MaxInt64
	ret.minSplitCount = 1
	ret.maxNodeCount = math.MaxInt64

	return ret
}

func (dt *DecisionTree) SetMaxSplits(splits int) {
	dt.maxSplits = splits
}

func (dt *DecisionTree) SetGainFunction(function GainFunction) {
	dt.gainFunction = function
}

func (dt *DecisionTree) SetMinSplitCount(count int) {
	dt.minSplitCount = count
}

func (dt DecisionTree) GetNodeCount() int {
	return dt.nodeCount
}

func (dt *DecisionTree) SetMaxNodeCount(count int) {
	dt.maxNodeCount = count
}

func (dt DecisionTree) GetClasses() []string {
	classes := make([]string, len(dt.domain[dt.predAttr]))
	var idx int
	for class := range dt.domain[dt.predAttr] {
		classes[idx] = class
		idx++
	}

	return classes
}

func (dt DecisionTree) GetPredictableAttribute() string {
	return dt.predAttr
}

func (dt DecisionTree) String() string {
	return fmt.Sprintf("digraph G {\n%s}", dt.tree.DotString())
}
