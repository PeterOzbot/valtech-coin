package blockchain

import (
	"math"
	"naivecoin/transactions"
	"time"
)

//SelectChain : Checks if new chain should be replaced with the existing one. The longest valid chain is always selected. Others are ignored.
func (chainState *ChainState) SelectChain(newChain []*Block, existingChain []*Block, currentTimestamp time.Time) ([]*Block, bool, []*transactions.UnspentTransactionOutput, error) {
	if newChain == nil {
		return existingChain, false, []*transactions.UnspentTransactionOutput{}, nil
	}
	if existingChain == nil {
		return newChain, true, []*transactions.UnspentTransactionOutput{}, nil
	}

	// checks if chain is valid
	isValidChain, unspentTransactionOutputs, err := chainState.ChainValidator.IsValidChain(newChain, currentTimestamp)
	if !isValidChain || err != nil {
		return existingChain, false, []*transactions.UnspentTransactionOutput{}, err
	}

	// check if new has larger accumulated difficulty
	if !checkAccumulatedDifficulty(newChain, existingChain) {
		return existingChain, false, []*transactions.UnspentTransactionOutput{}, nil
	}

	return newChain, true, unspentTransactionOutputs, nil
}

// GetBlockchain : Returns current valid block chain.
func (chainState *ChainState) GetBlockchain(currentBlockchain []*Block) []*Block {
	if currentBlockchain == nil || len(currentBlockchain) == 0 {
		currentBlockchain = append(currentBlockchain, GenesisBlock())
	}

	return currentBlockchain
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
