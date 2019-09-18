package encoder

import (
	"encoding/csv"
	"io"

	"github.com/Giulianos/ml-decision-tree/classifier"
)

func EncodeCSV(reader csv.Reader) ([]classifier.Example, error) {
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
		examples = append(examples, EncodeRecord(headers, record))
	}

	return examples, nil
}
