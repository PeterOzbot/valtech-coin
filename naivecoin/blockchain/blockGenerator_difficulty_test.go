package blockchain

import (
	"testing"
	"time"
)

//Test_GetDifficulty_GenesisOnlyChain : Tests if correct difficulty is returned if only genesis block is in the blockchain.
func Test_GetDifficulty_GenesisOnlyChain(t *testing.T) {
	// intervals
	blockGenerationInterval := 0
	difficultyAdjustmentInterval := 0
	// create genesis block
	var genesisBlock = GenesisBlock()

	// create block chain
	var currentBlockchain = []*Block{genesisBlock}

	// get difficulty
	var difficulty = GetDifficulty(genesisBlock, currentBlockchain, blockGenerationInterval, difficultyAdjustmentInterval)

	// difficulty should be the same as genesis
	if difficulty != genesisBlock.Difficulty {
		t.Errorf("Difficulty with only genesis block should be the same as that of the genesis block.")
	}
}

//Test_GetDifficulty_AdjustmentIntervalLargerThanBlockchain : Tests if correct difficulty is returned if DifficultyAdjustmentInterval is larger that the number of blocks in block chain.
func Test_GetDifficulty_AdjustmentIntervalLargerThanBlockchain(t *testing.T) {
	// intervals
	blockGenerationInterval := 0
	difficultyAdjustmentInterval := 10
	// create genesis block
	var genesisBlock = GenesisBlock()

	// create block chain
	var currentBlockchain = []*Block{genesisBlock}

	// get difficulty
	var difficulty = GetDifficulty(genesisBlock, currentBlockchain, blockGenerationInterval, difficultyAdjustmentInterval)

	// difficulty should be the same as genesis
	if difficulty != genesisBlock.Difficulty {
		t.Errorf("Difficulty with only genesis block should be the same as that of the genesis block no matter if difficultyAdjustmentInterval is larger than the whole chain.")
	}
}

//Test_GetDifficulty_BelowZero : Tests if difficulty is not decreased below zero.
func Test_GetDifficulty_BelowZero(t *testing.T) {
	// intervals
	blockGenerationInterval := 10
	difficultyAdjustmentInterval := 1

	// timestamp
	adjustedTimestamp := int64(100)

	// latest timestamp - should be more that two times blockGenerationInterval
	latestTimestamp := adjustedTimestamp + int64(((difficultyAdjustmentInterval*blockGenerationInterval)*2)+1)

	// adjusted difficulty
	adjustedBlockDifficulty := 0

	// create genesis and adjusted block
	var genesisBlock = GenesisBlock()
	var adjustedBlock = &Block{
		Index:      1,
		Timestamp:  time.Unix(adjustedTimestamp, 0),
		Difficulty: adjustedBlockDifficulty,
	}
	var latestBlock = &Block{
		Index:      2,
		Timestamp:  time.Unix(latestTimestamp, 0),
		Difficulty: adjustedBlockDifficulty,
	}

	// create block chain
	var currentBlockchain = []*Block{genesisBlock, adjustedBlock, latestBlock}

	// get difficulty
	var difficulty = GetDifficulty(latestBlock, currentBlockchain, blockGenerationInterval, difficultyAdjustmentInterval)

	// difficulty should be the same as genesis
	if difficulty != adjustedBlockDifficulty {
		t.Errorf("Difficulty should be zero.")
	}
}

//Test_GetDifficulty_BelowZero : Tests if difficulty is calculated only on difficultyAdjustmentInterval-nt block.
// in this test the difficulty should not be calculated and result should be latest's difficulty
func Test_GetDifficulty_MultipleBlocks(t *testing.T) {
	// intervals
	blockGenerationInterval := 10
	difficultyAdjustmentInterval := 3

	// timestamp
	adjustedTimestamp := int64(100)

	// latest timestamp - should be more that two times blockGenerationInterval
	latestTimestamp := adjustedTimestamp + int64(((difficultyAdjustmentInterval*blockGenerationInterval)*2)+1)

	// create genesis and adjusted block
	var genesisBlock = GenesisBlock()
	var adjustedBlock = &Block{
		Index:      1,
		Timestamp:  time.Unix(adjustedTimestamp, 0),
		Difficulty: 2,
	}
	var latestBlock = &Block{
		Index:      2,
		Timestamp:  time.Unix(latestTimestamp, 0),
		Difficulty: 3,
	}

	// create block chain
	var currentBlockchain = []*Block{genesisBlock, adjustedBlock, latestBlock}

	// get difficulty
	var difficulty = GetDifficulty(latestBlock, currentBlockchain, blockGenerationInterval, difficultyAdjustmentInterval)

	// difficulty should be the same as genesis
	if difficulty != latestBlock.Difficulty {
		t.Errorf("Difficulty should be zero.")
	}
}

