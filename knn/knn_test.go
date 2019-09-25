package knn

import (
	"math"
	"strconv"
	"testing"

	"github.com/Giulianos/ml-tp2/classifier"
)

func TestGetNearest(t *testing.T) {
	knn := New(3, "color", func(ex1, ex2 classifier.Example) float64 {
		height1, _ := strconv.ParseFloat(ex1["height"], 64)
		height2, _ := strconv.ParseFloat(ex2["height"], 64)

		return math.Abs(height1 - height2)
	})

	trainingSet := []classifier.Example{
		{"color": "red", "height": "11.4"},
		{"color": "blue", "height": "100"},
		{"color": "red", "height": "4.3"},
		{"color": "green", "height": "1000"},
		{"color": "red", "height": "15.2"},
		{"color": "red", "height": "17"},
	}

	knn.Fit(trainingSet)

	nearestNeighbors := knn.getKNearest(classifier.Example{"height": "10"})

	for _, neighbor := range nearestNeighbors {
		if neighbor.Distance > 6 {
			t.Errorf("distance should be at most 5.7, but got %f for %s", neighbor.Distance, neighbor.Neighbor)
		}
	}
}

func TestClassifyNotWeighted(t *testing.T) {
	knn := New(5, "color", func(ex1, ex2 classifier.Example) float64 {
		height1, _ := strconv.ParseFloat(ex1["height"], 64)
		height2, _ := strconv.ParseFloat(ex2["height"], 64)

		return math.Abs(height1 - height2)
	})
	knn.SetWeighted(false)

	trainingSet := []classifier.Example{
		{"color": "red", "height": "11.4"},
		{"color": "red", "height": "9.3"},
		{"color": "green", "height": "100"},
		{"color": "blue", "height": "4.3"},
		{"color": "green", "height": "150"},
		{"color": "blue", "height": "15.2"},
		{"color": "blue", "height": "8"},
	}

	knn.Fit(trainingSet)

	class, _ := knn.Classify(classifier.Example{"height": "10"})

	if class != "blue" {
		t.Errorf("class should be red, got %s", class)
	}

}

func TestClassifyWeighted(t *testing.T) {
	knn := New(5, "color", func(ex1, ex2 classifier.Example) float64 {
		height1, _ := strconv.ParseFloat(ex1["height"], 64)
		height2, _ := strconv.ParseFloat(ex2["height"], 64)

		return math.Abs(height1 - height2)
	})
	knn.SetWeighted(true)

	trainingSet := []classifier.Example{
		{"color": "red", "height": "11.4"},
		{"color": "red", "height": "9.3"},
		{"color": "green", "height": "100"},
		{"color": "blue", "height": "4.3"},
		{"color": "green", "height": "150"},
		{"color": "blue", "height": "15.2"},
		{"color": "blue", "height": "8"},
	}

	knn.Fit(trainingSet)

	class, _ := knn.Classify(classifier.Example{"height": "10"})

	if class != "red" {
		t.Errorf("class should be red, got %s", class)
	}

}

func TestGetClasses(t *testing.T) {
	knn := New(5, "color", func(ex1, ex2 classifier.Example) float64 {
		height1, _ := strconv.ParseFloat(ex1["height"], 64)
		height2, _ := strconv.ParseFloat(ex2["height"], 64)

		return math.Abs(height1 - height2)
	})
	knn.SetWeighted(true)

	trainingSet := []classifier.Example{
		{"color": "red", "height": "11.4"},
		{"color": "red", "height": "9.3"},
		{"color": "green", "height": "100"},
		{"color": "blue", "height": "4.3"},
		{"color": "green", "height": "150"},
		{"color": "blue", "height": "15.2"},
		{"color": "blue", "height": "8"},
	}

	knn.Fit(trainingSet)

	classes := knn.GetClasses()

	if len(classes) != 3 {
		t.Errorf("there should be 3 classes, got %d", len(classes))
	}

	// TODO: fix this test to look for all the classes to be in the slice (now they can be repeated)
	for _, class := range classes {
		if class != "red" && class != "blue" && class != "green" {
			t.Errorf("unexpected class")
		}
	}

}
