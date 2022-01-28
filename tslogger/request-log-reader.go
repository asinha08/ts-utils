package helper

import "net/http"

func RequestLogHelper(r *http.Request) (clientLang string, extraLogs map[string]interface{}) {
	clientLang = r.Header.Get("language")
	correlationId := r.Header.Get("x-correlation-id")

	extraLogs["clientLang"] = clientLang
	extraLogs["correlationId"] = correlationId
	return clientLang, extraLogs
}
