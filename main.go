package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query().Get("domain")
	port := r.URL.Query().Get("port")

	if domain == "" {
		http.Error(w, "Missing 'domain' parameter", http.StatusBadRequest)
		return
	}

	if port == "" {
		port = "80"
	}

	response := Check(domain, port)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", HealthCheckHandler).Methods("GET")

	port := "8080"
	fmt.Printf("Api rodando na porta %s...\n", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor", err)
	}
}
