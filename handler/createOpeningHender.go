package handler

import (
	"net/http"

	"github.com/brenommelo/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary CreateOpeningHender
// @Description create new job openinhg
// @Tags openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} CreateOpeningRespose
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHender(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Oppening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	logger.Infof("request recived: %+v", request)

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating oppening %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	sendSuccess(ctx, "create opening", opening)
}
