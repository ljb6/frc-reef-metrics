package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	// Endpoint
	end := router.Group("/api/reef-metrics/v1")

	// Route to get all data from event_key
	end.GET("/event/:event_key", func(ctx *gin.Context) {
		event_key := ctx.Param("event_key")
		ctx.JSON(http.StatusOK, gin.H {
			"test": event_key,
		})
	})

	// Route to get the sum of each row from event_key
	end.GET("/event/:event_key/simple", func(ctx *gin.Context) {
		event_key := ctx.Param("event_key")
		ctx.JSON(http.StatusOK, gin.H {
			"test": event_key,
		})
	})
}