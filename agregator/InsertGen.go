package agregator

// Агрегирующая функция, кладущая в базу новую генерацию.
func (agr *Aggregator) InsertGen(val int) (int, error) {
	var id int
	sql := `INSERT INTO generations (generation) VALUES ($1) RETURNING id;`
	err := agr.connection.QueryRow(sql, val).Scan(&id)
	return id, err
}
