package elements

import (
	helper_types "kontentaimanagementsdkgo/models/elements/types"
	helper_models "kontentaimanagementsdkgo/models/helper"
	"kontentaimanagementsdkgo/models/validation"
)

type TextElement struct {
	BaseElement
	Guidelines        string                              `json:"guidelines,omitempty"`
	IsRequired        bool                                `json:"is_required,omitempty"`
	IsNonLocalizable  bool                                `json:"is_non_localizable,omitempty"`
	ValidationRegex   *validation.RegexValidation         `json:"validation_regex,omitempty"`
	MaximumTextLength *validation.MaxTextLengthValidation `json:"maximum_text_length,omitempty"`
	Default           *helper_types.DefaultValue          `json:"default,omitempty"`
	ContentGroup      *helper_models.ItemReference        `json:"content_group,omitempty"`
}
