package types

import "time"

// Genesis contains the useful information about the genesis
type Genesis struct {
	Time          time.Time
	InitialHeight int64
	ChainId string
}

// NewGenesis allows to build a new Genesis instance
func NewGenesis(startTime time.Time, initialHeight int64,chainId string) *Genesis {
	return &Genesis{
		Time:          startTime,
		InitialHeight: initialHeight,
		ChainId:chainId,
	}
}

// Equal returns true iff g and other contain the same data
func (g *Genesis) Equal(other *Genesis) bool {
	return g.Time.Equal(other.Time) &&
		g.InitialHeight == other.InitialHeight&&
		g.ChainId ==other.ChainId
}
