package types

// TotalStakeRow represents a single row of the total_stake table
type TotalStakeByTypeRow struct {
	Height     int64  `db:"height"`
	Role       int8   `db:"role"`
	TotalStake uint64 `db:"total_stake"`
}

// Equal tells whether v and w represent the same rows
func (v TotalStakeByTypeRow) Equal(w TotalStakeByTypeRow) bool {
	return v.Height == w.Height &&
		v.Role == w.Role &&
		v.TotalStake == w.TotalStake
}

// TotalStakeRow allows to build a new TotalStakeRow
func NewTotalStakeByTypeRow(
	height int64,
	role int8,
	totalStake uint64) TotalStakeByTypeRow {
	return TotalStakeByTypeRow{
		Height:     height,
		Role:       role,
		TotalStake: totalStake,
	}
}

// StakeRequirementsRow represents a single row of the stake_requirements table
type StakeRequirementsRow struct {
	Height       int64  `db:"height"`
	Role         int8   `db:"role"`
	Requirements uint64 `db:"requirements"`
}

// Equal tells whether v and w represent the same rows
func (v StakeRequirementsRow) Equal(w StakeRequirementsRow) bool {
	return v.Height == w.Height &&
		v.Role == w.Role &&
		v.Requirements == w.Requirements
}

// StakeRequirementsRow allows to build a new StakeRequirementsRow
func NewStakeRequirementsRow(
	height int64,
	role int8,
	requirements uint64) StakeRequirementsRow {
	return StakeRequirementsRow{
		Height:       height,
		Role:         role,
		Requirements: requirements,
	}
}

// WeeklyPayoutRow represents a single row of the weekly_payout table
type WeeklyPayoutRow struct {
	Height int64  `db:"height"`
	Payout uint64 `db:"payout"`
}

// Equal tells whether v and w represent the same rows
func (v WeeklyPayoutRow) Equal(w WeeklyPayoutRow) bool {
	return v.Height == w.Height &&
		v.Payout == w.Payout
}

// WeeklyPayoutRow allows to build a new WeeklyPayoutRow
func NewWeeklyPayoutRow(
	height int64,
	payout uint64) WeeklyPayoutRow {
	return WeeklyPayoutRow{
		Height: height,
		Payout: payout,
	}
}
