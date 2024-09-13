package repositories

import (
	"encoding/json"
	"errors"
	"kontentaimanagementsdkgo/clients"
	"kontentaimanagementsdkgo/models"
	"kontentaimanagementsdkgo/models/elements"
)

type ContentTypeRepository struct {
	client clients.IManagementClient
}

func NewContentTypeRepository(client clients.IManagementClient) ContentTypeRepository {
	var repository = ContentTypeRepository{
		client:  client,	
	}
	return repository
}

func (repository ContentTypeRepository) GetContentTypes() (*[]models.ContentType, error) {
	htpClient := repository.client

	resp, error := htpClient.Get("types")
	if error != nil {
		return nil, error
	}

	var response models.KontentResponseWrapper
	
	error = json.Unmarshal(resp, &response)
	if error != nil {
		return nil, error		
	}
	
	var types []models.ContentType
	for _, contentType := range response.Types {
		var elements []models.BaseElement
		for _, elementRaw := range contentType.ElementsRaw {
			var element, err = unmarshallElement(elementRaw)
			if err != nil {
				return nil, err
			}
			elements = append(elements, *element)
		}

		contentType.Elements = elements

		types = append(types, contentType)
	}

	return &types, nil	
}

func (repository ContentTypeRepository) GetContentType(id string) (*models.ContentType, error) {
	htpClient := repository.client
	resp, error := htpClient.Get("types/" + id)

	if error != nil {
		return nil, error
	}

	var contentType models.ContentType
	
	error = json.Unmarshal(resp, &contentType)
	if error != nil {
		return nil, error		
	}
	
	for _, elementRaw := range contentType.ElementsRaw {
		var element, err = unmarshallElement(elementRaw)
		if err != nil {
			return nil, err
		}
		contentType.Elements = append(contentType.Elements, *element)
	}

	return &contentType, nil	
}

func (repository ContentTypeRepository) CreateContentType(contentType models.ContentType) (*models.ContentType, error) {
		httpClient := repository.client

	payload, error := json.Marshal(contentType)
	if error != nil {
		return nil, error
	}

	resp, error := httpClient.Post("types", string(payload))
	if error != nil {
		return nil, error
	}

	var createContentType models.ContentType
	
	error = json.Unmarshal(resp, &createContentType)
	if error != nil {
		return nil, error		
	}

	return &createContentType, nil	
}

func unmarshallElement(data json.RawMessage) (*models.BaseElement, error) {
	var element models.BaseElement
	var err = json.Unmarshal(data, &element)
	if err != nil {
		return nil, err
	}

	switch element.Type {
		case "url_slug":
			var slug elements.SlugElement
			err = json.Unmarshal(data, &slug)
			return &slug.BaseElement, err
		case "text":
			var text elements.TextElement
			err = json.Unmarshal(data, &text)
			return &text.BaseElement, err
		default:
			return nil, errors.New("Unknown element type " + element.Type)
	}
}