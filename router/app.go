package router

import (
	"cool/config"
	"cool/controller"
	"cool/dao"
	"cool/utils"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

var (
	mysqlOjb                            = dao.NewMysqlObject(config.DB)
	queryController                     = controller.NewQuery(mysqlOjb)
	brokerList                          = GetBrokerList()
	Producer        sarama.SyncProducer = config.SetupKafkaConnection(brokerList)
	kafkaController                     = controller.NewKafkaServer(Producer)

	redisController = controller.NewControllerRedisCli()
)

func InitApp() {
	server := gin.Default()
	config.DBConfs = InitDB()

	server.POST("/query", queryController.Query)
	server.POST("/execute", queryController.Execute)
	server.POST("/kafka", kafkaController.CollectQuery)

	server.POST("/redisDo", redisController.RedisExecute)

	server.Run(":9090")
}

func GetBrokerList() []string {
	//加载.env配置文件
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Unable to load .env file in SetupKafkaConnection func!")
	}
	brokers := os.Getenv("BROKERS")
	brokerList := strings.Split(brokers, ",")
	fmt.Println("brokerList:", brokerList)
	//log.Printf("Kafka brokers: %s", strings.Join(brokerList, ", "))
	return brokerList
}

// InitDB 初始化DB
func InitDB() (DBConfs *simplejson.Json) {
	var DBConfigPath = "/Users/cathy/dxyang/go/src/cool/config/dbconfig.json"
	dbConfs := utils.NewJSON().ReadJSON(DBConfigPath).Jsob
	//dbname := "testDb"
	//if ENV == "pro" {
	//	dbname = "cool"
	//}
	//testDB := dbConfs.Get("cooltest").Get(dbname)
	fmt.Println("------", dbConfs)
	return dbConfs
}
