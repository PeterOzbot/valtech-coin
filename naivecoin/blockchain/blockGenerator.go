package blockchain

import (
	"time"
)

//GenerateBlock : Generates new block from data and latest block in the block chain.
func GenerateBlock(blockData string, latestBlock *Block, currentTimestamp time.Time, difficulty int) *Block {
	if latestBlock == nil {
		return nil
	}

	// create new block without its hash
	var newBlock = &Block{
		Index:        latestBlock.Index + 1,
		PreviousHash: latestBlock.Hash,
		Timestamp:    currentTimestamp,
		Data:         blockData,
		Difficulty:   difficulty,
		Nonce:        0,
	}

	// calculate block hash
	newBlock.Hash = newBlock.CalculateHash()

	// return creates block
	return newBlock
}

//GetDifficulty : Determines the correct difficulty for next block.
//difficulty increses if difference in timestamp between latest and adjusted block is less than half of expected
//difficulty decreases if difference in timestamp between latest and adjusted block is more than double of expected
//difficulty stays the same if difference in timestamp between latest and adjusted block is not less than half of expected or twice as large
//adjusted block is located with difficultyAdjustmentInterval
//expected time per block is blockGenerationInterval -> so time difference is difficultyAdjustmentInterval * blockGenerationInterval
func GetDifficulty(latestBlock *Block, currentBlockchain []*Block, blockGenerationInterval int, difficultyAdjustmentInterval int) int {

	// check if difficulty should change
	if !shouldChangeDifficulty(currentBlockchain, difficultyAdjustmentInterval) {
		return latestBlock.Difficulty
	}

	// get adjusted block
	// reduce by additional 1 to get second to last
	var adjustedBlock = currentBlockchain[len(currentBlockchain)-difficultyAdjustmentInterval-1]

	// calculate block difference
	var timeDifference = latestBlock.Timestamp.Unix() - adjustedBlock.Timestamp.Unix()

	// calculate expected time difference in seconds
	var expectedDifference = int64(blockGenerationInterval * difficultyAdjustmentInterval)

	// check if difficulty should be increased
	if timeDifference < expectedDifference/2 {
		return adjustedBlock.Difficulty + 1
	} else if timeDifference > expectedDifference*2 {
		var newDifficulty = adjustedBlock.Difficulty - 1

		// check that we don't go lower that zero
		if newDifficulty > 0 {
			return newDifficulty
		}
		return 0
	} else {
		return adjustedBlock.Difficulty
	}
}

// checks if difficulty should even be changed
func shouldChangeDifficulty(currentBlockchain []*Block, difficultyAdjustmentInterval int) bool {

	// if only one block don't change difficulty
	if len(currentBlockchain) == 1 {
		return false
	}

	// if adjustment interval is zero, each block should change difficulty
	if difficultyAdjustmentInterval == 0 {
		return true
	}

	// calculate difficulty only each difficultyAdjustmentInterval block
	var modRemain = (len(currentBlockchain) - 1) % difficultyAdjustmentInterval
	if modRemain != 0 {
		return false
	}

	return true
}
