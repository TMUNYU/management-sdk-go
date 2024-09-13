package models

type KontentResponseWrapper struct {
	Types      []ContentType     `json:"types,omitempty"`
	Languages  []Language `json:"languages,omitempty"`
	Pagination Pagination `json:"pagination"`
}