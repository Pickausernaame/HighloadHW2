package handlers

import (
	"github.com/Pickausernaame/HighloadHW2/agregator"
	"github.com/Pickausernaame/HighloadHW2/models"
	"github.com/jackc/pgx"
)

type Handler struct {
	Agregator *agregator.Aggregator
}

// Функция, создающая новый обработчик
func New(config *models.Config) (*Handler, error) {

	// Инициализация конфигурационных структур для соединения с бд
	conf := pgx.ConnConfig{
		User:      config.DBUser,
		Password:  config.DBPassword,
		Host:      config.DBHost,
		Port:      config.DBPort,
		Database:  config.DBSpace,
		TLSConfig: nil,
	}
	confPool := pgx.ConnPoolConfig{
		ConnConfig:     conf,
		MaxConnections: 8,
	}

	// Проброс коннекта к бд
	pool, err := pgx.NewConnPool(confPool)
	if err != nil {
		return nil, err
	}

	var h = &Handler{}
	h.Agregator = agregator.New(pool)
	return h, nil
}
