package helper_models

type ItemReference struct {
	Id         string `json:"id,omitempty"`
	Codename   string `json:"codename,omitempty"`
	ExternalId string `json:"external_id,omitempty"`
}
