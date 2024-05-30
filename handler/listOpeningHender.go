package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpeningHender(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "LIST Opening"})
}