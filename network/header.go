package network

import "net/http"

func Reader(r *http.Request) (clientLang string, extraLogs map[string]interface{}) {
	clientLang = r.Header.Get("accept-language")
	correlationId := r.Header.Get("x-correlation-id")
	extraLogs = make(map[string]interface{})
	extraLogs["accept-language"] = clientLang
	extraLogs["x-correlation-id"] = correlationId
	return clientLang, extraLogs
}
