package models

type DependsOn struct {
	Element ItemReference `json:"element,omitempty"`
	Snippet ItemReference `json:"snippet,omitempty"`
}