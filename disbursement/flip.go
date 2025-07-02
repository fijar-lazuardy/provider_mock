package disbursement

import (
	"errors"
	"provider_mock/dto"

	"github.com/Pallinder/go-randomdata"
)

func ValidateAccount(AccountNumber string) (AccountName string, err error) {
	if AccountNumber == "11122233444" {
		return "", errors.New("error")
	}

	AccountName = randomdata.FullName(randomdata.RandomGender)
	return
}

func Disburse(Amount string) (response dto.FlipErrorResponse) {
	switch Amount {
	case "99900":
		return dto.FlipErrorResponse{Code: dto.FlipErrUndefined, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrUndefined, Message: ""}}}
	case "100100":
		return dto.FlipErrorResponse{Code: dto.FlipErrRequiredAttribute, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrRequiredAttribute, Message: ""}}}
	case "100200":
		return dto.FlipErrorResponse{Code: dto.FlipErrUncleanValue, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrUncleanValue, Message: ""}}}
	case "102000":
		return dto.FlipErrorResponse{Code: dto.FlipErrOnlyNumbers, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrOnlyNumbers, Message: ""}}}
	case "102100":
		return dto.FlipErrorResponse{Code: dto.FlipErrAmountBelowMinimum, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrAmountBelowMinimum, Message: ""}}}
	case "102200":
		return dto.FlipErrorResponse{Code: dto.FlipErrAmountAboveMaximum, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrAmountAboveMaximum, Message: ""}}}
	case "102400":
		return dto.FlipErrorResponse{Code: dto.FlipErrMaxCharExceeded, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrMaxCharExceeded, Message: ""}}}
	case "102500":
		return dto.FlipErrorResponse{Code: dto.FlipErrInvalidAccountNumber, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrInvalidAccountNumber, Message: ""}}}
	case "102600":
		return dto.FlipErrorResponse{Code: dto.FlipErrFraudSuspectedAccount, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrFraudSuspectedAccount, Message: ""}}}
	case "102700":
		return dto.FlipErrorResponse{Code: dto.FlipErrClosedAccount, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrClosedAccount, Message: ""}}}
	case "103200":
		return dto.FlipErrorResponse{Code: dto.FlipErrInvalidPagination, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrInvalidPagination, Message: ""}}}
	case "103300":
		return dto.FlipErrorResponse{Code: dto.FlipErrInvalidBankCode, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrInvalidBankCode, Message: ""}}}
	case "103400":
		return dto.FlipErrorResponse{Code: dto.FlipErrInvalidCountryCode, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrInvalidCountryCode, Message: ""}}}
	case "103500":
		return dto.FlipErrorResponse{Code: dto.FlipErrInsufficientBalance, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrInsufficientBalance, Message: ""}}}
	case "103800":
		return dto.FlipErrorResponse{Code: dto.FlipErrInvalidCountryOrCityCode, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrInvalidCountryOrCityCode, Message: ""}}}
	case "103900":
		return dto.FlipErrorResponse{Code: dto.FlipErrInvalidDateFormat, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrInvalidDateFormat, Message: ""}}}
	case "104000":
		return dto.FlipErrorResponse{Code: dto.FlipErrInvalidDate, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrInvalidDate, Message: ""}}}
	case "104100":
		return dto.FlipErrorResponse{Code: dto.FlipErrInvalidAttribute, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrInvalidAttribute, Message: ""}}}
	case "104200":
		return dto.FlipErrorResponse{Code: dto.FlipErrMissingIdempotencyKey, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrMissingIdempotencyKey, Message: ""}}}
	case "104300":
		return dto.FlipErrorResponse{Code: dto.FlipErrMissingBillTitle, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrMissingBillTitle, Message: ""}}}
	case "107000":
		return dto.FlipErrorResponse{Code: dto.FlipErrMaxBeneficiaryEmail, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrMaxBeneficiaryEmail, Message: ""}}}
	case "107100":
		return dto.FlipErrorResponse{Code: dto.FlipErrInvalidBeneficiaryEmail, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrInvalidBeneficiaryEmail, Message: ""}}}
	case "107200":
		return dto.FlipErrorResponse{Code: dto.FlipErrDisbursementIDNotFound, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrDisbursementIDNotFound, Message: ""}}}
	case "107300":
		return dto.FlipErrorResponse{Code: dto.FlipErrDisbursementIdemKeyNotFound, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrDisbursementIdemKeyNotFound, Message: ""}}}
	case "107400":
		return dto.FlipErrorResponse{Code: dto.FlipErrDailyLimitExceeded, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrDailyLimitExceeded, Message: ""}}}
	case "108000":
		return dto.FlipErrorResponse{Code: dto.FlipErrMaxActiveTransactions, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrMaxActiveTransactions, Message: ""}}}
	case "108800":
		return dto.FlipErrorResponse{Code: dto.FlipErrBankDisturbance, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrBankDisturbance, Message: ""}}}
	case "108900":
		return dto.FlipErrorResponse{Code: dto.FlipErrFlipAccountTarget, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrFlipAccountTarget, Message: ""}}}
	case "109000":
		return dto.FlipErrorResponse{Code: dto.FlipErrAgentKYCNotApproved, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrAgentKYCNotApproved, Message: ""}}}
	case "109100":
		return dto.FlipErrorResponse{Code: dto.FlipErrAgentNotActive, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrAgentNotActive, Message: ""}}}
	case "109200":
		return dto.FlipErrorResponse{Code: dto.FlipErrAgentUpdateNotAllowed, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrAgentUpdateNotAllowed, Message: ""}}}
	case "109300":
		return dto.FlipErrorResponse{Code: dto.FlipErrBankCutoffTime, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrBankCutoffTime, Message: ""}}}
	case "109400":
		return dto.FlipErrorResponse{Code: dto.FlipErrStaleRequest, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrStaleRequest, Message: ""}}}
	case "109500":
		return dto.FlipErrorResponse{Code: dto.FlipErrInvalidTimestampFormat, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrInvalidTimestampFormat, Message: ""}}}
	case "200100":
		return dto.FlipErrorResponse{Code: dto.FlipErrCharCountConstraint, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrCharCountConstraint, Message: ""}}}
	case "200200":
		return dto.FlipErrorResponse{Code: dto.FlipErrDuplicateAttribute, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrDuplicateAttribute, Message: ""}}}
	case "200300":
		return dto.FlipErrorResponse{Code: dto.FlipErrSanitizedValueOnly, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrSanitizedValueOnly, Message: ""}}}
	case "200400":
		return dto.FlipErrorResponse{Code: dto.FlipErrNonAlphanumeric, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrNonAlphanumeric, Message: ""}}}
	case "200500":
		return dto.FlipErrorResponse{Code: dto.FlipErrBelowBankMinimumAmount, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrBelowBankMinimumAmount, Message: ""}}}
	}
	return
}

func Inquiry() (response dto.FlipErrorResponse) {
	return dto.FlipErrorResponse{Code: dto.FlipErrDisbursementIdemKeyNotFound, Errors: []dto.FlipErrorAttribute{{Attribute: "", Code: dto.FlipErrDisbursementIdemKeyNotFound, Message: ""}}}
}
