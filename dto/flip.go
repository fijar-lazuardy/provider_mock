package dto

type FlipErrorAttribute struct {
	Attribute string `json:"attribute"`
	Code      string `json:"code"`
	Message   string `json:"message"`
}

type FlipErrorResponse struct {
	Code   string               `json:"code"`
	Errors []FlipErrorAttribute `json:"errors"`
}

const (
	FlipErrUndefined                   string = "999"
	FlipErrRequiredAttribute           string = "1001"
	FlipErrUncleanValue                string = "1002"
	FlipErrOnlyNumbers                 string = "1020"
	FlipErrAmountBelowMinimum          string = "1021"
	FlipErrAmountAboveMaximum          string = "1022"
	FlipErrMaxCharExceeded             string = "1024"
	FlipErrInvalidAccountNumber        string = "1025"
	FlipErrFraudSuspectedAccount       string = "1026"
	FlipErrClosedAccount               string = "1027"
	FlipErrInvalidPagination           string = "1032"
	FlipErrInvalidBankCode             string = "1033"
	FlipErrInvalidCountryCode          string = "1034"
	FlipErrInsufficientBalance         string = "1035"
	FlipErrInvalidCountryOrCityCode    string = "1038"
	FlipErrInvalidDateFormat           string = "1039"
	FlipErrInvalidDate                 string = "1040"
	FlipErrInvalidAttribute            string = "1041"
	FlipErrMissingIdempotencyKey       string = "1042"
	FlipErrMissingBillTitle            string = "1043"
	FlipErrMaxBeneficiaryEmail         string = "1070"
	FlipErrInvalidBeneficiaryEmail     string = "1071"
	FlipErrDisbursementIDNotFound      string = "1072"
	FlipErrDisbursementIdemKeyNotFound string = "1073"
	FlipErrDailyLimitExceeded          string = "1074"
	FlipErrMaxActiveTransactions       string = "1080"
	FlipErrBankDisturbance             string = "1088"
	FlipErrFlipAccountTarget           string = "1089"
	FlipErrAgentKYCNotApproved         string = "1090"
	FlipErrAgentNotActive              string = "1091"
	FlipErrAgentUpdateNotAllowed       string = "1092"
	FlipErrBankCutoffTime              string = "1093"
	FlipErrStaleRequest                string = "1094"
	FlipErrInvalidTimestampFormat      string = "1095"
	FlipErrCharCountConstraint         string = "2001"
	FlipErrDuplicateAttribute          string = "2002"
	FlipErrSanitizedValueOnly          string = "2003"
	FlipErrNonAlphanumeric             string = "2004"
	FlipErrBelowBankMinimumAmount      string = "2005"
)

