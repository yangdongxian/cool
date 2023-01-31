package controller

import (
	"cool/common"
	"cool/dao"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IKafkaServer interface {
	CollectQuery(context *gin.Context)
}
type kafkaServer struct {
	DataCollector sarama.SyncProducer
}

func NewKafkaServer(producer sarama.SyncProducer) IKafkaServer {
	return &kafkaServer{DataCollector: producer}
}

func (k *kafkaServer) CollectQuery(context *gin.Context) {
	var parameter dao.CollectQuery
	err := context.ShouldBindJSON(&parameter)
	if err != nil {
		response := common.ErrorResponse(101, "Failed to process request", err.Error(), common.Empty{})
		context.JSON(http.StatusBadRequest, response)
	}
	fmt.Println("CollectQuery:", parameter)
	partition, offset, err := k.DataCollector.SendMessage(&sarama.ProducerMessage{
		Topic: parameter.Topic,
		Value: sarama.StringEncoder(parameter.Value),
	})

	if err != nil {
		ret := fmt.Sprintf("Failed to store your data: %s", err)
		response := common.Response(100, "OK", ret)
		context.JSON(http.StatusOK, response)
	} else {
		//ret := fmt.Sprintf("Your data is stored with unique identifier important/%d/%d", partition, offset)
		response := common.Response(100, "OK", gin.H{"partition": partition, "offset": offset, "request": parameter})
		context.JSON(http.StatusOK, response)
	}
}
