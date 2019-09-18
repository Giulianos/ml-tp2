package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"log"
	"os"

	"github.com/Giulianos/ml-decision-tree/classifier"

	"github.com/Giulianos/ml-decision-tree/marshalling"
)

func randomSplit(examples []classifier.Example, testPortion float64) ([]classifier.Example, []classifier.Example) {
	// TODO: implement method
	return nil, nil
}

func writeSet(examples []classifier.Example, filename string) error {
	// Open file to write
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)
	csvWriter := csv.NewWriter(writer)
	marshalling.MarshallCSV(examples, *csvWriter)
	err = writer.Flush()
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Read flags
	dsFilename := flag.String("ds", "", "input filename for dataset to split, stdin if empty")
	trainFilename := flag.String("ds", "train.csv", "output filename for train set")
	testFilename := flag.String("ds", "test.csv", "output filename for test set")
	flag.Parse()

	// Open file to read
	dsFile, err := os.Open(*dsFilename)
	if err != nil {
		log.Fatal(err)
	}
	dsReader := bufio.NewReader(dsFile)
	dsCSVReader := csv.NewReader(dsReader)

	// Load entire dataset into memory (naive implementation for relatively small files)
	examples, err := marshalling.UnmarshallCSV(*dsCSVReader)
	if err != nil {
		log.Fatal(err)
	}

	// Split set in training and test
	training, test := randomSplit(examples, 0.1)

	// Save files
	err = writeSet(training, *trainFilename)
	if err != nil {
		log.Fatal(err)
	}
	err = writeSet(test, *testFilename)
	if err != nil {
		log.Fatal(err)
	}
}
