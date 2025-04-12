package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
  "ppe4peeps.com/order-api/services"
)

func setupRouter() *gin.Engine {
	api := gin.Default()

	api.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	api.POST("createOrder", func(c *gin.Context) {
		var newCreateOrderEvent services.OrderReceivedEvent

		if err := c.BindJSON(&newCreateOrderEvent); err != nil {
			return
		}


    if err := services.ProduceOrderEvent(newCreateOrderEvent); err != nil {
      return
    } 

		c.IndentedJSON(http.StatusCreated, newCreateOrderEvent)

	})

	return api
}

func main() {
	server := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	server.Run(":8080")
}
