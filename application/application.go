package application

import (
	"fmt"
	"github.com/Pickausernaame/AvitoRandom/handlers"
	"github.com/Pickausernaame/AvitoRandom/models"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router  *gin.Engine
	Handler *handlers.Handler
}

// Функция, создающая новое приложение
func New(config *models.Config) (*App, error) {

	a := &App{}
	// Создаем новый роутер
	a.Router = gin.New()

	// Подрубаем мидвары логер и рекавери
	a.Router.Use(gin.Logger())
	a.Router.Use(gin.Recovery())

	// Инициализируем хендлер и пробрасываем коннект к базе
	h, err := handlers.New(config)
	if err != nil {
		return nil, err
	}
	a.Handler = h

	// Объявляем эндпоинты и натравливаем на них хендлеры
	api := a.Router.Group("/api")
	{
		api.POST("/generate", a.Handler.Generate)
		api.GET("/retrieve/:id", a.Handler.Retrieve)
	}
	return a, nil
}

// Функция, запускающая приложение
func (a *App) Run(port int) error {
	// Запускаем роутер на переданном порту
	err := a.Router.Run(fmt.Sprintf(": %d", port))
	return err
}
