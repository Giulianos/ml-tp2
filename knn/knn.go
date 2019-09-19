package knn

import "github.com/Giulianos/ml-decision-tree/classifier"

type distanceFunc func(ex1, ex2 classifier.Example) float64

type KNN struct {
	predAttr string
	distance distanceFunc
	examples []classifier.Example
}

func New(predAttr string, distance distanceFunc) KNN {
	return KNN{
		distance: distance,
		predAttr: predAttr,
	}
}

func (k *KNN) Fit(examples []classifier.Example) {
	k.examples = make([]classifier.Example, len(examples))

	for idx, example := range examples {
		k.examples[idx] = example
	}
}

func (k KNN) Classify(example classifier.Example) (string, float64) {
	// TODO: implement method
	return "", 1.0
}

func (k KNN) GetClasses() []string {
	// TODO: implement method
	return []string{}
}

func (k KNN) GetPredictableAttribute() string {
	return k.predAttr
}
