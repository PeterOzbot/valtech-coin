package blockchain

import (
	"math"
	"time"
)

//SelectChain : Checks if new chain should be replaced with the existing one. The longest valid chain is always selected. Others are ignored.
func SelectChain(newChain []*Block, existingChain []*Block, currentTimestamp time.Time) ([]*Block, bool, error) {
	if newChain == nil {
		return existingChain, false, nil
	}
	if existingChain == nil {
		return newChain, true, nil
	}

	// checks if chain is valid
	isValidChain, err := IsValidChain(newChain, currentTimestamp)
	if !isValidChain || err != nil {
		return existingChain, false, err
	}

	// check if new has larger accumulated difficulty
	if !checkAccumulatedDifficulty(newChain, existingChain) {
		return existingChain, false, nil
	}

	return newChain, true, nil
}

// GetBlockchain : Returns current valid block chain.
func GetBlockchain(currentBlockchain []*Block) []*Block {
	if currentBlockchain == nil || len(currentBlockchain) == 0 {
		currentBlockchain = append(currentBlockchain, GenesisBlock())
	}

	return currentBlockchain
}

//AddBlockToChain : Adds block to the blockchain.
func AddBlockToChain(latestBlock *Block, newBlock *Block, currentBlockchain []*Block, currentTimestamp time.Time) (bool, []*Block, error) {

	// check if block is valid
	isValidNewBlock, err := IsValidNewBlock(newBlock, latestBlock, currentTimestamp)

	// if it is valid add it
	if isValidNewBlock && err == nil {
		newBlockChain := append(currentBlockchain, newBlock)
		return true, newBlockChain, nil
	}

	return false, currentBlockchain, err
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
