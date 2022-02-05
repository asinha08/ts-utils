package network

import (
	"net/http"
	"net/url"
)

func Reader(r *http.Request) (clientLang string, extraLogs map[string]interface{}) {
	clientLang = r.Header.Get("accept-language")
	correlationId := r.Header.Get("x-correlation-id")
	extraLogs = make(map[string]interface{})
	extraLogs["accept-language"] = clientLang
	extraLogs["x-correlation-id"] = correlationId
	return clientLang, extraLogs
}

func ReadQueryParams(r *http.Request) (queryParams url.Values, err error) {
	uri, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		return
	}
	queryParams, err = url.ParseQuery(uri.RawQuery)
	if err != nil {
		return
	}
	return
}
