package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Health(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
