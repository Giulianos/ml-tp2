package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Giulianos/ml-decision-tree/decisiontree"

	"github.com/Giulianos/ml-decision-tree/classifier"
	"github.com/Giulianos/ml-decision-tree/marshalling"
)

func loadDataset(filename string) ([]classifier.Example, error) {
	// Open file
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	// Create csv reader
	csvReader := csv.NewReader(f)
	examples, err := marshalling.UnmarshallCSV(*csvReader)
	if err != nil {
		return nil, err
	}
	return examples, nil
}

func main() {
	// Read flags
	trainFilename := flag.String("train", "", "input training set filename (required)")
	testFilename := flag.String("test", "", "input test set filename (required)")
	predictableAttr := flag.String("pred-attr", "", "column name of the predictable attribute (required)")
	flag.Parse()

	// Validate parameters
	if *trainFilename == "" || *testFilename == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Load training set
	training, err := loadDataset(*trainFilename)
	if err != nil {
		log.Fatal(err)
	}

	// Load test set
	test, err := loadDataset(*testFilename)
	if err != nil {
		log.Fatal(err)
	}

	// Create classifier
	classif := decisiontree.NewDecisionTree(*predictableAttr)
	/*
		Configure classifier
		classif.SetGainFunction(decisiontree.SHANNON_ENTROPY)
		classif.SetMinSplitCount(10)
	*/
	// Train classifier
	err = classif.Fit(training)
	if err != nil {
		log.Fatal(err)
	}

	// Print tree
	fmt.Println("Using the following tree")
	fmt.Println(classif)

	// Run evaluation
	eval := classifier.EvalClassifier(classif, test)

	// Print confusion matrix
	fmt.Println(eval.ConfusionMatrixToString())
}
