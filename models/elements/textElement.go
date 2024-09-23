package elements

import (
	helper_types "kontentaimanagementsdkgo/models/elements/types"
	helper_models "kontentaimanagementsdkgo/models/helper"
	"kontentaimanagementsdkgo/models/validation"
)

type TextElement struct {
	BaseElement
	Guidelines        string                             `json:"guidelines"`
	IsRequired        bool                               `json:"is_required"`
	IsNonLocalizable  bool                               `json:"is_non_localizable"`
	Type              string                             `json:"type"`
	ValidationRegex   validation.RegexValidation         `json:"validation_regex"`
	MaximumTextLength validation.MaxTextLengthValidation `json:"maximum_text_length"`
	Default           helper_types.DefaultValue                `json:"default"`
	ContentGroup      helper_models.ItemReference               `json:"content_group"`
}
