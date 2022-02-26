package utilities

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type response struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ApiResponse(message string, code int, status string, data interface{}) response {
	meta := meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	apires := response{
		Meta: meta,
		Data: data,
	}

	return apires
}

func FormatBindError(err error) interface{} {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return gin.H{"errors": errors}
}
