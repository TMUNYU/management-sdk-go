package elements

import (
	helper_models "kontentaimanagementsdkgo/models/helper"
	"kontentaimanagementsdkgo/models/validation"
)

type SlugElement struct {
	BaseElement
	DependsOn        helper_models.DependsOn           `json:"depends_on"`
	Guidelines       string                     `json:"guidelines"`
	IsRequired       bool                       `json:"is_required"`
	IsNonLocalizable bool                       `json:"is_non_localizable"`
	ValidationRegex  validation.RegexValidation `json:"validation_regex"`
	ContentGroup     helper_models.ItemReference       `json:"content_group"`
}
