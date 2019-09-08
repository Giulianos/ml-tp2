package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/Giulianos/ml-decision-tree/decisiontree"
)

func csvToExamples(r io.Reader) ([]decisiontree.Example, error) {
	csvReader := csv.NewReader(r)

	attrs, err := csvReader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading dataset")
	}

	var examples []decisiontree.Example

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		newExample := decisiontree.Example{}
		for fieldPos, fieldValue := range record {
			newExample[attrs[fieldPos]] = fieldValue
		}
		examples = append(examples, newExample)
	}

	return examples, nil
}

func HandleNewTree(dt *decisiontree.DecisionTree, dtMutex *sync.Mutex) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		bufWriter := bufio.NewWriter(writer)

		// Create dataset from received csv
		examples, err := csvToExamples(request.Body)

		// Get predicted attribute
		predAttr := request.URL.Query()["pred-attr"][0]
		log.Println(predAttr)

		// Create tree
		tree, err := decisiontree.NewDecisionTree(examples, predAttr)
		if err != nil {
			writer.WriteHeader(500)
			bufWriter.WriteString("Error building the tree")
		} else {
			// Save created tree
			dtMutex.Lock()
			*dt = tree
			dtMutex.Unlock()
			writer.WriteHeader(200)
			bufWriter.WriteString("OK!")
			log.Print(*dt)
		}
		bufWriter.Flush()
	}
}
