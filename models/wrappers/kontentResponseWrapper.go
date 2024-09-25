package wrappers

import (
	"kontentaimanagementsdkgo/models"
	helper_models "kontentaimanagementsdkgo/models/helper"
)

type KontentResponseWrapper struct {
	Types      []models.ContentType     `json:"types,omitempty"`
	Languages  []models.Language        `json:"languages,omitempty"`
	Pagination helper_models.Pagination `json:"pagination"`
}
