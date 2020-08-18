package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func indexGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	}
}
