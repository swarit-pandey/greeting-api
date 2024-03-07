package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/greeting", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from server!",
		})
	})

	_ = router.Run(":8080")
}
