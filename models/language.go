package models

import (
	helper_models "kontentaimanagementsdkgo/models/helper"
)

type Language struct {
	Id               string                      `json:"id"`
	Name             string                      `json:"name"`
	Codename         string                      `json:"codename"`
	ExternalId       string                      `json:"external_id"`
	IsActive         bool                        `json:"is_active"`
	IsDefault        bool                        `json:"is_default"`
	FallbackLanguage helper_models.ItemReference `json:"fallback_language,omitempty"`
}
