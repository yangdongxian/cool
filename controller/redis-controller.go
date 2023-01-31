package controller

import (
	"cool/common"
	"cool/config"
	"cool/dao"
	"cool/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IRedisQuery interface {
	RedisExecute(context *gin.Context)
}
type controllerRedisCli struct {
}

func NewControllerRedisCli() IRedisQuery {
	return &controllerRedisCli{}
}

func (r *controllerRedisCli) RedisExecute(context *gin.Context) {
	var parameter dao.RedisParameter
	err := context.ShouldBindJSON(&parameter)
	if err != nil {
		response := common.ErrorResponse(101, "Failed to process request", err.Error(), common.Empty{})
		context.JSON(http.StatusBadRequest, response)
	}

	cli := dao.NewRedisManger(parameter.Env, parameter.Name)
	fmt.Printf("------cli------- cli:%v\n", cli)
	for k, v := range config.RedisPool.Cli {
		fmt.Printf("------RedisPool------- key:%v value:%v \n", k, v)
	}

	key := toStr(parameter.Args[1])
	value := parameter.Args[2]
	//fmt.Printf("RedisExecute -- key:%s value:%v \n", key,value)
	ret, setErr := cli.Set(key, value)
	fmt.Printf("Set -- ret:%s \n", ret)
	if setErr != nil {
		response := common.ErrorResponse(101, "Redis Execute set is failed", err.Error(), common.Empty{})
		context.JSON(http.StatusBadRequest, response)
	}
	response := common.Response(100, "OK", ret)
	context.JSON(http.StatusOK, response)
}

// 转为string
func toStr(v interface{}) string {
	bt, _ := utils.Marshal(v)
	return string(bt)
}
