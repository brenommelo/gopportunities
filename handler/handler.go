package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHender(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "POST Opening"})
}

func ShowOpeningHender(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "GET Opening"})
}

func DeleteOpeningHender(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "DELETE Opening"})
}

func UpdateOpeningHender(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "PUT Opening"})
}

func ListOpeningHender(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "LIST Opening"})
}
