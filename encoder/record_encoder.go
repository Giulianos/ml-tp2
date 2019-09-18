package encoder

import "github.com/Giulianos/ml-decision-tree/classifier"

func EncodeRecord(headers []string, record []string) classifier.Example {
	example := classifier.Example{}
	for fieldIdx, fieldVal := range record {
		example[headers[fieldIdx]] = fieldVal
	}

	return example
}
