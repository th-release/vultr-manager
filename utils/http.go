package utils

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func PostRequest[T any](client *resty.Client, url string, data interface{}, queryParams map[string]string, token string) (*resty.Response, T, interface{}, error) {
	var result T
	var errorResult interface{}
	req := client.R().
		SetError(&errorResult).
		SetHeader("Authorization", "Bearer "+token).
		SetBody(data).
		SetResult(&result)

	if len(queryParams) > 0 {
		req = req.SetQueryParams(queryParams)
	}

	resp, err := req.Post(url)

	if err != nil {
		return nil, result, errorResult, err
	}

	return resp, result, errorResult, nil
}

func GetRequest[T any](client *resty.Client, url string, queryParams map[string]string, token string) (*resty.Response, T, interface{}, error) {
	var result T
	var errorResult interface{}
	req := client.R().
		SetHeader("Authorization", "Bearer "+token).
		SetResult(&result).
		SetError(&errorResult)

	if len(queryParams) > 0 {
		req = req.SetQueryParams(queryParams)
	}

	resp, err := req.Get(url)
	if err != nil {
		return nil, result, errorResult, err
	}
	return resp, result, errorResult, nil
}

func PutRequest[T any](client *resty.Client, url string, data interface{}, queryParams map[string]string, token string) (*resty.Response, T, interface{}, error) {
	var result T
	var errorResult interface{}
	req := client.R().
		SetHeader("Authorization", "Bearer "+token).
		SetError(&errorResult).
		SetBody(data).
		SetResult(&result)

	if len(queryParams) > 0 {
		req = req.SetQueryParams(queryParams)
	}

	resp, err := req.Put(url)

	if err != nil {
		return nil, result, errorResult, err
	}
	return resp, result, errorResult, nil
}

func DeleteRequest[T any](client *resty.Client, url string, data interface{}, queryParams map[string]string, token string) (*resty.Response, T, interface{}, error) {
	var result T
	var errorResult interface{}
	req := client.R().
		SetHeader("Authorization", "Bearer "+token).
		SetError(&errorResult).
		SetBody(data).
		SetResult(&result)

	if len(queryParams) > 0 {
		req = req.SetQueryParams(queryParams)
	}

	resp, err := req.Delete(url)
	if err != nil {
		return nil, result, errorResult, err
	}

	return resp, result, errorResult, nil
}

func PatchRequest[T any](client *resty.Client, url string, data interface{}, queryParams map[string]string, token string) (*resty.Response, T, interface{}, error) {
	var result T
	var errorResult interface{}
	req := client.R().
		SetHeader("Authorization", "Bearer "+token).
		SetError(&errorResult).
		SetBody(data).
		SetResult(&result)

	if len(queryParams) > 0 {
		req = req.SetQueryParams(queryParams)
	}

	resp, err := req.Patch(url)
	if err != nil {
		return nil, result, errorResult, err
	}
	return resp, result, errorResult, nil
}

func CheckResponse(resp *resty.Response) error {
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return nil
}
