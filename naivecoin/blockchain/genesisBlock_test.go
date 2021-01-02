package blockchain

import (
	"testing"
)

// Test_GenesisBlock : Tests the hash and validity of the genesis block.
func Test_GenesisBlock(t *testing.T) {
	// create genesis block
	genesisBlock := GenesisBlock()

	// block data
	hash := genesisBlock.CalculateHash()
	previousHash := ""
	index := 0

	// values should match
	if hash != genesisBlock.Hash {
		t.Errorf("Hash is incorrect, Actual: %s Expected: %s", hash, genesisBlock.Hash)
	}
	if index != genesisBlock.Index {
		t.Errorf("Index is incorrect, Actual: %d Expected: %d", index, genesisBlock.Index)
	}
	if previousHash != genesisBlock.PreviousHash {
		t.Errorf("PreviousHash is incorrect, Actual: %s Expected: %s", previousHash, genesisBlock.PreviousHash)
	}
}
