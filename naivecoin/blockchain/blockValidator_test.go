package blockchain

import (
	"testing"
	"time"
)

//Test_IsValidNewBlock_ValidatesValidNewBlock : Tests if the validation handles nil input for the new block.
func Test_IsValidNewBlock_ValidatesNilNewBlock(t *testing.T) {
	// create example blocks
	var latestBlock *Block = &Block{}

	// validate block
	result, _ := IsValidNewBlock(nil, latestBlock, time.Unix(int64(0), 0))

	// result should be false as new block is missing
	if result {
		t.Errorf("new block is nil and result should be false")
	}
}

//Test_IsValidNewBlock_ValidatesNilLatestBlock : Tests if the validation handles nil input for the latest block.
func Test_IsValidNewBlock_ValidatesNilLatestBlock(t *testing.T) {
	// create example blocks
	var newBlock *Block = &Block{}

	// validate block
	result, _ := IsValidNewBlock(newBlock, nil, time.Unix(int64(0), 0))

	// result should be false as latest block is missing
	if result {
		t.Errorf("latest block is nil and result should be false")
	}
}

//Test_IsValidNewBlock_ValidatesIndex : Tests if the validation validates index on the new block.
func Test_IsValidNewBlock_ValidatesIndex(t *testing.T) {
	// create example blocks
	var newBlock *Block = &Block{
		Index:        0,
		Message:      "Najnovejši block je ta.",
		PreviousHash: "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb",
		Timestamp:    time.Unix(1588430083866862505, 0),
		Hash:         "4dfce9398b1e7a7dda79ff524de9d44859479d0019d7101c81b0d61393cfc11d",
	}
	var latestBlock *Block = &Block{
		Index:        15,
		Message:      "Dober dan gospod kamplan.",
		PreviousHash: "ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae",
		Timestamp:    time.Unix(1588430083866862500, 0),
		Hash:         "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb",
	}

	// validate block
	result, _ := IsValidNewBlock(newBlock, latestBlock, time.Unix(int64(0), 0))

	// result should be false as Index is not correct
	if result {
		t.Errorf("new block is not valid and the validation result should be false - Index")
	}
}

//Test_IsValidNewBlock_ValidatesPreviousHash : Tests if the validation validates PreviousHash on the new block.
func Test_IsValidNewBlock_ValidatesPreviousHash(t *testing.T) {
	// create example blocks
	var newBlock *Block = &Block{
		Index:        16,
		Message:      "Najnovejši block je ta.",
		PreviousHash: "---",
		Timestamp:    time.Unix(1588430083866862505, 0),
		Hash:         "4dfce9398b1e7a7dda79ff524de9d44859479d0019d7101c81b0d61393cfc11d",
	}
	var latestBlock *Block = &Block{
		Index:        15,
		Message:      "Dober dan gospod kamplan.",
		PreviousHash: "ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae",
		Timestamp:    time.Unix(1588430083866862500, 0),
		Hash:         "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb",
	}

	// validate block
	result, _ := IsValidNewBlock(newBlock, latestBlock, time.Unix(int64(0), 0))

	// result should be false as PreviousHash is not correct
	if result {
		t.Errorf("new block is not valid and the validation result should be false - PreviousHash")
	}
}

//Test_IsValidNewBlock_ValidatesHash : Tests if the validation validates Hash on the new block.
func Test_IsValidNewBlock_ValidatesHash(t *testing.T) {
	// create example blocks
	var newBlock *Block = &Block{
		Index:        16,
		Message:      "Najnovejši block je ta.",
		PreviousHash: "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb",
		Timestamp:    time.Unix(1588430083866862505, 0),
		Hash:         "---",
	}
	var latestBlock *Block = &Block{
		Index:        15,
		Message:      "Dober dan gospod kamplan.",
		PreviousHash: "ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae",
		Timestamp:    time.Unix(1588430083866862500, 0),
		Hash:         "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb",
	}

	// validate block
	result, _ := IsValidNewBlock(newBlock, latestBlock, time.Unix(int64(0), 0))

	// result should be false as Hash is not correct
	if result {
		t.Errorf("new block is not valid and the validation result should be false - Hash")
	}
}

