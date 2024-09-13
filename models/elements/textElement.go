package elements

import (
	 "kontentaimanagementsdkgo/models"
	  "kontentaimanagementsdkgo/models/validation"
)

type TextElement struct {
	models.BaseElement
	Guidelines       string          `json:"guidelines"`
	IsRequired       bool            `json:"is_required"`
	IsNonLocalizable bool            `json:"is_non_localizable"`
	ValidationRegex  validation.ValidationRegexValidation `json:"validation_regex"`
	MaximumTextLength validation.MaximumTextLengthValidation `json:"maximum_text_length"`
	Default          models.Default `json:"default"`
	ContentGroup     models.ItemReference   `json:"content_group"`
}


/*
{
"name": "Title",
"guidelines": "Keep it short.",
"is_required": "false",
"is_non_localizable": "true",
"type": "text",
"id": "7dc115d0-e9f8-4947-9cfb-9a26485975ae",
"codename": "title",
"external_id": "my-text-element",
"maximum_text_length": {
"value": 60,
"applies_to": "characters"
},
"validation_regex": {
"regex": "^(https?:\\/\\/(?:www\\.|(?!www)))?[a-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,5}(:[0-9]{1,5})?(\\/.*)?$",
"flags": "i",
"validation_message": "Type a valid web URL such as https://example.org.",
"is_active": "true"
},
"default": {
"global": {
"value": "The greatest title"
}
},
"content_group": {
"id": "ad1b8cce-94d6-4682-a9ff-393d494e3a02"
}
}*/