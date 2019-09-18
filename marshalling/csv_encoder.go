package marshalling

import (
	"encoding/csv"
	"io"

	"github.com/Giulianos/ml-decision-tree/classifier"
)

func UnmarshallCSV(reader csv.Reader) ([]classifier.Example, error) {
	examples := make([]classifier.Example, 0)
	headers, err := reader.Read()
	if err != nil {
		return nil, err
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		examples = append(examples, UnmarshallRecord(headers, record))
	}

	return examples, nil
}

func MarshallHeaders(example classifier.Example) []string {
	headers := make([]string, len(example))
	var i int
	for key := range example {
		headers[i] = key
		i++
	}

	return headers
}

func MarshallCSV(examples []classifier.Example, writer csv.Writer) {
	if len(examples) == 0 {
		return
	}
	headers := MarshallHeaders(examples[0])
	writer.Write(headers)
	for _, example := range examples {
		writer.Write(MarshallRecord(headers, example))
	}
}
