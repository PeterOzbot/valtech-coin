package blockchain

import "time"

//GenerateBlock : Generates new block from data and latest block in the block chain.
func GenerateBlock(blockData string, latestBlock *Block, currentTimestamp int64) *Block {
	if latestBlock == nil {
		return nil
	}

	// create new block without its hash
	var newBlock = &Block{
		Index:        latestBlock.Index + 1,
		PreviousHash: latestBlock.Hash,
		Timestamp:    time.Unix(currentTimestamp, 0),
		Data:         blockData,
	}

	// calculate block hash
	newBlock.Hash = newBlock.CalculateHash()

	// return creates block
	return newBlock
}
