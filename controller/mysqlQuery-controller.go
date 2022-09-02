package controller

import (
	"cool/common"
	"cool/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IMysqlQuery interface {
	Query(context *gin.Context)
	Execute(context *gin.Context)
}

type mysqlObject struct {
	mysqlObj dao.IMysqlOperation
}

func NewQuery(db dao.IMysqlOperation) IMysqlQuery {
	return &mysqlObject{mysqlObj: db}
}
func (obj *mysqlObject) Query(context *gin.Context) {
	var parameter dao.QueryParameter
	err := context.ShouldBindJSON(&parameter)
	if err != nil {
		response := common.ErrorResponse(101, "Failed to process request", err.Error(), common.Empty{})
		context.JSON(http.StatusBadRequest, response)
	}

	ret := obj.mysqlObj.Query(parameter.Sql)
	//fmt.Printf("query:%#v\n",ret)
	response := common.Response(100, "OK", ret)
	context.JSON(http.StatusOK, response)
}
func (obj *mysqlObject) Execute(context *gin.Context) {
	var parameter dao.QueryParameter
	err := context.ShouldBindJSON(&parameter)
	if err != nil {
		response := common.ErrorResponse(101, "Failed to process request", err.Error(), common.Empty{})
		context.JSON(http.StatusBadRequest, response)
	}

	ret := obj.mysqlObj.Execute(parameter.Sql)

	ExecuteRet := map[string]int64{"rowsEffected": ret}
	response := common.Response(100, "OK", ExecuteRet)
	context.JSON(http.StatusOK, response)
}
