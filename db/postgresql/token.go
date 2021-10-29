package postgresql

func (db *Db) SaveSupply(supply uint64, height uint64) error {
	stmt := `INSERT INTO supply(height,supply) VALUES ($1,$2) ON CONFLICT (one_row_id) DO UPDATE 
    SET supply = excluded.supply,
    	height = excluded.height
WHERE supply.height <= excluded.height;`
	_, err := db.Sql.Exec(stmt,
		height,
		supply)
	return err
}
