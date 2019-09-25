package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/Giulianos/ml-tp2/knn"

	"github.com/Giulianos/ml-tp2/classifier"
	"github.com/Giulianos/ml-tp2/marshalling"
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

func exampleToTuple(ex classifier.Example) (wordCount, titleSentiment, sentimentValue float64) {
	wordCount, err := strconv.ParseFloat(ex["wordcount"], 64)
	if err != nil {
		log.Fatal(err)
	}
	titleSentiment, err = strconv.ParseFloat(ex["titleSentiment"], 64)
	if err != nil {
		log.Fatal(err)
	}
	sentimentValue, err = strconv.ParseFloat(ex["sentimentValue"], 64)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func euclideanDistance(ex1, ex2 classifier.Example) float64 {
	wc1, ts1, sv1 := exampleToTuple(ex1)
	wc2, ts2, sv2 := exampleToTuple(ex2)

	return math.Sqrt(math.Pow(wc1-wc2, 2) + math.Pow(ts1-ts2, 2) + math.Pow(sv1-sv2, 2))
}

func customDistance(ex1, ex2 classifier.Example) float64 {
	wc1, ts1, sv1 := exampleToTuple(ex1)
	wc2, ts2, sv2 := exampleToTuple(ex2)

	return math.Sqrt(math.Pow(wc1-wc2, 2) + math.Pow(2*(ts1-ts2), 2) + math.Pow(3*(sv1-sv2), 2))
}

func main() {
	// Parse flags
	trainFilename := flag.String("train", "", "training set filename (required)")
	testFilename := flag.String("test", "", "test set filename (required)")
	weighted := flag.Bool("weighted", false, "whether to weight contributions with the distance")
	flag.Parse()

	// Validate flags
	if *trainFilename == "" || *testFilename == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Load training set
	trainingSet, err := loadDataset(*trainFilename)
	if err != nil {
		log.Fatal(err)
	}

	// Create KNN classifier
	classif := knn.New(5, "StarRating", customDistance)
	classif.SetWeighted(*weighted)

	// Fit classifier
	classif.Fit(trainingSet)

	// Load test set
	testSet, err := loadDataset(*testFilename)
	if err != nil {
		log.Fatal(err)
	}

	eval := classifier.EvalClassifier(classif, testSet)

	fmt.Println(eval)
}
