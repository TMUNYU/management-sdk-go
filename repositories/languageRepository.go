package repositories

import (
	"encoding/json"
	"kontentaimanagementsdkgo/clients"
	"kontentaimanagementsdkgo/models"
)

type LanguageRepository struct {
	client clients.IManagementClient
}

func NewLanguageRepository(client clients.IManagementClient) LanguageRepository {
	var repository = LanguageRepository{
		client:  client,	
	}
	return repository
}

func (repository LanguageRepository) GetLanguages() (*[]models.Language, error) {
	htpClient := repository.client
	resp, error := htpClient.Get("languages")

	if error != nil {
		return nil, error
	}

	var response models.KontentResponseWrapper
	
	error = json.Unmarshal(resp, &response)
	if error != nil {
		return nil, error		
	}

	return &response.Languages, nil	
}

func (repository LanguageRepository) GetLanguage(id string) (*models.Language, error) {
	htpClient := repository.client
	resp, error := htpClient.Get("languages/" + id)

	if error != nil {
		return nil, error
	}

	var language models.Language
	
	error = json.Unmarshal(resp, &language)
	if error != nil {
		return nil, error		
	}

	return &language, nil	
}

func (repository LanguageRepository) CreateLanguage(language models.Language) (*models.Language, error) {
	htpClient := repository.client
	payload, error := json.Marshal(language)

	if error != nil {
		return nil, error
	}

	resp, error := htpClient.Post("languages", string(payload))

	if error != nil {
		return nil, error
	}

	var createdLanguage models.Language
	
	error = json.Unmarshal(resp, &createdLanguage)
	if error != nil {
		return nil, error		
	}

	return &createdLanguage, nil	
}

func (repository LanguageRepository) UpdateLanguage(language models.Language) (*models.Language, error) {
	
}

