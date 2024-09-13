package elements

import (
	"kontentaimanagementsdkgo/models"
)

type SlugElement struct {
	models.BaseElement
	DependsOn        models.DependsOn             `json:"depends_on"`
	Guidelines       string          `json:"guidelines"`
	IsRequired       bool            `json:"is_required"`
	IsNonLocalizable bool            `json:"is_non_localizable"`
	ValidationRegex  models.ValidationRegex `json:"validation_regex"`
	ContentGroup     models.ItemReference   `json:"content_group"`
}