package decisiontree

import (
	"math"

	"github.com/Giulianos/ml-decision-tree/classifier"
)

type ExampleFilter func(example classifier.Example) bool

func (dt DecisionTree) sEntropy(examples []classifier.Example) float64 {
	relFreq := make(map[string]float64)

	for _, example := range examples {
		relFreq[example[dt.predAttr]] += 1 / float64(len(examples))
	}

	var entropy float64
	for _, pi := range relFreq {
		entropy -= pi * math.Log2(pi)
	}

	return entropy
}

func (dt DecisionTree) svEntropy(examples []classifier.Example, attr string, val string) (entropy float64, svLen float64) {
	freqs := make(map[string]float64)
	var svSize float64

	for _, example := range examples {
		class := example[dt.predAttr]
		v := example[attr]
		if v == val {
			svSize++
			freqs[class] += 1
		}
	}

	var svEntropy float64
	for _, count := range freqs {
		pi := count / svSize
		svEntropy -= pi * math.Log2(pi)
	}

	return svEntropy, svSize
}

func (dt DecisionTree) gain(examples []classifier.Example, attr string) float64 {
	gain := dt.sEntropy(examples)
	examplesLen := float64(len(examples))
	for value := range dt.domain[attr] {
		svH, svLen := dt.svEntropy(examples, attr, value)
		gain -= (svLen / examplesLen) * svH
	}

	return gain
}

func (dt DecisionTree) giniIndex(examples []classifier.Example, attr string) float64 {
	var giniIndex float64
	for value := range dt.domain[attr] {
		branchGini, baselineProbability := dt.branchGiniIndex(examples, attributeFilter(attr, value))
		giniIndex += branchGini * baselineProbability
	}

	return giniIndex
}

func attributeFilter(attr string, val string) ExampleFilter {
	return func(example classifier.Example) bool {
		return example[attr] == val
	}
}

func (dt DecisionTree) branchGiniIndex(examples []classifier.Example, branchFilter ExampleFilter) (float64, float64) {
	classCounts := make(map[string]float64)
	var branchSize float64

	// Count occurences of each class
	for _, example := range examples {
		if branchFilter(example) {
			branchSize++
			classCounts[example[dt.predAttr]]++
		}
	}

	var giniIndex float64 = 1

	// Calculate gini index
	for _, classCount := range classCounts {
		classFrequency := classCount / branchSize
		giniIndex -= math.Pow(classFrequency, 2)
	}

	baselineProbability := branchSize / float64(len(examples))

	return giniIndex, baselineProbability
}
