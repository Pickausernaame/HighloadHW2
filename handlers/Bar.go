package handlers

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

// Хендлер-функция, возвращающая значение генерации по id
func Bar(c *gin.Context) {
	time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	c.Status(201)
}
