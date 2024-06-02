package handler

import (
	"net/http"

	"github.com/brenommelo/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningHender(ctx *gin.Context) {
	openings := []schemas.Oppening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Erro listing openings")
		return
	}
	sendSuccess(ctx, "list-openings", openings)
}
