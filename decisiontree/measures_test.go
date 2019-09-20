package decisiontree

import (
	"math"
	"testing"

	"github.com/Giulianos/ml-decision-tree/classifier"
)

var testExamples = []classifier.Example{
	{"pronostico": "soleado", "temp": "calido", "humedad": "alta", "viento": "debil", "juega": "no"},
	{"pronostico": "soleado", "temp": "calido", "humedad": "alta", "viento": "fuerte", "juega": "no"},
	{"pronostico": "nublado", "temp": "calido", "humedad": "alta", "viento": "debil", "juega": "si"},
	{"pronostico": "lluvioso", "temp": "templado", "humedad": "alta", "viento": "debil", "juega": "si"},
	{"pronostico": "lluvioso", "temp": "frio", "humedad": "normal", "viento": "debil", "juega": "si"},
	{"pronostico": "lluvioso", "temp": "frio", "humedad": "normal", "viento": "fuerte", "juega": "no"},
	{"pronostico": "nublado", "temp": "frio", "humedad": "normal", "viento": "fuerte", "juega": "si"},
	{"pronostico": "soleado", "temp": "templado", "humedad": "alta", "viento": "debil", "juega": "no"},
	{"pronostico": "soleado", "temp": "frio", "humedad": "normal", "viento": "debil", "juega": "si"},
	{"pronostico": "lluvioso", "temp": "templado", "humedad": "normal", "viento": "debil", "juega": "si"},
	{"pronostico": "soleado", "temp": "templado", "humedad": "normal", "viento": "fuerte", "juega": "si"},
	{"pronostico": "nublado", "temp": "templado", "humedad": "alta", "viento": "fuerte", "juega": "si"},
	{"pronostico": "nublado", "temp": "calido", "humedad": "normal", "viento": "debil", "juega": "si"},
	{"pronostico": "lluvioso", "temp": "templado", "humedad": "alta", "viento": "fuerte", "juega": "no"},
}

func TestEntropy(t *testing.T) {
	dt := NewDecisionTree("juega")

	err := dt.Fit(testExamples)

	if err != nil {
		t.Errorf("error fitting: %e", err)
	}

	expectedEntropy := -(9./14.)*math.Log2(9./14.) - (5./14.)*math.Log2(5./14.)
	actualEntropy := dt.sEntropy(testExamples)

	if math.Abs(expectedEntropy-actualEntropy) > 0.0001 {
		t.Errorf("expected: %f, got: %f", expectedEntropy, actualEntropy)
	}
}

func TestSVEntropy(t *testing.T) {
	dt := NewDecisionTree("juega")

	err := dt.Fit(testExamples)

	if err != nil {
		t.Errorf("error fitting: %e", err)
	}

	expectedEntropy := -(6./8.)*math.Log2(6./8.) - (2./8.)*math.Log2(2./8.)
	actualEntropy, _ := dt.svEntropy(testExamples, "viento", "debil")

	if math.Abs(expectedEntropy-actualEntropy) > 0.0001 {
		t.Errorf("expected: %f, got: %f", expectedEntropy, actualEntropy)
	}
}

func TestGain(t *testing.T) {
	dt := NewDecisionTree("juega")

	err := dt.Fit(testExamples)

	if err != nil {
		t.Errorf("error fitting: %e", err)
	}

	// Expected discriminant attribute
	expDiscAttr := "pronostico"

	// Find discriminant attribute
	discAttr := ""
	var discAttrGain float64
	for attr := range dt.domain {
		if attr == dt.predAttr {
			continue
		}
		gain := dt.gain(testExamples, attr)
		if gain > discAttrGain {
			discAttr = attr
			discAttrGain = gain
		}
	}

	if discAttr != expDiscAttr {
		t.Errorf("Expected: %s, got: %s", expDiscAttr, discAttr)
	}
}
