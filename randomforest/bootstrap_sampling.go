package randomforest

import "github.com/Giulianos/ml-decision-tree/classifier"

func (rf RandomForest) getBootstrapSample(examples []classifier.Example) []classifier.Example {
	sample := make([]classifier.Example, len(examples))

	for i := range sample {
		randIdx := rf.rng.Uint64()
		sample[i] = examples[randIdx]
	}

	return sample
}
