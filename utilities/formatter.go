package utilities

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type response struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type ApiError struct {
	Field   string `json:"field"`
	Tag     string `json:"type"`
	Message string `json:"message"`
}

type Field struct {
	Name  string
	Value string
}

func ApiResponse(message string, success bool, data interface{}) response {
	meta := meta{
		Message: message,
		Success: success,
	}

	apires := response{
		Meta: meta,
		Data: data,
	}

	return apires
}
func FormatBindError(err error) interface{} {
	var errors []ApiError

	for _, e := range err.(validator.ValidationErrors) {
		fmt.Println(e.Error())
		errors = append(errors, ApiError{
			Field:   e.Field(),
			Tag:     e.Tag(),
			Message: e.Error(),
		})
	}

	return gin.H{"errors": errors}
}
