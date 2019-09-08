package main

import (
	"bufio"
	"log"
	"net/http"
	"sync"

	"github.com/Giulianos/ml-decision-tree/decisiontree"
)

var examples = []decisiontree.Example{
	{"cielo": "sol", "temp": "calida", "humedad": "normal", "viento": "fuerte", "agua": "calida", "pronos": "estable", "disfruta": "si"},
	{"cielo": "sol", "temp": "calida", "humedad": "alta", "viento": "fuerte", "agua": "calida", "pronos": "estable", "disfruta": "si"},
	{"cielo": "nublado", "temp": "frio", "humedad": "alta", "viento": "fuerte", "agua": "calida", "pronos": "cambiante", "disfruta": "no"},
	{"cielo": "sol", "temp": "calida", "humedad": "alta", "viento": "fuerte", "agua": "fria", "pronos": "cambiante", "disfruta": "si"},
}

func HandleNewTree(dt *decisiontree.DecisionTree, dtMutex *sync.Mutex) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		bufWriter := bufio.NewWriter(writer)

		// Create tree
		tree, err := decisiontree.NewDecisionTree(examples, "disfruta")
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
