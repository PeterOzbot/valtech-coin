package blockchain

import (
	"testing"
	"time"
)

// Test_ValidatesValidNewBlock : Tests if the validation validates valid new block as valid.
func Test_ValidatesValidNewBlock(t *testing.T) {
	// create example blocks
	var newBlock *Block = &Block{
		index:        16,
		data:         "Najnovejši block je ta.",
		previousHash: "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb",
		timestamp:    time.Unix(1588430083866862505, 0),
		hash:         "4dfce9398b1e7a7dda79ff524de9d44859479d0019d7101c81b0d61393cfc11d",
	}
	var latestBlock *Block = &Block{
		index:        15,
		data:         "Dober dan gospod kamplan.",
		previousHash: "ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae",
		timestamp:    time.Unix(1588430083866862500, 0),
		hash:         "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb",
	}

	// hash block
	result := IsValidNewBlock(newBlock, latestBlock)

	// result should be true as the new block is valid
	if !result {
		t.Errorf("new block is valid")
	}
}
