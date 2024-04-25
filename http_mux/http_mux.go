package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var Pessoas []Pessoa = []Pessoa{}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/lista/pessoas", HandlePessoa).Methods("GET")
	r.HandleFunc("/add/pessoa", HandleAddPessoa).Methods("POST")
	http.ListenAndServe(":8080", r)
}

func HandlePessoa(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(listaPessoas())
}

type Pessoa struct {
	Nome  string `json:"nome"`
	Idade uint8  `json:"idade"`
}

func listaPessoas() []Pessoa {
	pessoas := []Pessoa{
		{
			Nome:  "Nati",
			Idade: 21,
		},
		{
			Nome:  "Emi",
			Idade: 49,
		},
	}
	if len(Pessoas) > 0 {
		pessoas = Pessoas
	}

	return pessoas

}

func HandleAddPessoa(w http.ResponseWriter, rq *http.Request) {
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
	// add a pessoa para a lista
	addPessoa(p)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}

func addPessoa(pessoa Pessoa) {
	Pessoas = append(listaPessoas(), pessoa)
}

// referencia : https://aprendagolang.com.br/?s=mux
// para baixar a lib go get -u github.com/gorilla/mux
