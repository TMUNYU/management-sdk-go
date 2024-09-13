package validation

type ValidationRegexValidation struct {
	Regex             string `json:"regex"`
	Flags             string `json:"flags"`
	ValidationMessage string `json:"validation_message"`
	IsActive          bool   `json:"is_active"`
}