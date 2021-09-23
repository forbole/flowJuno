package types

import (
	"time"
)

// TotalStakeRow represents a single row of the total_stake table
type TotalStakeRow struct { 
	Height int64 `db:"height"`
	Role int8 `db:"role"`
	TotalStake uint64 `db:"total_stake"`
	Timestamp time.Time `db:"timestamp"`
  }
  
  // Equal tells whether v and w represent the same rows
  func (v TotalStakeRow) Equal(w TotalStakeRow)bool{
	return v.Height==w.Height && 
  v.Role==w.Role && 
  v.TotalStake==w.TotalStake && 
  v.Timestamp==w.Timestamp }
  
   // TotalStakeRow allows to build a new TotalStakeRow
  func NewTotalStakeRow( 
	height int64,
	role int8,
	totalStake uint64,
	timestamp time.Time) TotalStakeRow{
   return TotalStakeRow{
   Height:height,
   Role:role,
   TotalStake:totalStake,
   Timestamp:timestamp,
  }
  }
  
