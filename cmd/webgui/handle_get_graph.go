package main

import (
	"bufio"
	"net/http"

	"github.com/Giulianos/ml-decision-tree/decisiontree"
)

func HandleGetGraph(dt *decisiontree.DecisionTree) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		bufWriter := bufio.NewWriter(writer)

		if !dt.Built {
			writer.WriteHeader(400)
			bufWriter.WriteString("Tree is not built")
		} else {
			writer.WriteHeader(200)
			writer.Header().Set("Content-Type", "text/vnd.graphviz")
			bufWriter.WriteString(dt.String())
		}
		bufWriter.Flush()
	}
}
