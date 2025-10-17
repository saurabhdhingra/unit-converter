package model

type RequestPayload struct {
	Value		float64		`json:"value"`
	UnitFrom	string		`json:"unitFrom"`
	UnitTo		string		`json:"unitTo"`
}

type APIResponse struct {
	OriginalValue	float64	`json:"originalValue"`
	FromUnit		string	`json:"fromUnit"`
	TargetValue		float64	`json:"result"`
	ToUnit			string	`json:"toUnit"`
	ConversionType	string	`json:"conversionType"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type ConversionType string

const (
	Temp	ConversionType = "Temperature"
	Len		ConversionType = "Length"
	Wgt		ConversionType = "Weight"
	None	ConversionType = ""
)