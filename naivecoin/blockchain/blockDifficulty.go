package blockchain

//GetDifficulty : Determines the correct difficulty for next block.
//difficulty increses if difference in timestamp between latest and adjusted block is less than half of expected
//difficulty decreases if difference in timestamp between latest and adjusted block is more than double of expected
//difficulty stays the same if difference in timestamp between latest and adjusted block is not less than half of expected or twice as large
//adjusted block is located with difficultyAdjustmentInterval
//expected time per block is blockGenerationInterval -> so time difference is difficultyAdjustmentInterval * blockGenerationInterval
func GetDifficulty(latestBlock *Block, currentBlockchain []*Block, blockGenerationInterval int, difficultyAdjustmentInterval int) int {

	// if adjustment interval is zero, each block should change difficulty
	if difficultyAdjustmentInterval == 0 {
		return latestBlock.Difficulty + 1
	}

	// check if difficulty should change
	if !shouldChangeDifficulty(currentBlockchain, difficultyAdjustmentInterval) {
		return latestBlock.Difficulty
	}

	// get adjusted block
	// reduce by additional 1 to get second to last
	// index out of range is checked before
	var adjustedBlock = currentBlockchain[len(currentBlockchain)-1-difficultyAdjustmentInterval]

	// calculate block difference
	var timeDifference = latestBlock.Timestamp.Unix() - adjustedBlock.Timestamp.Unix()

	// calculate expected time difference in seconds
	var expectedDifference = int64(blockGenerationInterval * difficultyAdjustmentInterval)

	// check if difficulty should be increased
	if timeDifference < expectedDifference/2 {
		// we increase difficulty if difference is smaller than half of expected
		return adjustedBlock.Difficulty + 1
	} else if timeDifference > expectedDifference*2 {
		// we decrease difficulty if difference is larger than double as expected
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

	// calculate difficulty only each difficultyAdjustmentInterval block
	var modRemain = (len(currentBlockchain) - 1) % difficultyAdjustmentInterval
	if modRemain != 0 {
		return false
	}

	return true
}
