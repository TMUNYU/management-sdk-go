package models

type KontentResponseWrapper struct {
	Languages  []Language `json:"languages"`
	Pagination Pagination `json:"pagination"`
}