package blockchain

import "time"

const (
	//BlockGenerationInterval : Defines how often a block should be found in seconds
	BlockGenerationInterval int = 100
	//DifficultyAdjustmentInterval : Defines how often the difficulty should adjust to the increasing or decreasing network hashrate
	DifficultyAdjustmentInterval int = 10
)

//GenesisBlock : Create first block of the blockchain - Genesis block.
func GenesisBlock() *Block {
	return &Block{
		Index:        0,
		Hash:         "b3aca084b966d9ff02364e3956d39b1e5bdbc268f85514338197e5da635159f5",
		Data:         "my genesis block!!",
		PreviousHash: "",
		Timestamp:    time.Unix(1465154705, 0),
		Nonce:        32,
		Difficulty:   0,
	}
}
