package goutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func SendHttpRequest(method, url string, data interface{}, headers *map[string]string, result interface{}) (*[]byte, error) {
	var requestBody io.Reader

	switch method {
	case http.MethodGet:
		if data != nil {
			queryString, err := generateQueryString(data)
			if err != nil {
				return nil, err
			}

			if queryString != nil && len(*queryString) > 0 {
				url = fmt.Sprintf("%s?%s", url, *queryString)
			}
		}
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodPatch:
		if data != nil {
			reqBody, err := generateRequestBody(data, headers)
			if err != nil {
				return nil, err
			}

			requestBody = *reqBody
		}
	}

	cl := &http.Client{}

	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		for k, v := range *headers {
			req.Header.Add(k, v)
		}
	}

	resp, err := cl.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	byteBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if !isHttpResponseOk(resp) {
		err = HttpResponseError{
			Code:         resp.StatusCode,
			Message:      resp.Status,
			ResponseBody: &byteBody,
		}

		return nil, err
	}

	if result != nil {
		if err = json.Unmarshal(byteBody, &result); err != nil {
			return nil, err
		}
	}

	return &byteBody, nil
}

func SendHttpGet(url string, params interface{}, headers *map[string]string, result interface{}) (*[]byte, error) {
	return SendHttpRequest(http.MethodGet, url, params, headers, result)
}

func SendHttpPost(url string, data interface{}, headers *map[string]string, result interface{}) (*[]byte, error) {
	return SendHttpRequest(http.MethodPost, url, data, headers, result)
}

func SendHttpPut(url string, data interface{}, headers *map[string]string, result interface{}) (*[]byte, error) {
	return SendHttpRequest(http.MethodPut, url, data, headers, result)
}

func SendHttpPatch(url string, data interface{}, headers *map[string]string, result interface{}) (*[]byte, error) {
	return SendHttpRequest(http.MethodPatch, url, data, headers, result)
}

func SendHttpDelete(url string, headers *map[string]string, result interface{}) (*[]byte, error) {
	return SendHttpRequest(http.MethodDelete, url, nil, headers, result)
}

func generateRequestBody(data interface{}, headers *map[string]string) (*io.Reader, error) {
	var contentType string

	if headers != nil {
		if cType, exist := (*headers)["Content-Type"]; exist {
			contentType = cType
		}
	}

	var result io.Reader

	switch contentType {
	default:
		bytePostData, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		result = bytes.NewBuffer(bytePostData)
	case "application/x-www-form-urlencoded":
		encodedData, err := generateQueryString(data)
		if err != nil {
			return nil, err
		}

		result = strings.NewReader(*encodedData)
	}

	return &result, nil
}

func generateQueryString(data interface{}) (*string, error) {
	by, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var jsonData map[string]interface{}

	if err = json.Unmarshal(by, &jsonData); err != nil {
		return nil, err
	}

	values := url.Values{}

	for k, v := range jsonData {
		values.Add(k, fmt.Sprintf("%+v", v))
	}

	encodedValues := values.Encode()

	return &encodedValues, nil
}

func isHttpResponseOk(resp *http.Response) bool {
	return resp.StatusCode >= 200 && resp.StatusCode < 300
}
