package bagging

import "github.com/Giulianos/ml-decision-tree/classifier"

func (b Bagging) getBootstrapSample(examples []classifier.Example) []classifier.Example {
	sample := make([]classifier.Example, len(examples))

	for i := range sample {
		randIdx := b.rng.Uint64()
		sample[i] = examples[randIdx]
	}

	return sample
}
