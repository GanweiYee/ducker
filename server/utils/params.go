package utils

import "github.com/gin-gonic/gin"

type RequestParams struct {
	Page   int
	Rows   int
	params interface{}
}

func NewRequestParams(context *gin.Context) *RequestParams {
	return &RequestParams{}
}
