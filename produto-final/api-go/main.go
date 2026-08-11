package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Term representa um termo técnico do glossário.
type Term struct {
	ID                int    `json:"id"`
	Termo             string `json:"termo"`
	Categoria         string `json:"categoria"`
	ExplicacaoSimples string `json:"explicacao_simples"`
	Exemplo           string `json:"exemplo"`
	Status            string `json:"status"`
}

// Armazenamento em memória: os termos ficam guardados aqui enquanto o servidor roda.
var terms = []Term{}
var nextID = 1

func isValidStatus(status string) bool {
	return status == "estudando" || status == "entendido"
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func writeJSONError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"erro": message})
}

// GET /terms
func listTerms(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, terms)
}

// GET /terms/{id}
func getTerm(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "id inválido")
		return
	}
	for _, t := range terms {
		if t.ID == id {
			writeJSON(w, http.StatusOK, t)
			return
		}
	}
	writeJSONError(w, http.StatusNotFound, "termo não encontrado")
}

// POST /terms
func createTerm(w http.ResponseWriter, r *http.Request) {
	var t Term
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		writeJSONError(w, http.StatusBadRequest, "JSON inválido")
		return
	}
	if !isValidStatus(t.Status) {
		writeJSONError(w, http.StatusBadRequest, "status deve ser 'estudando' ou 'entendido'")
		return
	}

	t.ID = nextID
	nextID++
	terms = append(terms, t)

	writeJSON(w, http.StatusCreated, t)
}

// PUT /terms/{id}
func updateTerm(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "id inválido")
		return
	}

	var updated Term
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		writeJSONError(w, http.StatusBadRequest, "JSON inválido")
		return
	}
	if !isValidStatus(updated.Status) {
		writeJSONError(w, http.StatusBadRequest, "status deve ser 'estudando' ou 'entendido'")
		return
	}

	for i, t := range terms {
		if t.ID == id {
			updated.ID = id
			terms[i] = updated
			writeJSON(w, http.StatusOK, updated)
			return
		}
	}
	writeJSONError(w, http.StatusNotFound, "termo não encontrado")
}

// DELETE /terms/{id}
func deleteTerm(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "id inválido")
		return
	}

	for i, t := range terms {
		if t.ID == id {
			terms = append(terms[:i], terms[i+1:]...)
			writeJSON(w, http.StatusOK, map[string]string{"mensagem": "termo removido"})
			return
		}
	}
	writeJSONError(w, http.StatusNotFound, "termo não encontrado")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /terms", listTerms)
	mux.HandleFunc("GET /terms/{id}", getTerm)
	mux.HandleFunc("POST /terms", createTerm)
	mux.HandleFunc("PUT /terms/{id}", updateTerm)
	mux.HandleFunc("DELETE /terms/{id}", deleteTerm)

	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
