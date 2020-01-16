package handler

func gameGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "lobby.html", nil)
	}
}