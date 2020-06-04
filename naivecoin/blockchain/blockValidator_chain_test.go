package blockchain

import (
	"testing"
	"time"
)

func prepareChain() []*Block {
	// create genesis block
	genesisBlock := GenesisBlock()

	// create chain and add genesis block
	var chain []*Block
	chain = append(chain, genesisBlock)

	// return chain
	return chain
}

//Test_ValidatesChain_Empty : Tests if the chain validation validates empty chain.
func Test_ValidatesChain_Empty(t *testing.T) {
	// create null chain
	chain := make([]*Block, 0)

	// validate block
	result := IsValidChain(chain)

	// result should be negative as the chain is empty
	if result {
		t.Errorf("chain validation succeeded even if chain is empty")
	}
}

//Test_ValidatesChain_Null : Tests if the chain validation validates Null chain.
func Test_ValidatesChain_Null(t *testing.T) {
	// declare chain
	var chain []*Block

	// validate block
	result := IsValidChain(chain)

	// result should be negative as the chain is null
	if result {
		t.Errorf("chain validation succeeded even if chain is Null")
	}
}

//Test_ValidatesGenesisBlock : Tests if the chain validation validates valid genesis block.
func Test_ValidatesChain_GenesisBlock(t *testing.T) {
	// create testing chain with genesis block
	chain := prepareChain()

	// validate block
	result := IsValidChain(chain)

	// result should be positive
	if !result {
		t.Errorf("chain validation failed even if the blocks are valid")
	}
}

//Test_ValidatesChain_InvalidGenesisBlock : Tests if the chain validation validates invalid genesis block.
func Test_ValidatesChain_InvalidGenesisBlock(t *testing.T) {
	// create testing chain with genesis block
	chain := prepareChain()
	// break block
	chain[0].Hash = "asdasd"

	// validate block
	result := IsValidChain(chain)

	// result should be false as the genesis block is not valid
	if result {
		t.Errorf("chain validation succeeded even if the genesis block is invalid")
	}
}

//Test_ValidatesChain_SecondBlock : Tests if the chain validation validates second block.
func Test_ValidatesChain_SecondBlock(t *testing.T) {
	// create testing chain with genesis block
	chain := prepareChain()
	// create second block and add
	secondBlock := &Block{
		Index:        1,
		Data:         "My second block. HaHaHa",
		PreviousHash: chain[0].Hash,
		Timestamp:    time.Now(),
	}
	secondBlock.Hash = secondBlock.CalculateHash()
	chain = append(chain, secondBlock)

	// validate block
	result := IsValidChain(chain)

	// result should be positive
	if !result {
		t.Errorf("chain validation failed even if the blocks are valid")
	}
}

//Test_ValidatesChain_ThirdBlock : Tests if the chain validation validates third invalid block.
func Test_ValidatesChain_ThirdInvalidBlock(t *testing.T) {
	// create testing chain with genesis block
	chain := prepareChain()
	// create second block and add
	secondBlock := &Block{
		Index:        1,
		Data:         "My second block. HaHaHa",
		PreviousHash: chain[0].Hash,
		Timestamp:    time.Now(),
	}
	secondBlock.Hash = secondBlock.CalculateHash()
	chain = append(chain, secondBlock)
	// create third and add
	thirdBlock := &Block{
		Index:        3,
		Data:         "My third block. HaHaHa",
		PreviousHash: secondBlock.Hash,
		Timestamp:    time.Now(),
	}
	thirdBlock.Hash = "invalid hash for šur"
	chain = append(chain, thirdBlock)

	// validate block
	result := IsValidChain(chain)

	// result should be false
	if result {
		t.Errorf("chain validation succeeded even if the blocks are NOT valid")
	}
}