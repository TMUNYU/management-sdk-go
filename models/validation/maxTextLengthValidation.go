package validation

type MaxTextLengthValidation struct {
	Value     int    `json:"value"`
	AppliesTo string `json:"applies_to"` // enum: ["characters", "words"]
}
