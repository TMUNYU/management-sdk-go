package elements

type BaseElement struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Codename string `json:"codename"`
	Type     string `json:"type,omitempty"`
}
