package blockchain

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

//Test_HandleNilInput : Tests if the generator does not crash if it gets nil as input.
func Test_HandleNilInput(t *testing.T) {
	// try to generate block
	result := GenerateBlock("empty", nil, 0)

	// the expected block should be equal to the result
	if result != nil {
		t.Errorf("new block generator should return nil")
	}
}

//Test_GeneratesCorrectBlock : Tests if the generator generates correct block from data and previous block.
func Test_GeneratesCorrectBlock(t *testing.T) {
	// create example block
	var latestBlock *Block = &Block{
		Index:        15,
		Data:         "Dober dan gospod kamplan.",
		PreviousHash: "ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae",
		Timestamp:    time.Unix(1588430083866862500, 0),
		Hash:         "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb",
	}
	newBlockData := "My new block data"
	currentTimestamp := int64(1588430083866862555)

	// create expected block
	var expectedBlock *Block = &Block{
		Index:        16,
		Data:         newBlockData,
		PreviousHash: "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb",
		Timestamp:    time.Unix(currentTimestamp, 0),
		Hash:         "079f8214cc7e47783b8ce6656de68002040c893afccebb99e7dfdc6df3729ccc",
	}

	// generate block
	result := GenerateBlock(newBlockData, latestBlock, currentTimestamp)

	// the expected block should be equal to the result
	if !cmp.Equal(result, expectedBlock) {
		t.Errorf("new block generator did not generate correct block")
	}
}
