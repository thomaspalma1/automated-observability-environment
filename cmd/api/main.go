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
		Help: "Total number of HTTP requests received by the service",
	},
)

func main() {
	router := gin.Default()

	router.GET("/status", func(context *gin.Context) {
		httpRequestsTotal.Inc()

		context.JSON(http.StatusOK, gin.H{
			"name": "http-server",
			"time": time.Now().UTC().Format(time.RFC3339),
		})
	})

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.Run(":8080")
}
