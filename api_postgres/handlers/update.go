package handlers

import (
	"api_postgres/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	var todo models.Todo
	if err != nil {
		log.Printf("ERRO AO TENTAR FAZER O PARSE DO ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("ERRO AO FAZER DECODE DO JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Update(int64(id), todo)
	if err != nil {
		log.Printf("ERRO AO ATUALIZAR REGISTRO: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if rows > 1 {
		log.Printf("Error: Foram atualizados : %d registros", rows)
	}
	resp := map[string]any{
		"Error":   false,
		"Message": "dados atualizados com sucesso",
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
