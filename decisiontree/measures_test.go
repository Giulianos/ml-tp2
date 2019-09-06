package decisiontree

import (
	"math"
	"testing"
)

var testExamples = []Example{
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

func TestSomething(t *testing.T) {
	dt := NewDecisionTree(testExamples, "juega")

	expectedEntropy := -(9/14)*math.Log2(9/14) - (5/14)*math.Log2(5/14)
	actualEntropy := dt.sEntropy(testExamples)

	if math.Abs(expectedEntropy-actualEntropy) > 0.001 {
		t.Errorf("Expected: %f, got: %f", expectedEntropy, actualEntropy)
	}
}
