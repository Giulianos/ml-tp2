package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/Giulianos/ml-decision-tree/classifier"

	"github.com/Giulianos/ml-decision-tree/decisiontree"
)

func csvToExamples(r io.Reader) ([]classifier.Example, error) {
	csvReader := csv.NewReader(r)

	attrs, err := csvReader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading dataset")
	}

	var examples []classifier.Example

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		newExample := classifier.Example{}
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
		log.Print(examples)

		// Get predicted attribute
		predAttr := request.URL.Query()["pred-attr"][0]
		log.Println(predAttr)

		// Create tree
		tree := decisiontree.NewDecisionTree(predAttr)
		tree.SetGainFunction(decisiontree.GINI)
		tree.SetMaxSplits(3)
		err = tree.Fit(examples)

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
