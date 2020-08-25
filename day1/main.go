package main

import (
	_ "database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Solution struct {
	A      int `json:"a"`
	B      int `json:"b"`
	C      int `json:"c"`
	NRoots int `json:"n_roots"`
}

var solutions []Solution

func main() {
	fmt.Println("API was started")

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/solve/{a}/{b}/{c}", AddSolution).Methods("POST")
	router.HandleFunc("/solution", GetLastSolution).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func AddSolution(writer http.ResponseWriter, request *http.Request) {
	var solution Solution
	reqBody := mux.Vars(request)
	solution.A, _ = strconv.Atoi(reqBody["a"])
	solution.B, _ = strconv.Atoi(reqBody["b"])
	solution.C, _ = strconv.Atoi(reqBody["c"])

	CountNRoots(&solution)
	solutions = append(solutions, solution)
	writer.WriteHeader(http.StatusCreated)
}

func CountNRoots(s *Solution) {
	if (s.A == 0 && s.B != 0) || (s.A != 0 && s.C == 0 && s.B == 0) || (s.A == s.B && s.C == 0) {
		s.NRoots = 1
		return
	}
	if s.A == 0 && s.B == 0 {
		s.NRoots = 0
		return
	}

	discr := s.B*s.B - 4*s.A*s.C
	if discr == 0 {
		s.NRoots = 1
	} else if discr > 0 {
		s.NRoots = 2
	} else {
		s.NRoots = 0
	}
}

func GetLastSolution(writer http.ResponseWriter, request *http.Request) {
	if len(solutions) == 0 {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode("Ничего не нашли")
	} else {
		json.NewEncoder(writer).Encode(solutions[len(solutions)-1])
	}
}