//Test_GetDifficulty_Increased : Tests if difficulty is increased when time between blocks is two times smaller than expected.
//difficulty increses if difference in timestamp between latest and adjusted block is less than half of expected
//adjusted block is located with difficultyAdjustmentInterval
//expected time is difficultyAdjustmentInterval*blockGenerationInterval
func Test_GetDifficulty_Increased(t *testing.T) {
	// intervals
	blockGenerationInterval := 10
	difficultyAdjustmentInterval := 1

	// timestamp
	adjustedTimestamp := int64(100)

	// latest timestamp - should be more that two times blockGenerationInterval
	latestTimestamp := adjustedTimestamp + int64((difficultyAdjustmentInterval*blockGenerationInterval/2)-1)

	// adjusted difficulty
	adjustedBlockDifficulty := 1

	// create genesis and adjusted block
	var genesisBlock = GenesisBlock()
	var adjustedBlock = &Block{
		Index:      1,
		Timestamp:  time.Unix(adjustedTimestamp, 0),
		Difficulty: adjustedBlockDifficulty,
	}
	var latestBlock = &Block{
		Index:      2,
		Timestamp:  time.Unix(latestTimestamp, 0),
		Difficulty: adjustedBlockDifficulty,
	}

	// create block chain
	var currentBlockchain = []*Block{genesisBlock, adjustedBlock, latestBlock}

	// get difficulty
	var difficulty = GetDifficulty(latestBlock, currentBlockchain, blockGenerationInterval, difficultyAdjustmentInterval)

	// difficulty should be the same as genesis
	if difficulty != adjustedBlockDifficulty+1 {
		t.Errorf("Difficulty should be increased.")
	}
}

//Test_GetDifficulty_Decreased : Tests if difficulty is decreased when time between blocks is two times larger than expected.
//difficulty decreases if difference in timestamp between latest and adjusted block is more than double of expected
//adjusted block is located with difficultyAdjustmentInterval
//expected time is difficultyAdjustmentInterval*blockGenerationInterval
func Test_GetDifficulty_Decreased(t *testing.T) {
	// intervals
	blockGenerationInterval := 10
	difficultyAdjustmentInterval := 1

	// timestamp
	adjustedTimestamp := int64(100)

	// latest timestamp - should be less that two times blockGenerationInterval
	latestTimestamp := adjustedTimestamp + int64(((difficultyAdjustmentInterval*blockGenerationInterval)*2)+1)

	// adjusted difficulty
	adjustedBlockDifficulty := 1

	// create genesis and adjusted block
	var genesisBlock = GenesisBlock()
	var adjustedBlock = &Block{
		Index:      1,
		Timestamp:  time.Unix(adjustedTimestamp, 0),
		Difficulty: adjustedBlockDifficulty,
	}
	var latestBlock = &Block{
		Index:      2,
		Timestamp:  time.Unix(latestTimestamp, 0),
		Difficulty: adjustedBlockDifficulty,
	}

	// create block chain
	var currentBlockchain = []*Block{genesisBlock, adjustedBlock, latestBlock}

	// get difficulty
	var difficulty = GetDifficulty(latestBlock, currentBlockchain, blockGenerationInterval, difficultyAdjustmentInterval)

	// difficulty should be the same as genesis
	if difficulty != adjustedBlockDifficulty-1 {
		t.Errorf("Difficulty should be decreased.")
	}
}

//Test_GetDifficulty_NoChanges : Tests if difficulty stays the same when time between blocks is the same as expected.
//difficulty stays the same if difference in timestamp between latest and adjusted block is not less than half of expected or twice as large
//adjusted block is located with difficultyAdjustmentInterval
//expected time is difficultyAdjustmentInterval*blockGenerationInterval
func Test_GetDifficulty_NoChanges(t *testing.T) {
	// intervals
	blockGenerationInterval := 10
	difficultyAdjustmentInterval := 1

	// timestamp
	adjustedTimestamp := int64(100)

	// latest timestamp - should be less that two times blockGenerationInterval
	latestTimestamp := adjustedTimestamp + int64((difficultyAdjustmentInterval * blockGenerationInterval))

	// adjusted difficulty
	adjustedBlockDifficulty := 1

	// create genesis and adjusted block
	var genesisBlock = GenesisBlock()
	var adjustedBlock = &Block{
		Index:      1,
		Timestamp:  time.Unix(adjustedTimestamp, 0),
		Difficulty: adjustedBlockDifficulty,
	}
	var latestBlock = &Block{
		Index:      2,
		Timestamp:  time.Unix(latestTimestamp, 0),
		Difficulty: adjustedBlockDifficulty,
	}

	// create block chain
	var currentBlockchain = []*Block{genesisBlock, adjustedBlock, latestBlock}

	// get difficulty
	var difficulty = GetDifficulty(latestBlock, currentBlockchain, blockGenerationInterval, difficultyAdjustmentInterval)

	// difficulty should be the same as genesis
	if difficulty != adjustedBlockDifficulty {
		t.Errorf("Difficulty should stay the same")
	}
}
