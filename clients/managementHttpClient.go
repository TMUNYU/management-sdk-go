package clients

import (
	"errors"
	"io"
	"kontentaimanagementsdkgo/models"
	"net/http"
	"strings"
)

type IManagementClient interface {
	Get(uri string) ([]byte, error)
	Post(uri string, payload string) ([]byte, error)
	Put(uri string, payload string) ([]byte, error)
	Patch(uri string, payload string) ([]byte, error)
	Delete(uri string) (bool, error)	
}

type KontentAiHttpManagementClient struct {
	httpClient    http.Client
	configuration models.KontentManagementConfiguration
}

func NewKontentManagementClient(configuration models.KontentManagementConfiguration) IManagementClient {
	var client = IManagementClient(KontentAiHttpManagementClient{
		httpClient: http.Client{},
		configuration: configuration,
	})

	return client
}

func (client KontentAiHttpManagementClient) Get(uri string) ([]byte, error) {
	return performRequest(client, http.MethodGet, uri, "", evaluateSuccessStatusCodeSuccessDefault)
}

func (client KontentAiHttpManagementClient) Post(uri string, payload string) ([]byte, error) {
	return performRequest(client, http.MethodPost, uri, payload, evaluateSuccessStatusCodeSuccessDefault)
}

func (client KontentAiHttpManagementClient) Put(uri string, payload string) ([]byte, error) {
	return performRequest(client, http.MethodPut, uri, payload, evaluateSuccessStatusCodeSuccessDefault)
}

func (client KontentAiHttpManagementClient) Patch(uri string, payload string) ([]byte, error) {
	return performRequest(client, http.MethodPatch, uri, payload, evaluateSuccessStatusCodeSuccessDefault)
}

func (client KontentAiHttpManagementClient) Delete(uri string) (bool, error) {
	_, err := performRequest(client, http.MethodDelete, uri, "", evaluateSuccessStatusCodeSuccessDefault)

	if err != nil {
		return false, err
	}

	return true, nil
}

func performRequest(client KontentAiHttpManagementClient, method string, uri string, payload string, isSuccess func(s int) bool) ([]byte, error) {
	req, err := prepareRequest(client, method, sanitizeUri(uri))

	if err != nil {
		return nil, err
	}

	if payload != "" {
		req.Body = io.NopCloser(strings.NewReader(payload))
	}

	resp, err := client.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	if !isSuccess(resp.StatusCode) {
		var bytes, err = readBody(resp)

		if err != nil {
			return nil, errors.New("Request failed with status code: " + resp.Status + ". Response: " + "Failed to read response body")
		}

		return nil, errors.New("Request failed with status code: " + resp.Status + ". Response: " + string(bytes))
	}

	bytes, err := readBody(resp)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func prepareRequest(client KontentAiHttpManagementClient, method string, uri string) (*http.Request, error) {
	var url, err = models.NewUrlWithSegments(client.configuration.Url, []string{"v2", "projects", client.configuration.EnvironmentId, uri})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url.ToUrlString(), nil)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"Authorization": []string{"Bearer " + client.configuration.ApiKey},
		"Content-Type":  []string{"application/json"},
	}

	return req, nil
}

func readBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func sanitizeUri(uri string) string {
	if uri[0] == '/' {
		return uri[1:]
	}

	return uri
}

func evaluateSuccessStatusCodeSuccessDefault(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}