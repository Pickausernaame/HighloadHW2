package handlers

import (
	"fmt"
	"github.com/Pickausernaame/HighloadHW2/models"
	"github.com/gin-gonic/gin"
	"math/rand"
)

// Хендлер-функция для генерации случайного значения
func (h *Handler) Generate(c *gin.Context) {
	var response models.Generation
	val := rand.Intn(1000000000)
	id, err := h.Agregator.InsertGen(val)
	if err != nil {
		fmt.Print(err)
		c.Status(403)
		return
	}
	response.Id = id
	response.Data = val
	c.JSON(201, response)
}
