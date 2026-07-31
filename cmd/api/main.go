package main

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/projeto-korp", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"nome":    "Projeto Korp",
			"horario": time.Now().UTC().Format("15:04:05"),
		})
	})

	router.Run(":8080")
}
