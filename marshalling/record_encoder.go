package marshalling

import "github.com/Giulianos/ml-decision-tree/classifier"

func UnmarshallRecord(headers []string, record []string) classifier.Example {
	example := classifier.Example{}
	for fieldIdx, fieldVal := range record {
		example[headers[fieldIdx]] = fieldVal
	}

	return example
}

func MarshallRecord(headers []string, example classifier.Example) []string {
	marshalledExample := make([]string, len(headers))

	for i, header := range headers {
		marshalledExample[i] = example[header]
	}

	return marshalledExample
}
