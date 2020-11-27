package blockchain

import (
	"testing"
)

// Test_GenesisBlock : Tests the hash from genesis block.
func Test_GenesisBlock(t *testing.T) {
	// create genesis block
	genesisBlock := GenesisBlock()

	// block data
	hash := "8d46e9fd83ead04cb38e6150130d556b97adf40cc865842395c1400ce48f724b"
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
