package config

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/joho/godotenv"
)

func SetupKafkaConnection(brokerList []string) sarama.SyncProducer {
	//加载.env配置文件，读取配置数据库连接信息
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Unable to load .env file in SetupKafkaConnection func!")
	}
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		fmt.Println("producer closed ,error: ", err)
	}

	return producer
}
