package models

import "encoding/json"

type ContentType struct {
	Id            string            `json:"id"`
	Codename      string            `json:"codename"`
	ExternalId    string            `json:"external_id"`
	LastModified  string            `json:"last_modified"`
	Name          string            `json:"name"`
	ContentGroups []ContentGroup    `json:"content_groups"`
	Elements      []BaseElement	 `json:"-"`
	ElementsRaw   []json.RawMessage `json:"elements"`
}