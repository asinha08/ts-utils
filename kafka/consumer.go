package kafka

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
)

func KafkaGroupConsumerSetUp(topic *string, groupId *string, brokerConfig *KafkaBroker, handler ConsumerGroupHandler) {
	// Init config, specify appropriate version
	config := sarama.NewConfig()
	config.Version = sarama.V1_0_0_0
	config.Consumer.Return.Errors = true

	// Start with a client
	client, err := sarama.NewClient(brokerConfig.HostList, config)
	if err != nil {
		panic(err)
	}
	defer func() { _ = client.Close() }()

	// Start a new consumer group
	group, err := sarama.NewConsumerGroupFromClient(*groupId, client)
	if err != nil {
		panic(err)
	}
	defer func() { _ = group.Close() }()

	// Track errors
	go func() {
		for err := range group.Errors() {
			fmt.Println("ERROR", err)
		}
	}()

	// Iterate over consumer sessions.
	ctx := context.Background()
	for {
		topics := []string{*topic}
		err := group.Consume(ctx, topics, handler)
		if err != nil {
			panic(err)
		}
	}
}
