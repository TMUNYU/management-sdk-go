package models_tests

import (
	"encoding/json"
	helper_models "kontentaimanagementsdkgo/models/helper"
	"testing"
)

func Test_ThatDependsOnUnmarshallsCorrectly(t *testing.T) {
	var objUnderTest helper_models.DependsOn

	var jsonData = `{
		"element": {
			"external_id": "external_id",
			"id": "id",
			"codename": "codename"
		},
		"snippet": {
			"external_id": "external_id",
			"id": "id",
			"codename": "codename"
		}
	}`

	err := json.Unmarshal([]byte(jsonData), &objUnderTest)

	if err != nil {
		t.Errorf("Unmarshalling failed")
	}

	if objUnderTest.Element.ExternalId != "external_id" {
		t.Errorf("Element.ExternalId is not equal")
	}

	if objUnderTest.Element.Id != "id" {
		t.Errorf("Element.Id is not equal")
	}

	if objUnderTest.Element.Codename != "codename" {
		t.Errorf("Element.Codename is not equal")
	}

	if objUnderTest.Snippet.ExternalId != "external_id" {
		t.Errorf("Snippet.ExternalId is not equal")
	}

	if objUnderTest.Snippet.Id != "id" {
		t.Errorf("Snippet.Id is not equal")
	}

	if objUnderTest.Snippet.Codename != "codename" {
		t.Errorf("Snippet.Codename is not equal")
	}
}
