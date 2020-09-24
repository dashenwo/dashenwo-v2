package global

import (
	"github.com/Shopify/sarama"
	"github.com/micro/go-micro/v2/client"
)

var (
	// elasticsearch客户端
	RequestClient client.Client
	Kafka         sarama.SyncProducer
)
