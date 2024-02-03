package client

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/open-wallstreet/go-avanza/avanza/models"
	"net/http"
	"time"
)

const BaseUrl = "https://www.avanza.se"
const DefaultRetryCount = 3
const DefaultUserAgent = "Avanza GO API client"

// Client defines an HTTP client for the Avanza API.
type Client struct {
	HTTP       *resty.Client
	encoder    *Encoder
	AuthTokens *models.AuthSessionTokens
	Logger     Logger
}

func New() *Client {
	rClient := resty.New()
	rClient.SetBaseURL(BaseUrl)
	rClient.SetRetryCount(DefaultRetryCount)
	rClient.SetTimeout(10 * time.Second)
	rClient.SetHeaders(map[string]string{
		"Accept":       "*/*",
		"Content-Type": "application/json",
		"User-Agent":   DefaultUserAgent,
	})
	return &Client{
		HTTP:    rClient,
		encoder: NewEncoder(),
		Logger:  NewNoop(),
	}
}

// Call makes an API call based on the request params and options. The response is automatically unmarshalled.
func (c *Client) Call(ctx context.Context, method, path string, params, response interface{}, opts ...models.RequestOption) error {
	uri, err := c.encoder.EncodeParams(path, params)
	if err != nil {
		return err
	}
	return c.CallURL(ctx, method, uri, response, params, opts...)
}

// CallURL makes an API call based on a request URI and options. The response is automatically unmarshalled.
func (c *Client) CallURL(ctx context.Context, method, uri string, response, params interface{}, opts ...models.RequestOption) error {
	options := mergeOptions(opts...)

	req := c.HTTP.R().SetContext(ctx)
	req.SetQueryParamsFromValues(options.QueryParams)
	req.SetHeaderMultiValues(options.Headers)
	req.SetResult(response).SetError(&models.ErrorResponse{})
	if method == http.MethodPost || method == http.MethodPut {
		req.SetBody(params)
	}
	res, err := req.Execute(method, uri)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	} else if res.IsError() {
		errRes := res.Error().(*models.ErrorResponse)
		errRes.StatusCode = res.StatusCode()
		return errRes
	}

	return nil
}

func mergeOptions(opts ...models.RequestOption) *models.RequestOptions {
	options := &models.RequestOptions{}
	for _, o := range opts {
		o(options)
	}

	return options
}
