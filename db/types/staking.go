package types

import (
	"time"
)

// TotalStakeRow represents a single row of the total_stake table
type TotalStakeRow struct {
	Height     int64     `db:"height"`
	Role       int8      `db:"role"`
	TotalStake uint64    `db:"total_stake"`
	Timestamp  time.Time `db:"timestamp"`
}

// Equal tells whether v and w represent the same rows
func (v TotalStakeRow) Equal(w TotalStakeRow) bool {
	return v.Height == w.Height &&
		v.Role == w.Role &&
		v.TotalStake == w.TotalStake &&
		v.Timestamp == w.Timestamp
}

// TotalStakeRow allows to build a new TotalStakeRow
func NewTotalStakeRow(
	height int64,
	role int8,
	totalStake uint64,
	timestamp time.Time) TotalStakeRow {
	return TotalStakeRow{
		Height:     height,
		Role:       role,
		TotalStake: totalStake,
		Timestamp:  timestamp,
	}
}

// StakeRequirementsRow represents a single row of the stake_requirements table
type StakeRequirementsRow struct {
	Height       int64     `db:"height"`
	Role         int8      `db:"role"`
	Requirements uint64    `db:"requirements"`
	Timestamp    time.Time `db:"timestamp"`
}

// Equal tells whether v and w represent the same rows
func (v StakeRequirementsRow) Equal(w StakeRequirementsRow) bool {
	return v.Height == w.Height &&
		v.Role == w.Role &&
		v.Requirements == w.Requirements &&
		v.Timestamp == w.Timestamp
}

// StakeRequirementsRow allows to build a new StakeRequirementsRow
func NewStakeRequirementsRow(
	height int64,
	role int8,
	requirements uint64,
	timestamp time.Time) StakeRequirementsRow {
	return StakeRequirementsRow{
		Height:       height,
		Role:         role,
		Requirements: requirements,
		Timestamp:    timestamp,
	}
}
