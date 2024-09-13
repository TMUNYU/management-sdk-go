package models

type BaseElement struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Codename string `json:"codename"`
	Type     string `json:"type"`
}