package randomforest

import "github.com/Giulianos/ml-tp2/classifier"

func (rf RandomForest) Classify(example classifier.Example) (string, float64) {
	classifications := make(map[string]float64)
	for _, decTree := range rf.classifiers {
		class, _ := decTree.Classify(example)
		classifications[class]++
	}

	return argMax(classifications), 1.0
}

func argMax(m map[string]float64) string {
	var maxVal float64
	var maxValKey string

	for k, v := range m {
		if v > maxVal {
			maxVal = v
			maxValKey = k
		}
	}

	return maxValKey
}
