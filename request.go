package goavanza

import (
	"net/http"

	"github.com/monaco-io/request"
)

type RequestOptions struct {
	headers map[string]string
	cookies []*http.Cookie
	query   map[string]string
}

func (a *api) request(path string, method string, data map[string]interface{}, config RequestOptions) (string, *http.Response, error) {
	headers := map[string]string{
		"Accept":                  "*/*",
		"Content-Type":            "application/json",
		"User-Agent":              UserAgent,
		"X-SecurityToken":         a.xSecurityToken,
		"X-AuthenticationSession": a.totpSession.AuthenticationSession,
	}
	if config.headers != nil {
		for k, v := range config.headers {
			headers[k] = v
		}
	}
	if config.cookies == nil {
		config.cookies = []*http.Cookie{}
	}

	a.logger.Debugf("%s: %s", method, Url+path)

	c := request.Client{
		URL:     Url + path,
		Method:  method,
		JSON:    data,
		Header:  headers,
		Cookies: config.cookies,
		Query:   config.query,
	}
	resp := c.Send()
	if !resp.OK() {
		// handle error
		a.logger.Errorf("err: %v", resp.Error())
		return "", resp.Response(), resp.Error()
	}
	return resp.String(), resp.Response(), nil
}
