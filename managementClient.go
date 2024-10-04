package main

import (
	"kontentaimanagementsdkgo/models"
	"kontentaimanagementsdkgo/http_clients"
	"kontentaimanagementsdkgo/repositories"
)

type KontentAiManagementClient struct {
	ApiKey string
}

func NewKontentAiManagementClient(apiKey string) *KontentAiManagementClient {
	return &KontentAiManagementClient{
		ApiKey: apiKey,
	}
}

func (client *KontentAiManagementClient) ListContentTypes() *[]models.ContentType {
	var httpClient = http_clients.NewKontentAiHttpManagementClient(models.KontentManagementConfiguration{
		ApiKey: client.ApiKey,
	})

	var contentTypeRepository = repositories.NewContentTypeRepository(httpClient)

	var contenTypes, err = contentTypeRepository.GetContentTypes()
	if err != nil {
		return nil
	}

	return contenTypes
}
