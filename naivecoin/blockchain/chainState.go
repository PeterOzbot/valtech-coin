package blockchain

//Blockchain : Current block chain.
var currentBlockchain []*Block

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

// GetBlockchain : Returns current valid block chain.
func GetBlockchain() []*Block {
	if currentBlockchain == nil {
		currentBlockchain = append(currentBlockchain, GenesisBlock())
	}

	return currentBlockchain
}

//SetBlockchain : sets block chain.
func SetBlockchain(blockchain []*Block) {
	currentBlockchain = blockchain
}

//GetLatestBlock : Returns latest block in the block chain.
func GetLatestBlock() *Block {
	blockchain := GetBlockchain()
	return blockchain[len(blockchain)-1]
}

//AddBlockToChain : Adds block to the blockchain.
func AddBlockToChain(newBlock *Block) bool {
	// get latest block
	latestBlock := GetLatestBlock()

	// check if block is valid and if it is add it
	if IsValidNewBlock(newBlock, latestBlock) {
		blockchain := GetBlockchain()
		SetBlockchain(append(blockchain, newBlock))
		return true
	} else {
		return false
	}
}
