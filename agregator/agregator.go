package agregator

import (
	"github.com/jackc/pgx"
)

type Aggregator struct {
	connection *pgx.ConnPool
}

// Функция, создающая новый агрегатор
func New(pool *pgx.ConnPool) *Aggregator {
	return &Aggregator{connection: pool}
}
