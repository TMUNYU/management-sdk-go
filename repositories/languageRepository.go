package repositories

import (
	"encoding/json"
	"errors"
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
	resp, error := htpClient.Get("languages" + id)

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

func (repository LanguageRepository) UpdateLanguage(old models.Language, new models.Language) (*models.Language, error) {
	var differences, err = old.GetDifferences(new)
	if err != nil {
		return nil, err
	}

	if len(differences) == 0 {
		return &old, nil
	}

	if !old.IsActive && old.IsActive == new.IsActive && differences[0].PropertyName != "is_active" {
		return nil, errors.New("cannot update an inactive language. Until activated, only action allowed is to activate the language")
	}

	payload, err := json.Marshal(differences)
	if err != nil {
		return nil, err
	}

	htpClient := repository.client

	resp, err := htpClient.Patch("languages/" + old.Id, string(payload))
	if err != nil {
		return nil, err
	}

	var updatedLanguage models.Language
	err = json.Unmarshal(resp, &updatedLanguage)
	if err != nil {
		return nil, err
	}

	return &updatedLanguage, nil
}

