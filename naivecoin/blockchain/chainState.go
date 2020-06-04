package blockchain

//BlockChain : Current block chain.
var currentBlockChain []*Block

//SelectChain : Checks if new chain should be replaced with the existing one. The longest valid chain is always selected. Others are ignored.
func SelectChain(newChain []*Block, existingChain []*Block) []*Block {
	if newChain == nil {
		return existingChain
	}
	if existingChain == nil {
		return newChain
	}

	if IsValidChain(newChain) && len(newChain) > len(existingChain) {
		return newChain
	}

	return existingChain
}

// GetBlockChain : Returns current valid block chain.
func GetBlockChain() []*Block {
	if currentBlockChain == nil {
		currentBlockChain = append(currentBlockChain, GenesisBlock())
	}

	return currentBlockChain
}
