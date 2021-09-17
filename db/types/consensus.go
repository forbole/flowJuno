package types

import (
	"reflect"
	"time"
)

type GenesisRow struct {
	OneRowID      bool      `db:"one_row_id"`
	Time          time.Time `db:"time"`
	InitialHeight int64     `db:"initial_height"`
}

func NewGenesisRow(time time.Time, initialHeight int64) GenesisRow {
	return GenesisRow{
		OneRowID:      true,
		Time:          time,
		InitialHeight: initialHeight,
	}
}

func (r GenesisRow) Equal(s GenesisRow) bool {
	return r.Time.Equal(s.Time) &&
		r.InitialHeight == s.InitialHeight
}

// -------------------------------------------------------------------------------------------------------------------
// AverageTimeRow is the average block time each minute/hour/day
type AverageTimeRow struct {
	OneRowID    bool    `db:"one_row_id"`
	AverageTime float64 `db:"average_time"`
	Height      int64   `db:"height"`
}

func NewAverageTimeRow(averageTime float64, height int64) AverageTimeRow {
	return AverageTimeRow{
		OneRowID:    true,
		AverageTime: averageTime,
		Height:      height,
	}
}

// Equal return true if two AverageTimeRow are true
func (r AverageTimeRow) Equal(s AverageTimeRow) bool {
	return r.AverageTime == s.AverageTime &&
		r.Height == s.Height
}

// -------------------------------------------------------------------------------------------------------------------

// BlockRow represents a single block row stored inside the database
type BlockRow struct {
	Height               int64     `db:"height"`
	Id                   string    `db:"id"`
	ParentId             string    `db:"parent_id"`
	CollectionGuarantees []string  `db:"collection_guarantees"`
	Timestamp            time.Time `db:"timestamp"`
}

// Equal tells whether v and w represent the same rows
func (v BlockRow) Equal(w BlockRow) bool {
	return v.Height == w.Height &&
		v.Id == w.Id &&
		v.ParentId == w.ParentId &&
		reflect.DeepEqual(v.CollectionGuarantees, w.CollectionGuarantees) &&
		v.Timestamp.Equal(w.Timestamp)
}

// BlockRow allows to build a new BlockRow
func NewBlockRow(
	height int64,
	id string,
	parentId string,
	collectionGuarantees []string,
	timestamp time.Time) BlockRow {
	return BlockRow{
		Height:               height,
		Id:                   id,
		ParentId:             parentId,
		CollectionGuarantees: collectionGuarantees,
		Timestamp:            timestamp,
	}
}