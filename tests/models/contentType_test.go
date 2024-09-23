package models_tests

import (
	"encoding/json"
	"kontentaimanagementsdkgo/models"
	"testing"
)

func Test_ThatContentTypeBaseUnmarshallsCorrectly(t *testing.T) {
	var jsonData = `{
		"id": "e47a155b-4ca3-53a1-b7a4-47f4844f238e",
		"codename": "manually_added_content_type",
		"external_id": "manually_added_content_type_external_id",
		"last_modified": "2024-09-20T13:42:05.1194857Z",
		"name": "Manually Added Content Type"
	}`

	var objUnderTest models.ContentType

	err := json.Unmarshal([]byte(jsonData), &objUnderTest)

	if err != nil {
		t.Errorf("Unmarshalling failed")
	}

	if objUnderTest.Id != "e47a155b-4ca3-53a1-b7a4-47f4844f238e" {
		t.Errorf("Id is not equal")
	}

	if objUnderTest.Codename != "manually_added_content_type" {
		t.Errorf("Codename is not equal")
	}

	if objUnderTest.ExternalId != "manually_added_content_type_external_id" {
		t.Errorf("ExternalId is not equal")
	}

	if objUnderTest.LastModified != "2024-09-20T13:42:05.1194857Z" {
		t.Errorf("LastModified is not equal")
	}

	if objUnderTest.Name != "Manually Added Content Type" {
		t.Errorf("Name is not equal")
	}
}

func Test_ThatContentTypeUnmarshallsContentGroupsCorrectly(t *testing.T) {
	var jsonData = `{
		"content_groups": [
			{
				"id": "e47a155b-4ca3-53a1-b7a4-47f4844f238e",
				"name": "Manually Added Content Group",
				"codename": "manually_added_content_group"
			}
		]
	}`

	var objUnderTest models.ContentType

	err := json.Unmarshal([]byte(jsonData), &objUnderTest)

	if err != nil {
		t.Errorf("Unmarshalling failed")
	}

	if objUnderTest.ContentGroups[0].Id != "e47a155b-4ca3-53a1-b7a4-47f4844f238e" {
		t.Errorf("ContentGroups[0].Id is not equal")
	}

	if objUnderTest.ContentGroups[0].Name != "Manually Added Content Group" {
		t.Errorf("ContentGroups[0].Name is not equal")
	}

	if objUnderTest.ContentGroups[0].Codename != "manually_added_content_group" {
		t.Errorf("ContentGroups[0].Codename is not equal")
	}
}

func Test_ThatContentTypeUnmarshallsElementsCorrectly(t *testing.T) {
	var slugJsonData = `{
		"id": "e47a155b-4ca3-53a1-b7a4-47f4844f238e",
		"name": "Manually Added Element",
		"codename": "manually_added_element",
		"type": "url_slug",
		"guidelines": "This is a manually added element"
	}`
	
	var numberJsonData = `{
        "name": "Number",
        "guidelines": "Provide a number.",
        "is_required": "false",
        "is_non_localizable": "true",
        "type": "number",
        "id": "b18cbb42-9649-4870-b2d6-805238a34ae9",
        "codename": "number",
        "external_id": "my-number",
        "default": {
            "global": {
                "value": 42
            }
        },
        "content_group": {
            "id": "ad1b8cce-94d6-4682-a9ff-393d494e3a02"
        }
    }`

	var jsonData = `{
		"elements": [
			` + slugJsonData + `,
			` + numberJsonData + `
		]
	}`

	var objUnderTest models.ContentType

	err := json.Unmarshal([]byte(jsonData), &objUnderTest)

	if err != nil {
		t.Errorf("Unmarshalling failed")
	}
}