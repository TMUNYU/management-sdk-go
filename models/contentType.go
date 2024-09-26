package models

import (
	"encoding/json"
	"fmt"
	"kontentaimanagementsdkgo/models/elements"
	helper_types "kontentaimanagementsdkgo/models/elements/types"
	"log"
)

type ContentType struct {
	Id            string                      `json:"id,omitempty"`
	Codename      string                      `json:"codename,omitempty"`
	ExternalId    string                      `json:"external_id,omitempty"`
	LastModified  string                      `json:"last_modified,omitempty"`
	Name          string                      `json:"name,omitempty"`
	ContentGroups []helper_types.ContentGroup `json:"content_groups,omitempty"`
	Elements      []any                       `json:"elements,omitempty"`
}

func (d *ContentType) UnmarshalJSON(data []byte) error {
	type ContentTypeAlias ContentType

	aux := &struct {
		ElementsRaw *[]json.RawMessage `json:"elements"`
		Elements    []any              `json:"-"`
		*ContentTypeAlias
	}{
		ContentTypeAlias: (*ContentTypeAlias)(d),
	}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	if aux.ElementsRaw == nil {
		return nil
	}

	for _, elementsRaw := range *aux.ElementsRaw {
		var err error
		var element any
		var elementBase elements.BaseElement
		err = json.Unmarshal(elementsRaw, &elementBase)

		log.Println(string(elementsRaw))

		if err == nil {
			switch elementBase.Type {
			case "text":
				element, err = marshallElement(elementsRaw, &elements.TextElement{})
			case "number":
				element, err = marshallElement(elementsRaw, &elements.NumberElement{})
			case "url_slug":
				element, err = marshallElement(elementsRaw, &elements.SlugElement{})
			default:
				err = fmt.Errorf("element type %s is not supported", elementBase.Type)
			}
		}

		if err != nil {
			return err
		}

		d.Elements = append(d.Elements, element)
	}

	return err
}

func marshallElement[TType interface{}](rawMessage json.RawMessage, element *TType) (*TType, error) {
	err := json.Unmarshal(rawMessage, element)
	if err != nil {
		return nil, err
	}

	return element, nil
}
