package blockchain

import "time"

//Blockchain : Current block chain.
var currentBlockchain []*Block

//SelectChain : Checks if new chain should be replaced with the existing one. The longest valid chain is always selected. Others are ignored.
func SelectChain(newChain []*Block, existingChain []*Block, currentTimestamp time.Time) ([]*Block, bool) {
	if newChain == nil {
		return existingChain, false
	}
	if existingChain == nil {
		return newChain, true
	}

	if IsValidChain(newChain, currentTimestamp) && len(newChain) > len(existingChain) {
		return newChain, true
	}

	return existingChain, false
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

//AddBlockToChain : Adds block to the blockchain.
func AddBlockToChain(latestBlock *Block, newBlock *Block, currentBlockchain []*Block, currentTimestamp time.Time) bool {

	// check if block is valid and if it is add it
	if IsValidNewBlock(newBlock, latestBlock, currentTimestamp) {
		SetBlockchain(append(currentBlockchain, newBlock))
		return true
	}

	return false
}
