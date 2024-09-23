package repositories

import (
	"encoding/json"
	"kontentaimanagementsdkgo/clients"
	"kontentaimanagementsdkgo/models"
	"kontentaimanagementsdkgo/models/wrappers"
)

type ContentTypeRepository struct {
	client clients.IManagementClient
}

func NewContentTypeRepository(client clients.IManagementClient) ContentTypeRepository {
	var repository = ContentTypeRepository{
		client: client,
	}
	return repository
}

func (repository ContentTypeRepository) GetContentTypes() (*[]models.ContentType, error) {
	htpClient := repository.client

	resp, error := htpClient.Get("types")
	if error != nil {
		return nil, error
	}

	var response wrappers.KontentResponseWrapper

	error = json.Unmarshal(resp, &response)
	if error != nil {
		return nil, error
	}

	var types []models.ContentType

	return &types, nil
}

func (repository ContentTypeRepository) GetContentType(id string) (*models.ContentType, error) {
	htpClient := repository.client
	resp, err := htpClient.Get("types/" + id)

	if err != nil {
		return nil, err
	}

	var contentType models.ContentType

	err = json.Unmarshal(resp, &contentType)
	if err != nil {
		return nil, err
	}

	return &contentType, nil
}

func (repository ContentTypeRepository) CreateContentType(contentType models.ContentType) (*models.ContentType, error) {
	panic("not implemented")
}