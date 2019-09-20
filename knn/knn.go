package knn

import (
	"container/heap"
	"math"

	"github.com/Giulianos/ml-decision-tree/classifier"
	"github.com/Giulianos/ml-decision-tree/knn/neighborheap"
)

type distanceFunc func(ex1, ex2 classifier.Example) float64

type KNN struct {
	predAttr string
	distance distanceFunc
	examples []classifier.Example
	k        int
	weighted bool
}

func New(k int, predAttr string, distance distanceFunc) KNN {
	return KNN{
		distance: distance,
		predAttr: predAttr,
		k:        k,
	}
}

func (knn *KNN) SetWeighted(value bool) {
	knn.weighted = value
}

func (knn *KNN) Fit(examples []classifier.Example) {
	knn.examples = make([]classifier.Example, len(examples))
	copy(knn.examples, examples)
}

func (knn KNN) getKNearest(example classifier.Example) neighborheap.NeighborHeap {
	nearest := &neighborheap.NeighborHeap{}
	heap.Init(nearest)

	for i, neighbor := range knn.examples {
		dist := knn.distance(example, neighbor)
		heap.Push(nearest, neighborheap.Neighbor{Neighbor: &knn.examples[i], Distance: dist})

		// Check if we have to delete the excess
		if nearest.Len() > knn.k {
			heap.Pop(nearest)
		}
	}

	return *nearest
}

func (knn KNN) Classify(example classifier.Example) (string, float64) {
	contrib := map[string]float64{}

	nearest := knn.getKNearest(example)

	for _, neigh := range nearest {
		if neigh.Distance == 0 {
			return (*neigh.Neighbor)[knn.predAttr], 1.0
		}
		var weight float64 = 1
		if knn.weighted {
			weight *= 1 / math.Pow(neigh.Distance, 2)
		}
		contrib[(*neigh.Neighbor)[knn.predAttr]] += weight
	}

	// Find the mode in the contrib map
	var maxClass string
	var maxContrib float64

	for class, val := range contrib {
		if val > maxContrib {
			maxClass = class
			maxContrib = val
		}
	}

	return maxClass, 1.0
}

func (knn KNN) GetClasses() []string {
	// TODO: implement method
	return []string{}
}

func (knn KNN) GetPredictableAttribute() string {
	return knn.predAttr
}
