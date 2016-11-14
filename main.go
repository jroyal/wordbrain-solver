package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jroyal/wordbrain-solver/dictionary"
	"github.com/jroyal/wordbrain-solver/grid"
)

func SolveEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var myGrid grid.Grid
	log.Printf("Request Received %v", params)
	_ = json.NewDecoder(req.Body).Decode(&myGrid)
	results := map[int][]string{}
	for _, elem := range myGrid.Words {
		results[elem] = myGrid.GetAllPossibleWords(elem)
	}
	json.NewEncoder(w).Encode(results)
}

func main() {
	router := mux.NewRouter()
	dictionary.LoadDict()
	router.HandleFunc("/solve", SolveEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
