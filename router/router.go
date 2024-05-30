package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//inicializa o router utilizando a configs default do gin
	router := gin.Default()
	//define rotas
	router.GET("/ping", func(c *gin.Context) { //esse funcao seria um handler
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//aqui toda a api
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
