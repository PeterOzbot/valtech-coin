package blockchain

import (
	"testing"
	"time"
)

//Test_ValidatesValidNewBlock : Tests if the validation handles nil input for the new block.
func Test_ValidatesNilNewBlock(t *testing.T) {
	// create example blocks
	var latestBlock *Block = &Block{}

	// validate block
	result := IsValidNewBlock(nil, latestBlock)

	// result should be false as new block is missing
	if result {
		t.Errorf("new block is nil and result should be false")
	}
}

//Test_ValidatesNilLatestBlock : Tests if the validation handles nil input for the latest block.
func Test_ValidatesNilLatestBlock(t *testing.T) {
	// create example blocks
	var newBlock *Block = &Block{}

	// validate block
	result := IsValidNewBlock(newBlock, nil)

	// result should be false as latest block is missing
	if result {
		t.Errorf("latest block is nil and result should be false")
	}
}

//Test_ValidatesIndex : Tests if the validation validates index on the new block.
func Test_ValidatesIndex(t *testing.T) {
	// create example blocks
	var newBlock *Block = &Block{
		Index:        0,
		Data:         "Najnovejši block je ta.",
		PreviousHash: "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb",
		Timestamp:    time.Unix(1588430083866862505, 0),
		Hash:         "4dfce9398b1e7a7dda79ff524de9d44859479d0019d7101c81b0d61393cfc11d",
	}
	var latestBlock *Block = &Block{
		Index:        15,
		Data:         "Dober dan gospod kamplan.",
		PreviousHash: "ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae",
		Timestamp:    time.Unix(1588430083866862500, 0),
		Hash:         "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb",
	}

	// validate block
	result := IsValidNewBlock(newBlock, latestBlock)

	// result should be false as Index is not correct
	if result {
		t.Errorf("new block is not valid and the validation result should be false - Index")
	}
}

//Test_ValidatesPreviousHash : Tests if the validation validates PreviousHash on the new block.
func Test_ValidatesPreviousHash(t *testing.T) {
	// create example blocks
	var newBlock *Block = &Block{
		Index:        16,
		Data:         "Najnovejši block je ta.",
		PreviousHash: "---",
		Timestamp:    time.Unix(1588430083866862505, 0),
		Hash:         "4dfce9398b1e7a7dda79ff524de9d44859479d0019d7101c81b0d61393cfc11d",
	}
	var latestBlock *Block = &Block{
		Index:        15,
		Data:         "Dober dan gospod kamplan.",
		PreviousHash: "ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae",
		Timestamp:    time.Unix(1588430083866862500, 0),
		Hash:         "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb",
	}

	// validate block
	result := IsValidNewBlock(newBlock, latestBlock)

	// result should be false as PreviousHash is not correct
	if result {
		t.Errorf("new block is not valid and the validation result should be false - PreviousHash")
	}
}

//Test_ValidatesHash : Tests if the validation validates Hash on the new block.
func Test_ValidatesHash(t *testing.T) {
	// create example blocks
	var newBlock *Block = &Block{
		Index:        16,
		Data:         "Najnovejši block je ta.",
		PreviousHash: "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb",
		Timestamp:    time.Unix(1588430083866862505, 0),
		Hash:         "---",
	}
	var latestBlock *Block = &Block{
		Index:        15,
		Data:         "Dober dan gospod kamplan.",
		PreviousHash: "ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae",
		Timestamp:    time.Unix(1588430083866862500, 0),
		Hash:         "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb",
	}

	// validate block
	result := IsValidNewBlock(newBlock, latestBlock)

	// result should be false as Hash is not correct
	if result {
		t.Errorf("new block is not valid and the validation result should be false - Hash")
	}
}

//Test_ValidatesValidNewBlock : Tests if the validation validates valid new block as valid.
func Test_ValidatesValidNewBlock(t *testing.T) {
	// create example blocks
	var newBlock *Block = &Block{
		Index:        16,
		Data:         "Najnovejši block je ta.",
		PreviousHash: "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb",
		Timestamp:    time.Unix(1588430083866862505, 0),
		Hash:         "4dfce9398b1e7a7dda79ff524de9d44859479d0019d7101c81b0d61393cfc11d",
	}
	var latestBlock *Block = &Block{
		Index:        15,
		Data:         "Dober dan gospod kamplan.",
		PreviousHash: "ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae",
		Timestamp:    time.Unix(1588430083866862500, 0),
		Hash:         "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb",
	}

	// validate block
	result := IsValidNewBlock(newBlock, latestBlock)

	// result should be true as the new block is valid
	if !result {
		t.Errorf("new block is valid and the validation result should be true")
	}
}
