package models

import (
	"kontentaimanagementsdkgo/guards"
	"strings"
)

type Url struct {
	Scheme      string
	Host        string
	Segments    []string
	QueryParams map[string]string
}

func NewUrlWithSegments(url string, segments []string) (Url, error) {
	guards.UrlSyntacticallyValid(url)
	
	var urlObj, err = parseUrl(url)
	if err != nil {
		return Url{}, err
	}

	urlObj.Segments = segments

	return urlObj, nil
}

func (url Url) ToUrlString() string {
	var urlStr = url.Scheme + "://" + url.Host

	for _, segment := range url.Segments {
		urlStr += "/" + segment
	}

	if len(url.QueryParams) > 0 {
		urlStr += "?"

		for key, value := range url.QueryParams {
			urlStr += key + "=" + value + "&"
		}

		urlStr = urlStr[:len(urlStr)-1]
	}

	return urlStr
}

func parseUrl(url string) (Url, error) {
	var querySplit = strings.Split(url, "?")
	var baseSplit = strings.Split(querySplit[0], "://")

	var scheme = baseSplit[0]
	var host = strings.Split(baseSplit[1], "/")[0]
	var segments = strings.Split(baseSplit[1], "/")[1:]
	var queryParams = make(map[string]string)

	var querySegments []string
	if len(querySplit) > 1 {
		querySegments = strings.Split(querySplit[1], "&")
		for _, querySegment := range querySegments {
			var queryParam = strings.Split(querySegment, "=")
			queryParams[queryParam[0]] = queryParam[1]
		}		
	}

	var urlObj = Url{
		Scheme:      scheme,
		Host:        host,
		Segments:    segments,
		QueryParams: queryParams,
	}

	return urlObj, nil
}