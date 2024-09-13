```csharp
package main

import (
	"encoding/json"
	"kontentaimanagementsdkgo/clients"
	"kontentaimanagementsdkgo/models"
	"kontentaimanagementsdkgo/repositories"
	"log"
)

func main() {
	var config = models.KontentManagementConfiguration{
		Url:           "https://manage.kontent.ai",
		ApiVersion:    "v2",
		ApiKey:        "from content",
		EnvironmentId: "8dad9221-54da-0027-e05e-e4fc51de6369",
	}

	var client = clients.NewKontentAiHttpManagementClient(config)
	var contentTypeRepo = repositories.NewContentTypeRepository(client)

	contentTypes, err := contentTypeRepo.GetContentTypes()
	if err != nil {
		log.Fatal(err)
	}

	for _, contentType := range *contentTypes {
		var json, _ = json.MarshalIndent(contentType, "", "	")
		log.Println(string(json))
	}
}

```
