package types

import (
	"time"
)

type StakeRequirements struct {
	Height     int64
	Role       uint8
	Requirement uint64
	Timestamp  time.Time
}

// Equal tells whether v and w represent the same rows
func (v StakeRequirements) Equal(w StakeRequirements) bool {
	return v.Height == w.Height &&
		v.Role == w.Role &&
		v.Requirement == w.Requirement &&
		v.Timestamp.Equal(w.Timestamp)
}

// StakeRequirements allows to build a new StakeRequirements
func NewStakeRequirements(
	height int64,
	role uint8,
	requirement uint64,
	timestamp time.Time) StakeRequirements {
	return StakeRequirements{
		Height:     height,
		Role:       role,
		Requirement: requirement,
		Timestamp:  timestamp,
	}
}

type TotalStake struct {
	Height     int64
	Role       int8
	TotalStake uint64
	Timestamp  time.Time
}

// Equal tells whether v and w represent the same rows
func (v TotalStake) Equal(w TotalStake) bool {
	return v.Height == w.Height &&
		v.Role == w.Role &&
		v.TotalStake == w.TotalStake &&
		v.Timestamp.Equal(w.Timestamp)
}

// TotalStake allows to build a new TotalStake
func NewTotalStake(
	height int64,
	role int8,
	totalStake uint64,
	timestamp time.Time) TotalStake {
	return TotalStake{
		Height:     height,
		Role:       role,
		TotalStake: totalStake,
		Timestamp:  timestamp,
	}
}
