package handlers

import "github.com/gin-gonic/gin"

func Healthcheck(c *gin.Context) {
	c.Status(200)
}
