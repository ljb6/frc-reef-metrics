package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	// endpoint
	end := router.Group("/api/reef-metrics/v1")

	// test
	end.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H {
			"message": "pong",
		})
	})
}