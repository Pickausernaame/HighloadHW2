package application

import (
	"fmt"
	"github.com/Pickausernaame/HighloadHW2/handlers"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

type App struct {
	Router *gin.Engine
}

// Функция, создающая новое приложение
func New() (*App, error) {

	a := &App{}
	// Создаем новый роутер
	a.Router = gin.New()

	// Подрубаем мидвары логер и рекавери
	a.Router.Use(gin.Logger())
	a.Router.Use(gin.Recovery())

	p := ginprometheus.NewPrometheus("gin")

	p.Use(a.Router)

	// Объявляем эндпоинты и натравливаем на них хендлеры
	api := a.Router.Group("/api")
	{
		api.POST("/bar", handlers.Bar)
		api.GET("/foo", handlers.Foo)
		api.GET("/healthcheck", handlers.Healthcheck)
	}
	return a, nil
}

// Функция, запускающая приложение
func (a *App) Run(port int) error {
	// Запускаем роутер на переданном порту
	err := a.Router.Run(fmt.Sprintf(": %d", port))
	return err
}
