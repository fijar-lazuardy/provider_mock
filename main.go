package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"provider_mock/disbursement"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from root!")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"status":"ok"}`)
}

func flipPingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Pong from /flip/ping")
}

func flipEchoHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	fmt.Fprintf(w, "Echo: %s", string(body))
}

func flipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[FLIP] %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func flipDisburse(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	amount := r.FormValue("idempotency-key") // returns string
	fmt.Println(amount)

	resp := disbursement.Inquiry()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity) // optional if 200
	json.NewEncoder(w).Encode(resp)
}

func flipInquiry(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	amount := r.FormValue("idempotency-key") // returns string

	resp := disbursement.Disburse(amount)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity) // optional if 200
	json.NewEncoder(w).Encode(resp)
}

func flipValidateAccount(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	fmt.Fprintf(w, "Echo: %s", string(body))
}

func main() {
	r := mux.NewRouter()

	// Root routes
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/health", healthHandler).Methods("GET")

	// Group: /flip
	flip := r.PathPrefix("/flip").Subrouter()
	flip.Use(flipMiddleware)
	flip.HandleFunc("/ping", flipPingHandler).Methods("GET")
	flip.HandleFunc("/echo", flipEchoHandler).Methods("POST")
	flip.HandleFunc("/echo", flipEchoHandler).Methods("POST")
	flip.HandleFunc("/v2/disbursement/bank-account-inquiry", flipEchoHandler).Methods(http.MethodPost)
	flip.HandleFunc("/v3/special-disbursement", flipDisburse).Methods(http.MethodPost)
	flip.HandleFunc("/v3/get-disbursement", flipDisburse).Methods(http.MethodPost)

	port := "0.0.0.0:3131"
	fmt.Println("Server running on", port)
	log.Fatal(http.ListenAndServe(port, r))
}
