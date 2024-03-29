package randomforest

import "github.com/Giulianos/ml-tp2/classifier"

func (rf RandomForest) getBootstrapSample(examples []classifier.Example) []classifier.Example {
	sample := make([]classifier.Example, len(examples))

	for i := range sample {
		randIdx := rf.rng.Intn(len(examples))
		sample[i] = examples[randIdx]
	}

	return sample
}
