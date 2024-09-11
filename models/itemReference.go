package models

import (
	"errors"
	"kontentaimanagementsdkgo/models/difference"
)

type ItemReference struct {
	Id         string `json:"id,omitempty"`
	Codename   string `json:"codename,omitempty"`
	ExternalId string `json:"external_id,omitempty"`
}

func (lang ItemReference) GetDifferences(newModel interface{}) ([]difference.Difference, error) {
	var newItemRef, ok = newModel.(ItemReference)

	if !ok {
		return nil, errors.New("the provided model is not of the correct type. Expected ItemReference")
	}

	var differences []difference.Difference

	if lang.Id != newItemRef.Id {
		differences = append(differences, difference.Difference{
			Operation: "replace",
			PropertyName: "id",
			Value:        newItemRef.Id,
		})
	}

	if lang.Codename != newItemRef.Codename {
		differences = append(differences, difference.Difference{
			Operation: "replace",
			PropertyName: "codename",
			Value:        newItemRef.Codename,
		})
	}

	if lang.ExternalId != newItemRef.ExternalId {
		differences = append(differences, difference.Difference{
			Operation: "replace",
			PropertyName: "external_id",
			Value:        newItemRef.ExternalId,
		})
	}

	return differences, nil
}