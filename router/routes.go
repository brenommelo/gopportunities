package router

import (
	"github.com/brenommelo/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		v1.GET("/opening", func(ctx *gin.Context) {
			handler.ShowOpeningHender(ctx)
		})

		v1.POST("/opening", func(ctx *gin.Context) {
			handler.CreateOpeningHender(ctx)
		})

		v1.DELETE("/opening", func(ctx *gin.Context) {
			handler.DeleteOpeningHender(ctx)
		})

		v1.PUT("/opening", func(ctx *gin.Context) {
			handler.UpdateOpeningHender(ctx)
		})

		v1.GET("/openings", func(ctx *gin.Context) {
			handler.ListOpeningHender(ctx)
		})
	}
}
