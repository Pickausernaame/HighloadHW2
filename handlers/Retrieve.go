package handlers

import (
	"fmt"
	"github.com/Pickausernaame/HighloadHW2/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Хендлер-функция, возвращающая значение генерации по id
func (h *Handler) Retrieve(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print(err)
		c.Status(409)
		return
	}
	val, err := h.Agregator.GetGenById(id)
	if err != nil {
		fmt.Print(err)
		c.Status(404)
		return
	}
	c.JSON(200, models.Generation{Id: id, Data: val})
	return
}
