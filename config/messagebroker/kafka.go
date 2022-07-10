package messagebroker

import (
	"github.com/Shopify/sarama"
	"github.com/yanuar-nc/migration-database-microservice/config"
)

func GetKafkaConfig(c config.EventConfig) *sarama.Config {

	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.ClientID = c.ClientID
	config.Version = sarama.V0_10_0_0

	return config
}
