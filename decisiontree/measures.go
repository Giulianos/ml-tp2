package decisiontree

import "math"

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

type svKey struct {
	class string
	val   string
}

func (dt DecisionTree) svEntropy(examples []Example, attr string, val string) (entropy float64, svLen float64) {
	relFreq := make(map[svKey]float64)
	var svSize float64

	for _, example := range examples {
		class := example[dt.predAttr]
		v := example[attr]
		if v == val {
			svSize++
			relFreq[svKey{class, val}] += 1 / float64(len(examples))
		}
	}

	var svEntropy float64
	for _, pi := range relFreq {
		svEntropy -= pi * math.Log2(pi)
	}

	return svEntropy, svSize
}

func (dt DecisionTree) gain(examples []Example, attr string) float64 {
	gain := dt.sEntropy(examples)
	examplesLen := float64(len(examples))
	for _, value := range dt.domain[attr] {
		svH, svLen := dt.svEntropy(examples, attr, value)
		gain += (svLen / examplesLen) * svH
	}

	return gain
}
