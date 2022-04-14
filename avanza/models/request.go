package models

import (
	"net/http"
	"net/url"
)

// RequestOptions are used to configure client calls.
type RequestOptions struct {
	// Headers to apply to the request
	Headers http.Header
	// Query params to apply to the request
	QueryParams url.Values
}

// RequestOption changes the configuration of RequestOptions.
type RequestOption func(o *RequestOptions)

// Header sets a header as an option.
func Header(key, value string) RequestOption {
	return func(o *RequestOptions) {
		if o.Headers == nil {
			o.Headers = make(http.Header)
		}

		o.Headers.Add(key, value)
	}
}

// QueryParam sets a query param as an option.
func QueryParam(key, value string) RequestOption {
	return func(o *RequestOptions) {
		if o.QueryParams == nil {
			o.QueryParams = make(url.Values)
		}

		o.QueryParams.Add(key, value)
	}
}