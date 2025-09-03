package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"provider_mock/disbursement"
	"strconv"

	"github.com/Pallinder/go-randomdata"
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

func flipInquiry(w http.ResponseWriter, r *http.Request) {
	amount := r.URL.Query().Get("idempotency-key")
	fmt.Println(amount)

	resp := disbursement.Inquiry()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity) // optional if 200
	json.NewEncoder(w).Encode(resp)
}

func flipDisburse(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	amount := r.FormValue("idempotency-key") // returns string

	resp := disbursement.Disburse(amount)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError) // optional if 200
	json.NewEncoder(w).Encode(resp)
}

func flipValidateAccount(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"bank_code":          "bri",
		"account_number":     "0013000397",
		"account_holder":     "Dummy Name",
		"status":             "SUCCESS",
		"inquiry_key":        "",
		"is_virtual_account": "false", // must be string, since map is map[string]string
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // optional if 200
	json.NewEncoder(w).Encode(data)

}

func jackValidateAccount(w http.ResponseWriter, r *http.Request) {
	account_number := r.URL.Query().Get("account_number")
	bank_name := r.URL.Query().Get("bank_name")
	payload := map[string]any{
		"status": 200,
		"data": map[string]any{
			"id":           randomdata.Alphanumeric(20), // or generate your own
			"account_no":   account_number,
			"bank_name":    bank_name,
			"account_name": randomdata.FullName(randomdata.RandomGender),
		},
	}

	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	// If you want to return it:
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func jackDisbursement(w http.ResponseWriter, r *http.Request) {
	payload := map[string]any{
		"status": 2,
		"state":  6,
	}
	jsonBytes, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError) // optional if 200
	json.NewEncoder(w).Encode(jsonBytes)
}

func simulateHandle(w http.ResponseWriter, r *http.Request) {
	ref_id := r.URL.Query().Get("reference_id")
	amount := r.URL.Query().Get("amount")
	amountInt, err := strconv.Atoi(amount)
	if err != nil {
		return
	}

	baseUrl := "https://rest.doitpay.dev/checkout/v1/public/simulate"
	bodyRequest := map[string]any{
		"amount":       amountInt,
		"reference_id": ref_id,
		"payment_type": "direct_debit",
	}
	bodyJson, _ := json.Marshal(bodyRequest)
	httpReq, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, baseUrl, bytes.NewReader(bodyJson))
	if err != nil {
		return
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}
	httpClient.Do(httpReq)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // optional if 200
	payload := map[string]any{
		"status": "success",
	}
	json.NewEncoder(w).Encode(payload)

}

func main() {
	r := mux.NewRouter()

	// Root routes
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/health", healthHandler).Methods("GET")
	r.HandleFunc("/simulate", simulateHandle).Methods(http.MethodGet)

	// Group: /flip
	flip := r.PathPrefix("/flip").Subrouter()
	flip.Use(flipMiddleware)
	flip.HandleFunc("/ping", flipPingHandler).Methods("GET")
	flip.HandleFunc("/echo", flipEchoHandler).Methods("POST")
	flip.HandleFunc("/echo", flipEchoHandler).Methods("POST")
	flip.HandleFunc("/v2/disbursement/bank-account-inquiry", flipValidateAccount).Methods(http.MethodPost)
	flip.HandleFunc("/v3/special-disbursement", flipDisburse).Methods(http.MethodPost)
	flip.HandleFunc("/v3/get-disbursement", flipInquiry).Methods(http.MethodGet)
	flip.HandleFunc("/v2/disbursement/bank-account-inquiry", flipValidateAccount).Methods(http.MethodGet)

	jack := r.PathPrefix("/jack").Subrouter()
	jack.HandleFunc("/validation_bank_account", jackValidateAccount).Methods(http.MethodGet)
	jack.HandleFunc("/transactions", jackDisbursement).Methods(http.MethodPost)

	port := "0.0.0.0:313"

	fmt.Println("Server is running!!!! on", port)

	log.Fatal(http.ListenAndServe(port, r))
}
