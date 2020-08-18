package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func gamePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "lobby.html", nil)
	}
}
