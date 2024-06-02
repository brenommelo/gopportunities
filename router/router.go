package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {

	//inicializa o router utilizando a configs default do gin
	router := gin.Default()

	//Iniciando rotas
	initializeRoutes(router)

	//aqui toda a api
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
