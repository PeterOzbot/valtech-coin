package blockchain

import "time"

//IsValidNewBlock : Checks if new block is valid regarding the latest block.
func IsValidNewBlock(newBlock *Block, latestBlock *Block, currentTimestamp time.Time) bool {
	if newBlock == nil || latestBlock == nil {
		return false
	}

	// validate index
	expectedIndex := latestBlock.Index + 1
	if newBlock.Index != expectedIndex {
		return false
	}

	// validate previous hash
	if newBlock.PreviousHash != latestBlock.Hash {
		return false
	}

	// validate current hash
	expectedHash := newBlock.CalculateHash()
	if newBlock.Hash != expectedHash {
		return false
	}

	// validates if hash matches difficulty
	if !newBlock.HashMatchesDifficulty() {
		return false
	}

	// validate timestamp
	if !newBlock.ValidateTimestamp(latestBlock, currentTimestamp) {
		return false
	}

	// new block is valid
	return true
}

//IsValidChain : Validates the whole block chain, including the first genesis block and all the rest.
func IsValidChain(blockchain []*Block, currentTimestamp time.Time) bool {
	if blockchain == nil || len(blockchain) == 0 {
		return false
	}

	// lets validate genesis block first
	if !isGenesisBlockValid(blockchain[0]) {
		return false
	}

	// loop through chain and validate all blocks with each other
	for index := 1; index < len(blockchain); index++ {

		if !IsValidNewBlock(blockchain[index], blockchain[index-1], currentTimestamp) {
			return false
		}
	}

	return true
}

func isGenesisBlockValid(block *Block) bool {
	// generate genesis block
	genesisBlock := GenesisBlock()

	// validate block values
	if genesisBlock.Data != block.Data {
		return false
	}

	if genesisBlock.Hash != block.Hash {
		return false
	}

	if genesisBlock.Index != block.Index {
		return false
	}

	if genesisBlock.PreviousHash != block.PreviousHash {
		return false
	}

	if genesisBlock.Timestamp != block.Timestamp {
		return false
	}

	return true
}

//ValidateTimestamp : Checks if block timestamp is not too much in the future or in the past.
func (block *Block) ValidateTimestamp(latestBlock *Block, currentTimestamp time.Time) bool {

	// block should not be in the future more that 60s
	return block.Timestamp.Unix() <= currentTimestamp.Unix()+60 &&
		// block should not be more that 60s in the past from latest block
		block.Timestamp.Unix() >= latestBlock.Timestamp.Unix()-60
}
