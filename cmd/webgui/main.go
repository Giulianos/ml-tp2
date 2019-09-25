package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"

	"github.com/Giulianos/ml-tp2/decisiontree"
	"github.com/gorilla/handlers"
)

var dt decisiontree.DecisionTree
var dtMutex sync.Mutex

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/tree", HandleNewTree(&dt, &dtMutex)).Methods("POST")
	r.HandleFunc("/graph", HandleGetGraph(&dt)).Methods("GET")

	corsHandler := handlers.CORS(
		handlers.AllowedHeaders([]string{
			"X-Requested-With",
			"Content-Type",
			"Authorization",
		}), handlers.AllowedMethods([]string{
			"GET",
			"POST",
			"PUT",
			"HEAD",
			"OPTIONS",
		}), handlers.AllowedOrigins([]string{"*"}))(r)

	srv := &http.Server{
		Handler:      corsHandler,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
