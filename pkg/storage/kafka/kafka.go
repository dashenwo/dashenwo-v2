package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/util/log"
	"sync"
)

var (
	once     sync.Once
	conf     Kafka
	Producer sarama.SyncProducer
	err      error
)

type Kafka struct {
	// 主机地址
	Hosts []string
	Topic string
}

func Init() (sarama.SyncProducer, error) {
	once.Do(func() {
		conf = Kafka{}
		err = config.Get("kafka").Scan(&conf)
		if err != nil {
			log.Fatal(err)
		}
		producerConf := sarama.NewConfig()
		// 等待服务器所有副本都保存成功后的响应
		producerConf.Producer.RequiredAcks = sarama.WaitForAll
		// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
		producerConf.Producer.Partitioner = sarama.NewRandomPartitioner
		// 是否等待成功和失败后的响应
		producerConf.Producer.Return.Successes = true
		// 是否等待成功和失败后的响应
		producerConf.Producer.Return.Errors = true

		// 超时时间
		producerConf.Producer.Timeout = 3
		// 使用给定代理地址和配置创建一个同步生产者
		Producer, err = sarama.NewSyncProducer(conf.Hosts, producerConf)
		if err != nil {
			log.Fatal(err)
		}
	})
	return Producer, err
}
