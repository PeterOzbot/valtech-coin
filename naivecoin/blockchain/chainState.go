package blockchain

import (
	"math"
	"time"
)

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

	// checks if chain is valid
	if !IsValidChain(newChain, currentTimestamp) {
		return existingChain, false
	}

	// check if new has larger accumulated difficulty
	if !checkAccumulatedDifficulty(newChain, existingChain) {
		return existingChain, false
	}

	return newChain, true
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

// returns true if new chain has larger accumulated difficulty
func checkAccumulatedDifficulty(newChain []*Block, existingChain []*Block) bool {
	// calculate accumulated difficulty
	var newChainAccumulatedDifficulty = calculateAccumulatedDifficulty(newChain)
	var existingChainAccumulatedDifficulty = calculateAccumulatedDifficulty(existingChain)

	// return if new is larger
	return newChainAccumulatedDifficulty > existingChainAccumulatedDifficulty
}

func calculateAccumulatedDifficulty(chain []*Block) int64 {
	var accumulatedDifficulty int64 = 0
	for _, block := range chain {
		accumulatedDifficulty += int64(math.Pow(float64(block.Difficulty), 2))
	}
	return accumulatedDifficulty
}
