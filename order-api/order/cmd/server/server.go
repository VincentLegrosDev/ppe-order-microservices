package server

import (
	"github.com/gin-gonic/gin"
	"ppe4peeps.com/order-api/models"
	"ppe4peeps.com/order-api/order/internal/handlers"
)

func SetupRouter() *gin.Engine {
	api := gin.Default()
	api.GET("/ping", handlers.Health)

	api.POST("orderReceived", func(c *gin.Context) {
		handlers.PublishOrderEvent(c, models.OrderReceived)
	})

	api.POST("orderConfirmed", func(c *gin.Context) {
		handlers.PublishOrderEvent(c, models.OrderConfirmed)
	})

	api.POST("orderPackedAndPicked", func(c *gin.Context) {
		handlers.PublishOrderEvent(c, models.OrderPackedAndPicked)
	})

	return api
}
