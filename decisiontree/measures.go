package decisiontree

import (
	"math"
)

func (dt DecisionTree) sEntropy(examples []Example) float64 {
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

func (dt DecisionTree) svEntropy(examples []Example, attr string, val string) (entropy float64, svLen float64) {
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

func (dt DecisionTree) gain(examples []Example, attr string) float64 {
	gain := dt.sEntropy(examples)
	examplesLen := float64(len(examples))
	for value := range dt.domain[attr] {
		svH, svLen := dt.svEntropy(examples, attr, value)
		gain -= (svLen / examplesLen) * svH
	}

	return gain
}
