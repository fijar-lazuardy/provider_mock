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
	"provider_mock/cards"
	"provider_mock/disbursement"
	"provider_mock/dto"
	"strconv"
	"time"

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

func nobuTransfer(w http.ResponseWriter, r *http.Request) {
	customResponse := r.Header.Get("CustomResponse")
	data := map[string]interface{}{}

	switch customResponse {
	case "TIMEOUT":
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusGatewayTimeout)
		json.NewEncoder(w).Encode(data)

	case "PROCESSING":
		data = map[string]interface{}{
			"additionalInfo": map[string]interface{}{
				"beneficiaryAccountName":   "TEST 4",
				"beneficiaryAccountStatus": "01",
				"beneficiaryAccountType":   "CACC",
				"currency":                 "IDR",
				"customerReference":        fmt.Sprint(randomdata.Number(12)),
				"transactionDate":          time.Now().Format("2006-01-02T15:04:05-07:00"),
			},
			"amount":               map[string]string{"currency": "IDR", "value": "123456.00"},
			"beneficiaryAccountNo": "510654304",
			"beneficiaryBankCode":  "SIHBIDJ1",
			"partnerReferenceNo":   "",
			"referenceNo":          fmt.Sprint(randomdata.Number(12)),
			"responseCode":         "2021800",
			"responseMessage":      "Request has been processed successfully",
			"sourceAccountNo":      "10110889307",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(data)

	default:
		data = map[string]interface{}{
			"additionalInfo": map[string]interface{}{
				"beneficiaryAccountName":   "TEST 4",
				"beneficiaryAccountStatus": "01",
				"beneficiaryAccountType":   "CACC",
				"currency":                 "IDR",
				"customerReference":        fmt.Sprint(randomdata.Number(12)),
				"transactionDate":          time.Now().Format("2006-01-02T15:04:05-07:00"),
			},
			"amount":               map[string]string{"currency": "IDR", "value": "123456.00"},
			"beneficiaryAccountNo": "510654304",
			"beneficiaryBankCode":  "SIHBIDJ1",
			"partnerReferenceNo":   "",
			"referenceNo":          fmt.Sprint(randomdata.Number(12)),
			"responseCode":         "2001800",
			"responseMessage":      "Request has been processed successfully",
			"sourceAccountNo":      "10110889307",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}

func nobuCheckStatus(w http.ResponseWriter, r *http.Request) {
	customResponse := r.Header.Get("CustomResponse")
	data := map[string]interface{}{}

	switch customResponse {
	case "TIMEOUT":
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusGatewayTimeout)
		json.NewEncoder(w).Encode(data)

	case "NOT FOUND":
		data = map[string]interface{}{
			"responseCode":    "4043601",
			"responseMessage": "Transaction Not Found",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(data)

	default:
		data = map[string]interface{}{
			"additionalInfo": map[string]interface{}{
				"beneficiaryAccountName": "An***n",
				"referenceNo":            "20220818LFIBIDJ1010O0200001841",
			},
			"amount":                     map[string]interface{}{"currency": "IDR", "value": "100001.00"},
			"beneficiaryAccountNo":       "51******1",
			"beneficiaryBankCode":        "GN*****A",
			"latestTransactionStatus":    "00",
			"originalPartnerReferenceNo": "202507041112213258216820887277",
			"originalReferenceNo":        "220818001861",
			"referenceNumber":            "20220818LFIBIDJ1010O0200001841",
			"responseCode":               "2003600",
			"responseMessage":            "Request has been processed successfully",
			"serviceCode":                "18",
			"sourceAccountNo":            "10********3",
			"transactionStatusDesc":      "Success",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
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

func createCCDoku(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateCreditCardPaymentRequest

	defer r.Body.Close()
	bodyReq, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(bodyReq, &request)
	if err != nil {
		return
	}
	resp, err := cards.GenerateDokuCardPaymentPage(r.Context(), request)
	if err != nil {
		return
	}

	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // optional if 200
	w.Write(jsonBytes)
}

func showDokuCCPage(w http.ResponseWriter, r *http.Request) {
	invoiceNumber := r.URL.Query().Get("invoice_number")
	callbackUrl := r.URL.Query().Get("callback_url")

	renderedHTML, err := cards.DokuRenderFormPage(r.Context(), invoiceNumber, callbackUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(renderedHTML))
}

func dokuProcessCCPayment(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	invoiceNumber := r.FormValue("invoice_number")
	callbackUrl := r.FormValue("callback_url")
	order := int64(10000)

	cards.DokuPaymentAndSendCallback(r.Context(), dto.DokuPayment{
		Amount:        order,
		InvoiceNumber: invoiceNumber,
		CallbackURL:   callbackUrl,
	})

	// ccNumber := r.FormValue("ccNumber")
	// exp := r.FormValue("exp")
	// cvv := r.FormValue("cvv")
	// jsonBytes, err := json.Marshal(resp)
	// if err != nil {
	// 	return
	// }
	w.WriteHeader(http.StatusOK)
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

	nobu := r.PathPrefix("/nobu").Subrouter()
	nobu.HandleFunc("/v1.3/transfer-interbank/", nobuTransfer).Methods(http.MethodPost)
	nobu.HandleFunc("/v1.1/transfer/status/", nobuCheckStatus).Methods(http.MethodPost)

	jack := r.PathPrefix("/jack").Subrouter()
	jack.HandleFunc("/validation_bank_account", jackValidateAccount).Methods(http.MethodGet)
	jack.HandleFunc("/transactions", jackDisbursement).Methods(http.MethodPost)

	doku := r.PathPrefix("/doku").Subrouter()
	doku.HandleFunc("/credit-card/v1/payment-page", createCCDoku).Methods(http.MethodPost)
	doku.HandleFunc("/payment-page", showDokuCCPage).Methods(http.MethodGet)
	doku.HandleFunc("/process-payment", dokuProcessCCPayment).Methods(http.MethodPost)

	port := "0.0.0.0:3131"

	fmt.Println("Server is running!!!! on", port)

	log.Fatal(http.ListenAndServe(port, r))
}
