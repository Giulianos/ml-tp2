package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/Giulianos/ml-decision-tree/decisiontree"
	"github.com/gorilla/mux"
)

var dt decisiontree.DecisionTree
var dtMutex sync.Mutex

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/tree", HandleNewTree(&dt, &dtMutex)).Methods("POST")
	r.HandleFunc("/graph", HandleGetGraph(&dt)).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
