package kafka

type KafkaBroker struct {
	HostList []string `json:"kafka_service_list"`
}
