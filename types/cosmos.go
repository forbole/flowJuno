package types

import (
	"fmt"
	"time"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
)

// Validator contains the data of a single validator
type Validator struct {
	ConsAddr   string
	ConsPubKey string
}

// NewValidator allows to build a new Validator instance
func NewValidator(consAddr string, consPubKey string) *Validator {
	return &Validator{
		ConsAddr:   consAddr,
		ConsPubKey: consPubKey,
	}
}

// -------------------------------------------------------------------------------------------------------------------

// CommitSig contains the data of a single validator commit signature
type CommitSig struct {
	Height           int64
	ValidatorAddress string
	VotingPower      int64
	ProposerPriority int64
	Timestamp        time.Time
}

// NewCommitSig allows to build a new CommitSign object
func NewCommitSig(validatorAddress string, votingPower, proposerPriority, height int64, timestamp time.Time) *CommitSig {
	return &CommitSig{
		Height:           height,
		ValidatorAddress: validatorAddress,
		VotingPower:      votingPower,
		ProposerPriority: proposerPriority,
		Timestamp:        timestamp,
	}
}

// -------------------------------------------------------------------------------------------------------------------

// Block contains the data of a single chain block
type Block struct {
	Height          int64
	Hash            string
	TxNum           int
	TotalGas        uint64
	ProposerAddress string
	Timestamp       time.Time
}

// NewBlock allows to build a new Block instance
func NewBlock(
	height int64, hash string, txNum int, totalGas uint64, proposerAddress string, timestamp time.Time,
) *Block {
	return &Block{
		Height:          height,
		Hash:            hash,
		TxNum:           txNum,
		TotalGas:        totalGas,
		ProposerAddress: proposerAddress,
		Timestamp:       timestamp,
	}
}

// NewBlockFromTmBlock builds a new Block instance from a given ResultBlock object
func NewBlockFromTmBlock(blk *tmctypes.ResultBlock, totalGas uint64) *Block {
	return NewBlock(
		blk.Block.Height,
		blk.Block.Hash().String(),
		len(blk.Block.Txs),
		totalGas,
		ConvertValidatorAddressToBech32String(blk.Block.ProposerAddress),
		blk.Block.Time,
	)
}


// Message represents the data of a single message
type Message struct {
	TxHash    string
	Index     int
	Type      string
	Value     string
	Addresses []string
}

// NewMessage allows to build a new Message instance
func NewMessage(txHash string, index int, msgType string, value string, addresses []string) *Message {
	return &Message{
		TxHash:    txHash,
		Index:     index,
		Type:      msgType,
		Value:     value,
		Addresses: addresses,
	}
}
