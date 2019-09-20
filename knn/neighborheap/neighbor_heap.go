package neighborheap

import "github.com/Giulianos/ml-decision-tree/classifier"

type Neighbor struct {
	Neighbor *classifier.Example
	Distance float64
}

type NeighborHeap []Neighbor

func (nh NeighborHeap) Len() int { return len(nh) }

func (nh NeighborHeap) Less(i, j int) bool { return nh[i].Distance > nh[j].Distance }

func (nh NeighborHeap) Swap(i, j int) { nh[i], nh[j] = nh[j], nh[i] }

func (nh *NeighborHeap) Push(x interface{}) { *nh = append(*nh, x.(Neighbor)) }

func (nh *NeighborHeap) Pop() interface{} {
	old := *nh
	n := len(old)
	ret := old[n-1]
	*nh = old[0 : n-1]
	return ret
}
