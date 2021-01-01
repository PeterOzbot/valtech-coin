package blockchain

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

//Test_GenerateBlock_HandleNilInput : Tests if the generator does not crash if it gets nil as input.
func Test_GenerateBlock_HandleNilInput(t *testing.T) {
	// try to generate block
	result := GenerateBlock("empty", nil, time.Unix(int64(0), 0), 0)

	// the expected block should be equal to the result
	if result != nil {
		t.Errorf("new block generator should return nil")
	}
}

//Test_GenerateBlock_GeneratesCorrectBlock : Tests if the generator generates correct block from data and previous block.
func Test_GenerateBlock_GeneratesCorrectBlock(t *testing.T) {
	// create example block
	var latestBlock *Block = &Block{
		Index:        15,
		Data:         "Dober dan gospod kamplan.",
		PreviousHash: "ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae",
		Timestamp:    time.Unix(1588430083866862500, 0),
		Hash:         "06dbca6186aa219b47be16de064d6180635f98a5b7a5b0cf2419ee0ae84c2ff2",
		Difficulty:   5,
		Nonce:        21,
	}
	newBlockData := "My new block data"
	currentTimestamp := time.Unix(int64(1588430083866862555), 0)

	// create expected block
	var expectedBlock *Block = &Block{
		Index:        16,
		Data:         newBlockData,
		PreviousHash: "06dbca6186aa219b47be16de064d6180635f98a5b7a5b0cf2419ee0ae84c2ff2",
		Timestamp:    currentTimestamp,
		Hash:         "bc98c44b998a09e3205103942a4184782f350dd87930221ff77a93e6cfb16a01",
		Difficulty:   5,
	}

	// generate block
	result := GenerateBlock(newBlockData, latestBlock, currentTimestamp, 5)

	// the expected block should be equal to the result
	if !cmp.Equal(result, expectedBlock) {
		t.Errorf("new block generator did not generate correct block")
	}
}
