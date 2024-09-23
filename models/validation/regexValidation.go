package validation

type RegexValidation struct {
	Regex             string `json:"regex"`
	Flags             string `json:"flags"`
	ValidationMessage string `json:"validation_message"`
	IsActive          bool   `json:"is_active"`
}
