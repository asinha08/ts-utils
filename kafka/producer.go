package kafka

import (
	"github.com/Shopify/sarama"
	"log"
	"time"
)

func SyncMessageProducer(brokerConfig *KafkaBroker, clientName string, msg *sarama.ProducerMessage) (partition int32, offset int64, err error) {
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 3
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms
	config.ClientID = clientName

	producer, err := sarama.NewSyncProducer(brokerConfig.HostList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}
	partition, offset, err = producer.SendMessage(msg)
	defer func() {
		_ = producer.Close()
	}()
	return partition, offset, err
}

func GetAsyncProducer(brokerConfig *KafkaBroker, clientName string) sarama.AsyncProducer {

	config := sarama.NewConfig()
	config.Producer.Retry.Max = 3
	config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms
	config.ClientID = clientName

	producer, err := sarama.NewAsyncProducer(brokerConfig.HostList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}

	// We will just log to STDOUT if we're not able to produce messages.
	// Note: messages will only be returned here after all retry attempts are exhausted.
	go func() {
		for err := range producer.Errors() {
			log.Fatalln("Failed to write access log entry:", err)
		}
	}()

	/*producer.Input() <- &sarama.ProducerMessage{
		Topic:     "",
		Key:       nil,
		Value:     nil,
		Headers:   nil,
		Metadata:  nil,
		Offset:    0,
		Partition: 0,
		Timestamp: time.Time{},
	}
	*/
	return producer
}