//Test_IsValidNewBlock_ValidatesDifficulty : Tests if the validation validates Difficulty on the new block.
func Test_IsValidNewBlock_ValidatesDifficulty(t *testing.T) {
	// create example blocks
	var newBlock *Block = &Block{
		Index:        16,
		Message:      "Najnovejši block je ta.",
		PreviousHash: "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb",
		Timestamp:    time.Unix(1588430083866862505, 0),
		Hash:         "fc1f17f211942842cae5b581b4db0fcbc917f6c1c9ecae4b7552225cac4651e8",
		Difficulty:   6,
		Nonce:        40,
	}
	var latestBlock *Block = &Block{
		Index:        15,
		Message:      "Dober dan gospod kamplan.",
		PreviousHash: "ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae",
		Timestamp:    time.Unix(1588430083866862500, 0),
		Hash:         "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb",
	}

	// validate block
	result, _ := IsValidNewBlock(newBlock, latestBlock, time.Unix(int64(0), 0))

	// result should be false as Hash is not correct
	if result {
		t.Errorf("new block is not valid and the validation result should be false - Hash")
	}
}

//Test_IsValidNewBlock_ValidatesTimestamp : Tests if the validation validates timestamp on the new block regarding previous block.
// new block should not have its timestamp more that 60s in the past from latest block
func Test_IsValidNewBlock_ValidatesTimestampPreviousBlock(t *testing.T) {
	// create example blocks
	var newBlockTimestamp = int64(1588430083866862505)
	var newBlock *Block = &Block{
		Index:        16,
		Message:      "Najnovejši block je ta.",
		PreviousHash: "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb",
		Timestamp:    time.Unix(newBlockTimestamp, 0),
		Hash:         "02a1357876740fe77e0a934badad846902cc2f6861c706e185b68f4250ae53e2",
		Difficulty:   5,
		Nonce:        40,
	}
	var latestBlock *Block = &Block{
		Index:        15,
		Message:      "Dober dan gospod kamplan.",
		PreviousHash: "ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae",
		Timestamp:    time.Unix(newBlockTimestamp+61, 0),
		Hash:         "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb",
	}

	// validate block
	result, _ := IsValidNewBlock(newBlock, latestBlock, time.Unix(newBlockTimestamp, 0))

	// result should be false as Hash is not correct
	if result {
		t.Errorf("new block is not valid and the validation result should be false - Timestamp")
	}
}

//Test_IsValidNewBlock_ValidatesTimestamp : Tests if the validation validates timestamp on the new block regarding current time.
// if new block has its timestamp more that 60s in the future its invalid
func Test_IsValidNewBlock_ValidatesTimestampCurrentTime(t *testing.T) {
	// create example blocks
	var newBlockTimestamp = int64(1588430083866862505)
	var newBlock *Block = &Block{
		Index:        16,
		Message:      "Najnovejši block je ta.",
		PreviousHash: "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb",
		Timestamp:    time.Unix(newBlockTimestamp, 0),
		Hash:         "02a1357876740fe77e0a934badad846902cc2f6861c706e185b68f4250ae53e2",
		Difficulty:   5,
		Nonce:        40,
	}
	var latestBlock *Block = &Block{
		Index:        15,
		Message:      "Dober dan gospod kamplan.",
		PreviousHash: "ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae",
		Timestamp:    time.Unix(1588430083866862500, 0),
		Hash:         "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb",
	}

	// validate block
	result, _ := IsValidNewBlock(newBlock, latestBlock, time.Unix(newBlockTimestamp-61, 0))

	// result should be false as Hash is not correct
	if result {
		t.Errorf("new block is not valid and the validation result should be false - Timestamp")
	}
}

//Test_IsValidNewBlock_ValidatesValidNewBlock : Tests if the validation validates valid new block as valid.
func Test_IsValidNewBlock_ValidatesValidNewBlock(t *testing.T) {
	// current time
	var currentTimestamp = time.Unix(1588430083866862505, 0)
	// create example blocks
	var newBlock *Block = &Block{
		Index:        16,
		Message:      "Najnovejši block je ta.",
		PreviousHash: "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb",
		Timestamp:    currentTimestamp,
		Hash:         "0010d0af2526fbce2b6c23b434669748d6467d04bad4fa24cc303a18a77b41b1",
		Difficulty:   5,
		Nonce:        150,
	}
	var latestBlock *Block = &Block{
		Index:        15,
		Message:      "Dober dan gospod kamplan.",
		PreviousHash: "ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae",
		Timestamp:    time.Unix(1588430083866862500, 0),
		Hash:         "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb",
	}

	// validate block
	result, _ := IsValidNewBlock(newBlock, latestBlock, currentTimestamp)

	// result should be true as the new block is valid
	if !result {
		t.Errorf("new block is valid and the validation result should be true")
	}
}
