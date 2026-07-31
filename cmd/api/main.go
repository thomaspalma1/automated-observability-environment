package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/projeto-korp", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"nome":    "Projeto Korp",
			"horario": "00:00",
		})
	})

	router.Run(":8080")
}
