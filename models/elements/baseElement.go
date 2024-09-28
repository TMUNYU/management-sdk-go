package elements

type BaseElement struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Codename string `json:"codename,omitempty"`
	Type     string `json:"type,omitempty"`
}
