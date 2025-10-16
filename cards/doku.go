package cards

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"provider_mock/dto"
	"text/template"
)

func GenerateDokuCardPaymentPage(ctx context.Context, request dto.CreateCreditCardPaymentRequest) (response dto.CreateCreditCardPaymentResponse, err error) {
	callbackUrl := base64.URLEncoding.EncodeToString([]byte(request.Order.CallbackURL))
	response = dto.CreateCreditCardPaymentResponse{
		Order: dto.Order{
			InvoiceNumber: request.Order.InvoiceNumber,
			Amount:        request.Order.Amount,
			CallbackURL:   request.Order.CallbackURL,
			FailedURL:     "",
			AutoRedirect:  false,
			Descriptor:    "",
			SessionID:     "",
		},
		CreditCardPaymentPage: dto.CreditCardPaymentPage{
			URL: fmt.Sprintf("https://provider.lazu.dev/doku/payment-page?invoice_number=%s&callback_url=%s", request.Order.InvoiceNumber, callbackUrl),
		},
		AdditionalInfo: dto.AdditionalInfo{},
	}
	return
}

// form.html is in the same directory as this file
func DokuRenderFormPage(ctx context.Context, invoiceNumber string, callbackUrl string) (string, error) {
	// Read the HTML template file
	data, err := os.ReadFile("form.html")
	if err != nil {
		return "", fmt.Errorf("failed to read form.html: %w", err)
	}

	// Parse the template
	tmpl, err := template.New("form").Parse(string(data))
	if err != nil {
		return "", fmt.Errorf("failed to parse form.html: %w", err)
	}

	// Inject dynamic data (like invoice number)
	var renderedHTML string
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, map[string]any{
		"InvoiceNumber": invoiceNumber,
		"CallbackURL":   callbackUrl,
	})
	if err != nil {
		return "", fmt.Errorf("failed to execute form.html: %w", err)
	}

	renderedHTML = buf.String()
	return renderedHTML, nil
}

func DokuPaymentAndSendCallback(ctx context.Context, request dto.DokuPayment) (err error) {
	httpClient := &http.Client{
		Timeout: 5,
	}

	callbackUrl, err := base64.URLEncoding.DecodeString(request.CallbackURL)

	body := dto.DokuCreditCardNotifyPaymentRequest{
		Order: dto.Order{
			InvoiceNumber: request.InvoiceNumber,
			CallbackURL:   request.CallbackURL,
			Amount:        request.Amount,
		},
	}
	jsonBody, _ := json.Marshal(body)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, string(callbackUrl), bytes.NewBuffer(jsonBody))
	httpClient.Do(req)
	return
}
