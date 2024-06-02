package router

import (
	docs "github.com/brenommelo/gopportunities/docs"
	"github.com/brenommelo/gopportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {

	handler.InitializerHandler()

	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
