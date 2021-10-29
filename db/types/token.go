package types

// SupplyRow represents a single row of the supply table
type SupplyRow struct {
	OneRowId bool   `db:"one_row_id"`
	Height   uint64 `db:"height"`
	Supply   uint64 `db:"supply"`
}

// Equal tells whether v and w represent the same rows
func (v SupplyRow) Equal(w SupplyRow) bool {
	return v.OneRowId == w.OneRowId &&
		v.Height == w.Height &&
		v.Supply == w.Supply
}

// SupplyRow allows to build a new SupplyRow
func NewSupplyRow(
	height uint64,
	supply uint64) SupplyRow {
	return SupplyRow{
		OneRowId: true,
		Height:   height,
		Supply:   supply,
	}
}
