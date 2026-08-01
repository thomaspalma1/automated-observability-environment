package main

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// `httpRequestsTotal` tracks the total number of HTTP requests received by the service.
var httpRequestsTotal = promauto.NewCounter(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total de requisições HTTP recebidas pelo serviço",
	},
)

func main() {
	router := gin.Default()

	router.GET("/projeto-korp", func(context *gin.Context) {
		httpRequestsTotal.Inc()

		context.JSON(http.StatusOK, gin.H{
			"nome":    "Projeto Korp",
			"horario": time.Now().UTC().Format(time.RFC3339),
		})
	})

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.Run(":8080")
}
