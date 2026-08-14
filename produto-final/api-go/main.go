package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

// validateTerm confere se todos os campos obrigatórios do termo foram preenchidos
// e se o status é um dos valores permitidos. Retorna uma mensagem de erro, ou "" se estiver tudo certo.
func validateTerm(t Term) string {
	if strings.TrimSpace(t.Termo) == "" {
		return "campo 'termo' é obrigatório"
	}
	if strings.TrimSpace(t.Categoria) == "" {
		return "campo 'categoria' é obrigatório"
	}
	if strings.TrimSpace(t.ExplicacaoSimples) == "" {
		return "campo 'explicacao_simples' é obrigatório"
	}
	if strings.TrimSpace(t.Exemplo) == "" {
		return "campo 'exemplo' é obrigatório"
	}
	if !isValidStatus(t.Status) {
		return "status deve ser 'estudando' ou 'entendido'"
	}
	return ""
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

// corsMiddleware libera o acesso da API para o front-end, que roda em outra origem,
// e responde ao preflight (OPTIONS) que o navegador envia antes de POST/PUT/DELETE
// com corpo JSON.
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
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
	if msg := validateTerm(t); msg != "" {
		writeJSONError(w, http.StatusBadRequest, msg)
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
	if msg := validateTerm(updated); msg != "" {
		writeJSONError(w, http.StatusBadRequest, msg)
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
	http.ListenAndServe(":8080", corsMiddleware(mux))
}
