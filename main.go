package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/blevesearch/bleve"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var dbc = sqlx.MustOpen("postgres",
	"postgres://charlie@localhost/KFoodNutrients?sslmode=disable")

func main() {
	// Initializes profiler
	// os.Remove("./cpu.prof")
	// prof, err := os.Create("./cpu.prof")
	// if err != nil {
	// 	panic("could not open cpu profiler")
	// }
	// defer prof.Close()

	// pprof.StartCPUProfile(prof)
	// defer pprof.StopCPUProfile()

	arg := os.Args[1]

	var index bleve.Index

	// Either rebuilds the index or loads it
	if arg == "rebuild" {
		index = initIndex()
	} else {
		index = openIndex()
	}

	// Runs functions in response to commands
	switch arg {
	case "serve":
		serve(index)
	}
}

func serve(index bleve.Index) {
	r := mux.NewRouter()

	r.HandleFunc("/query", makeQueryHandler(index))
	r.HandleFunc("/nutrients", getNutrients)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":4321", nil))
}

func makeQueryHandler(index bleve.Index) func(w http.ResponseWriter,
	r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := r.URL.Query()["q"]

		if len(queries) < 1 {
			http.Error(w, "Must include query", 400)
			return
		}

		searchQuery := queries[0]

		results, err := getSearchResultsJSON(index, searchQuery)
		allowLocalhost(w)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		fmt.Fprintf(w, results)
	}
}

func getNutrients(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()["q"]

	if len(queries) < 1 {
		http.Error(w, "Must include query", 400)
		return
	}

	fdcID := queries[0]

	nutrients, err := getNutrientsJSON(fdcID)
	allowLocalhost(w)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Fprintf(w, nutrients)

}

func allowLocalhost(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
}
