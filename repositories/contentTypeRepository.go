package repositories

import (
	"encoding/json"
	"kontentaimanagementsdkgo/http_clients"
	"kontentaimanagementsdkgo/models"
	"kontentaimanagementsdkgo/models/wrappers"
	"log"
)

type ContentTypeRepository struct {
	client http_clients.IKontentAiManagementClient
}

func NewContentTypeRepository(client http_clients.IKontentAiManagementClient) ContentTypeRepository {
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
	httpClient := repository.client
	resp, err := httpClient.Get("types/" + id)

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
	httpClient := repository.client

	payload, err := json.Marshal(contentType)
	if err != nil {
		return nil, err
	}

	log.Println(string(payload))

	result, err := httpClient.Post("types", string(payload))
	if err != nil {
		return nil, err
	}

	var newContentType models.ContentType
	err = json.Unmarshal(result, &newContentType)
	if err != nil {
		return nil, err
	}

	return &newContentType, nil
}
