package handler

import (
	"fmt"
	"net/http"

	"github.com/brenommelo/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("value", "result")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("value", "result")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateOpeningRespose struct {
	Message string                   `json:"message"`
	Data    schemas.OppeningResponse `json:"data"`
}

type UpdateOpeningRespose struct {
	Message string                   `json:"message"`
	Data    schemas.OppeningResponse `json:"data"`
}
type DeleteOpeningRespose struct {
	Message string                   `json:"message"`
	Data    schemas.OppeningResponse `json:"data"`
}
