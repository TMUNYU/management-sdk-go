package models

import (
	"errors"
	"kontentaimanagementsdkgo/models/difference"
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

func (lang Language) GetDifferences(newModel interface{}) ([]difference.Difference, error) {
	var newLang, ok = newModel.(Language)

	if !ok {
		return nil, errors.New("the provided model is not of the correct type. Expected Language")
	}

	var differences []difference.Difference

	if lang.Name != newLang.Name {
		differences = append(differences, difference.Difference{
			Operation:    "replace",
			PropertyName: "name",
			Value:        newLang.Name,
		})
	}

	if lang.Codename != newLang.Codename {
		differences = append(differences, difference.Difference{
			Operation:    "replace",
			PropertyName: "codename",
			Value:        newLang.Codename,
		})
	}

	if lang.ExternalId != newLang.ExternalId {
		differences = append(differences, difference.Difference{
			Operation:    "replace",
			PropertyName: "external_id",
			Value:        newLang.ExternalId,
		})
	}

	if lang.IsActive != newLang.IsActive {
		differences = append(differences, difference.Difference{
			Operation:    "replace",
			PropertyName: "is_active",
			Value:        newLang.IsActive,
		})
	}

	if lang.IsDefault != newLang.IsDefault {
		differences = append(differences, difference.Difference{
			Operation:    "replace",
			PropertyName: "is_default",
			Value:        newLang.IsDefault,
		})
	}

	var fallBackDiff, err = difference.IModel(lang.FallbackLanguage).GetDifferences(newLang.FallbackLanguage)
	if err != nil {
		return nil, err
	}

	differences = append(differences, fallBackDiff...)

	return differences, nil
}
