package knn

import (
	"container/heap"

	"github.com/Giulianos/ml-decision-tree/classifier"
	"github.com/Giulianos/ml-decision-tree/knn/neighborheap"
)

type distanceFunc func(ex1, ex2 classifier.Example) float64

type KNN struct {
	predAttr string
	distance distanceFunc
	examples []classifier.Example
	k        int
}

func New(k int, predAttr string, distance distanceFunc) KNN {
	return KNN{
		distance: distance,
		predAttr: predAttr,
		k:        k,
	}
}

func (knn *KNN) Fit(examples []classifier.Example) {
	knn.examples = make([]classifier.Example, len(examples))
	copy(knn.examples, examples)
}

func (knn KNN) getKNearest(example classifier.Example) neighborheap.NeighborHeap {
	nearest := &neighborheap.NeighborHeap{}
	heap.Init(nearest)

	for _, neighbor := range knn.examples {
		dist := knn.distance(example, neighbor)
		heap.Push(nearest, neighborheap.Neighbor{Neighbor: &neighbor, Distance: dist})

		// Check if we have to delete the excess
		if nearest.Len() > knn.k {
			heap.Pop(nearest)
		}
	}

	return *nearest
}

func (knn KNN) Classify(example classifier.Example) (string, float64) {
	// TODO: implement method
	return "", 1.0
}

func (knn KNN) GetClasses() []string {
	// TODO: implement method
	return []string{}
}

func (knn KNN) GetPredictableAttribute() string {
	return knn.predAttr
}
