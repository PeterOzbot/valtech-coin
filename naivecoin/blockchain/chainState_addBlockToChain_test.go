package blockchain

import (
	"testing"
	"time"
)

//Test_AddBlockToChain_InvalidBlock : Tests if result is false and block is not added when trying to add new invalid block.
func Test_AddBlockToChain_InvalidBlock(t *testing.T) {
	// create input
	var newBlock = &Block{
		Index: 1,
	}
	var latestBlock = &Block{
		Index: 0,
	}
	var currentBlockchain = []*Block{
		latestBlock,
	}
	var currentTimestamp = time.Unix(1588430083866862500, 0)

	// validate block
	isValidNewBlock, newBlockchain, _ := AddBlockToChain(latestBlock, newBlock, currentBlockchain, currentTimestamp)

	// check if block was added
	if isValidNewBlock {
		t.Errorf("when new chain is null the existing chain should be selected")
	}
	// check if blockchain did not get new block
	if len(newBlockchain) != 1 {
		t.Errorf("when new chain is null the existing chain should be selected")
	}
}

//Test_AddBlockToChain_ValidBlock : Tests if result is true and block is added when trying to add new valid block.
func Test_AddBlockToChain_ValidBlock(t *testing.T) {
	// create input
	var currentTimestamp = time.Unix(1588430083866862505, 0)
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

	var currentBlockchain = []*Block{
		latestBlock,
	}

	// validate block
	isValidNewBlock, newBlockchain, _ := AddBlockToChain(latestBlock, newBlock, currentBlockchain, currentTimestamp)

	// check if block was added
	if !isValidNewBlock {
		t.Errorf("when new chain is null the existing chain should be selected")
	}
	// check if blockchain did not get new block
	if len(newBlockchain) != 2 {
		t.Errorf("when new chain is null the existing chain should be selected")
	}
}
