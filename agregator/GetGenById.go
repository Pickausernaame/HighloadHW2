package agregator

// Агрегирующая функция, которая забирает генерацию по id
func (agr *Aggregator) GetGenById(id int) (int, error) {
	var val int
	sql := `SELECT generation FROM generations WHERE id = $1;`
	err := agr.connection.QueryRow(sql, id).Scan(&val)
	return val, err
}
