package fluentd

import "github.com/fluent/fluent-logger-golang/fluent"
import "fmt"

var logger *fluent.Fluent
var tag string

type TSLogConfig struct {
	FluentPort int
	FluentHost string
	TagPrefix  string
}

func SetFluentdLogger(tagName string, config *TSLogConfig) {
	l, err := fluent.New(fluent.Config{
		FluentPort: config.FluentPort,
		FluentHost: config.FluentHost,
		Async:      true,
		TagPrefix:  config.TagPrefix,
	})
	if err != nil {
		fmt.Println(err)
	}
	tag = tagName
	logger = l
}

func LogError(data map[string]string) {
	if logger == nil {
		fmt.Println("Logger not set")
		return
	}
	data["logType"] = "error"
	err := logger.Post(tag, data)
	if err != nil {
		panic(err)
	}
}

func LogInfo(data map[string]string) {
	if logger == nil {
		fmt.Println("Logger not set")
		return
	}
	data["logType"] = "info"
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
