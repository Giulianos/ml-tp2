package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

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

func averageWordCount(examples []classifier.Example) float64 {
	var wordCount float64
	var reviewCount float64
	for _, example := range examples {
		words, err := strconv.ParseInt(example["wordcount"], 10, 64)
		if err == nil {
			wordCount += float64(words)
			reviewCount++
		}
	}
	return wordCount / reviewCount
}

func filter(examples []classifier.Example, keep func(classifier.Example) bool) []classifier.Example {
	filtered := []classifier.Example{}

	for _, example := range examples {
		if keep(example) {
			filtered = append(filtered, example)
		}
	}

	return filtered
}

func main() {
	// Define flags
	dsFilename := flag.String("ds", "", "filename of the dataset (required)")
	flag.Parse()

	// Validate flags
	if *dsFilename == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Read dataset
	ds, err := loadDataset(*dsFilename)
	if err != nil {
		log.Fatal(err)
	}

	// Average word count
	x := averageWordCount(filter(ds, func(ex classifier.Example) bool {
		// Filter only one star reviews
		return ex["Star Rating"] == "1"
	}))

	fmt.Printf("One star reviews have on average, %f words\n", x)

}
