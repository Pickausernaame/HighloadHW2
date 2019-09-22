package handlers

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

// Хендлер-функция для генерации случайного значения
func Foo(c *gin.Context) {
	time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	c.Status(200)
}
