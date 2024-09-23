package models

import (
	"encoding/json"
	"kontentaimanagementsdkgo/models/elements"
	helper_types "kontentaimanagementsdkgo/models/elements/types"
)

type ContentType struct {
	Id            string            `json:"id"`
	Codename      string            `json:"codename"`
	ExternalId    string            `json:"external_id"`
	LastModified  string            `json:"last_modified"`
	Name          string            `json:"name"`
	ContentGroups []helper_types.ContentGroup    `json:"content_groups"`
	Elements      []elements.BaseElement     `json:"-"`
}

func (d *ContentType) UnmarshalJSON(data []byte) error {
	type ContentTypeAlias ContentType

	var elementsRaw []json.RawMessage

	aux := &struct {
		Elements    *[]elements.BaseElement     `json:"-"`
		ElementsRaw *[]json.RawMessage `json:"elements"`
		*ContentTypeAlias
	}{
		ContentTypeAlias: (*ContentTypeAlias)(d),
		Elements:         &d.Elements,
		ElementsRaw:      &elementsRaw,
	}
	err := json.Unmarshal(data, &aux)

	// returnValue := d.unmarshallElements(elementsRaw)
	// if err != nil {
	// 	return returnValue
	// }

	return err
}
