package fluentd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/asinha08/ts-utils/tsdate"
	"github.com/fluent/fluent-logger-golang/fluent"
)

var logger *fluent.Fluent
var tag string
var serviceName string

type TSLogConfig struct {
	FluentPort  int
	FluentHost  string
	TagPrefix   string
	Tag         string
	ServiceName string
}

func SetFluentdLogger(config *TSLogConfig) {
	l, err := fluent.New(fluent.Config{
		FluentPort:    config.FluentPort,
		FluentHost:    config.FluentHost,
		Async:         true,
		MarshalAsJSON: true,
		TagPrefix:     config.TagPrefix,
	})
	if err != nil {
		fmt.Println(err)
	}
	tag = config.Tag
	serviceName = config.ServiceName
	logger = l
}

func LogError(data map[string]interface{}, r *http.Request) {
	if logger == nil {
		fmt.Println("Logger not set")
		return
	}
	data["logType"] = "error"
	data["serviceName"] = serviceName
	now := time.Now().UTC()
	data["createdAt"] = tsdate.DateTimeTZFormatter(&now)
	if r != nil {
		data["REMOTE_ADDR"] = r.RemoteAddr
		data["METHOD"] = r.Method
		data["URI"] = r.RequestURI
		data["REFERER"] = r.Referer()
	}
	err := logger.Post(tag, data)
	if err != nil {
		panic(err)
	}
}

func LogInfo(data map[string]interface{}, r *http.Request) {
	if logger == nil {
		fmt.Println("Logger not set")
		return
	}
	data["logType"] = "info"
	data["serviceName"] = serviceName
	now := time.Now().UTC()
	data["createdAt"] = tsdate.DateTimeTZFormatter(&now)
	if r != nil {
		data["REMOTE_ADDR"] = r.RemoteAddr
		data["METHOD"] = r.Method
		data["URI"] = r.RequestURI
		data["REFERER"] = r.Referer()
	}
	err := logger.Post(tag, data)
	if err != nil {
		panic(err)
	}
}

func CloseLogger() {
	err := logger.Close()
	if err != nil {
		panic(err)
	}
}
