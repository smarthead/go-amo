package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ClientOptions struct {
	Url          string
	ClientId     string
	ClientSecret string
	AccessToken  string
	RefreshToken string
}

type Client struct {
	options *ClientOptions
	BaseUrl *url.URL
}

type requestOptions struct {
	HttpMethod string
	Body       interface{}
	Headers    map[string]string
}

type OauthTokenResponse struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewClient(options *ClientOptions) (*Client, error) {
	if len(options.AccessToken) == 0 || len(options.RefreshToken) == 0 || len(options.Url) == 0 {
		return nil, errors.New("AmoCrm: Invalid options")
	}

	resolvedUrl, err := url.Parse(options.Url)
	if err != nil {
		return nil, err
	}

	return &Client{
		options: options,
		BaseUrl: resolvedUrl,
	}, nil
}

func (api *Client) doRequest(resourceUrl string, requestParams requestOptions, result interface{}) error {
	resolvedUrl, err := url.Parse(resourceUrl)
	if err != nil {
		return err
	}

	requestUrl := api.BaseUrl.ResolveReference(resolvedUrl)

	requestBody := new(bytes.Buffer)
	if requestParams.Body != nil {
		encoder := json.NewEncoder(requestBody)
		if err := encoder.Encode(requestParams.Body); err != nil {
			return err
		}
	}

	request, err := http.NewRequest(requestParams.HttpMethod, requestUrl.String(), requestBody)
	if err != nil {
		return err
	}

	if requestParams.Headers != nil {
		for k, v := range requestParams.Headers {
			request.Header.Set(k, v)
		}
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return errors.New(fmt.Sprintf("Request error: %s %d %s %s", err.Error(), response.StatusCode, requestParams.HttpMethod, resourceUrl))
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		return errors.New(fmt.Sprintf("%s %d %s %s", string(bodyBytes), response.StatusCode, requestParams.HttpMethod, resourceUrl))
	}

	if result != nil {
		decoder := json.NewDecoder(response.Body)
		err = decoder.Decode(result)
		if err != nil {
			return err
		}
	}

	return nil
}

func (api *Client) RefreshToken() (*OauthTokenResponse, error) {
	result := new(OauthTokenResponse)
	request := map[string]string{
		"client_id":     api.options.ClientId,
		"client_secret": api.options.ClientSecret,
		"grant_type":    "refresh_token",
		"refresh_token": api.options.RefreshToken,
		"redirect_uri":  api.options.Url,
	}

	err := api.doRequest("/oauth2/access_token", requestOptions{
		HttpMethod: http.MethodPost,
		Body:       request,
		Headers: map[string]string{
			"Accept":        "application/json",
			"Cache-Control": "no-cache",
			"Content-Type":  "application/json",
		},
	}, result)

	return result, err
}

func (api *Client) Get(resource string, result interface{}) error {
	return api.doRequest(resource, requestOptions{
		HttpMethod: http.MethodGet,
		Body:       nil,
		Headers: map[string]string{
			"Accept":        "application/json",
			"Cache-Control": "no-cache",
			"Content-Type":  "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", api.options.AccessToken),
		},
	}, result)
}

func (api *Client) Post(resource string, request interface{}, result interface{}) error {
	return api.doRequest(resource, requestOptions{
		HttpMethod: http.MethodPost,
		Body:       request,
		Headers: map[string]string{
			"Accept":        "application/json",
			"Cache-Control": "no-cache",
			"Content-Type":  "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", api.options.AccessToken),
		},
	}, result)
}

func (api *Client) Patch(resource string, request interface{}, result interface{}) error {
	return api.doRequest(resource, requestOptions{
		HttpMethod: http.MethodPatch,
		Body:       request,
		Headers: map[string]string{
			"Accept":        "application/json",
			"Cache-Control": "no-cache",
			"Content-Type":  "application/json",
			"Authorization": fmt.Sprintf("Bearer %s", api.options.AccessToken),
		},
	}, result)
}
