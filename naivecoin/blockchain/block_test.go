package blockchain

import (
	"testing"
)

// Test_GenesisBlock : Tests the hash from genesis block.
func Test_GenesisBlock(t *testing.T) {
	// create genesis block
	genesisBlock := GenesisBlock()

	// hash block
	result := genesisBlock.CalculateHash()

	// result should be correct hash
	if result != genesisBlock.hash {
		t.Errorf("hashing result is incorrect, Actual: %s Expected: %s", result, genesisBlock.hash)
	}
}
