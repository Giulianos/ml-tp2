package printTree

import (
	"fmt"

	"github.com/Giulianos/ml-decision-tree/decisiontree"
)

var testExamples = []decisiontree.Example{
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

func main() {
	dt := decisiontree.NewDecisionTree(testExamples, "juega")

	fmt.Print(dt)
}
