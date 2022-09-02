package common

import "strings"

type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type Empty struct{}

func Response(status int, msg string, data interface{}) response {
	return response{
		Status:  status,
		Message: msg,
		Error:   nil,
		Data:    data,
	}
}

func ErrorResponse(status int, msg string, error string, data interface{}) response {
	splitedError := strings.Split(error, "\n")
	return response{
		Status:  status,
		Message: msg,
		Error:   splitedError,
		Data:    data,
	}
}
