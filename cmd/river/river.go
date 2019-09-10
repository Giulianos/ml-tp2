package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Giulianos/ml-decision-tree/decisiontree"
)

func main() {
	// Configure flags
	withNewExample := flag.Bool("with-new", false, "weather or not to include new example")
	flag.Parse()

	examples := []decisiontree.Example{
		{"cielo": "sol", "temp": "calida", "humedad": "normal", "viento": "fuerte", "agua": "calida", "pronos": "estable", "disfruta": "si"},
		{"cielo": "sol", "temp": "calida", "humedad": "alta", "viento": "fuerte", "agua": "calida", "pronos": "estable", "disfruta": "si"},
		{"cielo": "nublado", "temp": "frio", "humedad": "alta", "viento": "fuerte", "agua": "calida", "pronos": "cambiante", "disfruta": "no"},
		{"cielo": "sol", "temp": "calida", "humedad": "alta", "viento": "fuerte", "agua": "fria", "pronos": "cambiante", "disfruta": "si"},
	}

	if *withNewExample {
		newExample := decisiontree.Example{
			"cielo": "sol", "temp": "calida", "humedad": "normal", "viento": "debil", "agua": "calida", "pronos": "estable", "disfruta": "no",
		}
		examples = append(examples, newExample)
	}

	dt, err := decisiontree.NewDecisionTree(examples, "disfruta")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(dt)
}
