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

	// new block is valid
	return true
}
