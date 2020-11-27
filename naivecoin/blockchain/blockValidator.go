package blockchain

//IsValidNewBlock : Checks if new block is valid regarding the latest block.
func IsValidNewBlock(newBlock *Block, latestBlock *Block) bool {
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

	// new block is valid
	return true
}

//IsValidChain : Validates the whole block chain, including the first genesis block and all the rest.
func IsValidChain(blockchain []*Block) bool {
	if blockchain == nil || len(blockchain) == 0 {
		return false
	}

	// lets validate genesis block first
	if !isGenesisBlockValid(blockchain[0]) {
		return false
	}

	// loop through chain and validate all blocks with each other
	for index := 1; index < len(blockchain); index++ {

		if !IsValidNewBlock(blockchain[index], blockchain[index-1]) {
			return false
		}
	}

	return true
}

func isGenesisBlockValid(block *Block) bool {
	// generate genesis block
	genesisBlock := GenesisBlock()

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
