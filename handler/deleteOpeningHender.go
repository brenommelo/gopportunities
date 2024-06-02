package handler

import (
	"fmt"
	"net/http"

	"github.com/brenommelo/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHender(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParamter").Error())
		return
	}

	opening := schemas.Oppening{}
	//find Opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening not found for id: %s", id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Erro deleting openig with id: %s", id))
		return
	}

	sendSuccess(ctx, "delete-opening", opening)

}
