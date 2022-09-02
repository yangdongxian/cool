package router

import (
	"cool/config"
	"cool/controller"
	"cool/dao"
	"github.com/gin-gonic/gin"
)

var (
	mysqlOjb        = dao.NewMysqlObject(config.DB)
	queryController = controller.NewQuery(mysqlOjb)
)

func InitApp() {
	server := gin.Default()

	server.POST("/query", queryController.Query)
	server.POST("/execute", queryController.Execute)
	server.Run(":9090")
}
