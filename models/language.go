package models

type Language struct {
	Id                string        `json:"id"`
	Name              string        `json:"name"`
	Codename          string        `json:"codename"`
	ExternalId        string        `json:"external_id"`
	IsActive          bool          `json:"is_active"`
	IsDefault         bool          `json:"is_default"`
	Fallback_language ItemReference `json:"fallback_language,omitempty"`
}