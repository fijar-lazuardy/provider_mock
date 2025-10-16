package dto

import "time"

type CreateCreditCardPaymentRequest struct {
	Order                 Order                  `json:"order"`
	Card                  *Card                  `json:"card,omitempty"`
	Customer              *Customer              `json:"customer,omitempty"`
	Payment               Payment                `json:"payment"`
	OverrideConfiguration *OverrideConfiguration `json:"override_configuration,omitempty"`
	AdditionalInfo        AdditionalInfo         `json:"additional_info"`
}

type CreateCreditCardPaymentResponse struct {
	Order                 Order                 `json:"order"`
	CreditCardPaymentPage CreditCardPaymentPage `json:"credit_card_payment_page"`
	AdditionalInfo        AdditionalInfo        `json:"additional_info"`
}

type CreditCardPaymentPage struct {
	URL string `json:"url"`
}

type Order struct {
	InvoiceNumber string `json:"invoice_number"`
	Amount        int64  `json:"amount,omitempty"`
	CallbackURL   string `json:"callback_url"`
	FailedURL     string `json:"failed_url"`
	AutoRedirect  bool   `json:"auto_redirect"`
	Descriptor    string `json:"descriptor,omitempty"`
	SessionID     string `json:"session_id,omitempty"`
}

type Card struct {
	Token string `json:"token"`
	Save  bool   `json:"save"`
}

type Customer struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Country string `json:"country"`
}

type Payment struct {
	Type              string `json:"type,omitempty"`
	Acquirer          string `json:"acquirer,omitempty"`
	Tenor             int    `json:"tenor,omitempty"`
	OriginalRequestID string `json:"original_request_id,omitempty"`
}

type OverrideConfiguration struct {
	Themes     *Themes   `json:"themes,omitempty"`
	Promo      []Promo   `json:"promo"`
	AllowBin   *[]string `json:"allow_bin,omitempty"`
	AllowTenor *[]int    `json:"allow_tenor,omitempty"`
}

type Themes struct {
	Language              string `json:"language"`
	BackgroundColor       string `json:"background_color"`
	FontColor             string `json:"font_color"`
	ButtonBackgroundColor string `json:"button_background_color"`
	ButtonFontColor       string `json:"button_font_color"`
}

type Promo struct {
	Bin            string `json:"bin"`
	DiscountAmount int    `json:"discount_amount"`
}

type AdditionalInfo struct {
	OverrideNotificationURL string      `json:"override_notification_url"`
	Disclaimer              *Disclaimer `json:"disclaimer,omitempty"`
}

type Disclaimer struct {
	ID string `json:"id"`
	EN string `json:"en"`
}

type CheckCreditCardPaymentStatusResponse struct {
	Order       Order       `json:"order"`
	Transaction Transaction `json:"transaction"`
	Service     Service     `json:"service"`
	Acquirer    Acquirer    `json:"acquirer"`
	Channel     Channel     `json:"channel"`
	CardPayment CardPayment `json:"card_payment"`
}

type Transaction struct {
	Status            string    `json:"status"`
	Date              time.Time `json:"date"`
	Type              string    `json:"type"`
	OriginalRequestID string    `json:"original_request_id"`
}

type Service struct {
	ID string `json:"id"`
}

type Acquirer struct {
	ID string `json:"id"`
}

type Channel struct {
	ID string `json:"id"`
}

type CardPayment struct {
	CardMasked           string    `json:"card_masked"`
	ApprovalCode         string    `json:"approval_code"`
	ResponseCode         string    `json:"response_code"`
	ResponseMessage      string    `json:"response_message"`
	Type                 string    `json:"type"`
	AcquiringOffUsStatus string    `json:"acquiring_off_us_status"`
	RequestID            string    `json:"request_id"`
	CardType             string    `json:"card_type"`
	ThreeDSecureStatus   string    `json:"three_dsecure_status"`
	Issuer               string    `json:"issuer"`
	TransactionStatus    string    `json:"transaction_status"`
	Brand                string    `json:"brand"`
	Date                 time.Time `json:"date"`
}

type RefundCreditCardPaymentRequest struct {
	Order   Order   `json:"order"`
	Payment Payment `json:"payment"`
	Refund  Refund  `json:"refund"`
}

type RefundCreditCardPaymentResponse struct {
	Order   Order   `json:"order"`
	Payment Payment `json:"payment"`
	Refund  Refund  `json:"refund"`
}

type Refund struct {
	Amount       int     `json:"amount"`
	Type         *string `json:"type,omitempty"`
	Status       *string `json:"status,omitempty"`
	Message      *string `json:"message,omitempty"`
	ApprovalCode *string `json:"approval_code,omitempty"`
}

type DokuPayment struct {
	Amount        int64
	InvoiceNumber string
	CallbackURL   string
}

type DokuCreditCardNotifyPaymentRequest struct {
	Order       Order       `json:"order"`
	Customer    Customer    `json:"customer"`
	Transaction Transaction `json:"transaction"`
	Service     Service     `json:"service"`
	Acquirer    Acquirer    `json:"acquirer"`
	Channel     Channel     `json:"channel"`
	CardPayment CardPayment `json:"card_payment"`
	AuthorizeID string      `json:"authorize_id"`
}
