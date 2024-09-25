package elements

import helper_types "kontentaimanagementsdkgo/models/elements/types"

type NumberElement struct {
	BaseElement
	Guidelines       string                    `json:"guidelines"`
	IsRequired       bool                      `json:"is_required,string"`        // Use bool and convert from string
	IsNonLocalizable bool                      `json:"is_non_localizable,string"` // Use bool and convert from string
	ExternalID       string                    `json:"external_id"`
	Default          helper_types.DefaultValue `json:"default"`
	ContentGroup     helper_types.ContentGroup `json:"content_group"`
}
