package main

import (
	"encoding/json"
	"net/http"
)

type Pessoa struct {
	Nome  string `json:"nome"`
	Idade uint8  `json:"idade"`
}

func main() {

	http.HandleFunc("/pessoa", HandlePerson)
	http.ListenAndServe(":8080", nil)
}

func HandlePerson(w http.ResponseWriter, rq *http.Request) {

	// Permitido apenas POST requests
	if rq.Method != "POST" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var p Pessoa
	err := json.NewDecoder(rq.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(p)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	defer rq.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)

}
