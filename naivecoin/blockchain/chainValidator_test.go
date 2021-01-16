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

//Test_IsValidChain_Empty : Tests if the chain validation validates empty chain.
func Test_IsValidChain_Empty(t *testing.T) {
	// create null chain
	chain := make([]*Block, 0)

	// validate block
	result, _, _ := IsValidChain(chain, time.Unix(int64(0), 0))

	// result should be negative as the chain is empty
	if result {
		t.Errorf("chain validation succeeded even if chain is empty")
	}
}

//Test_IsValidChain_Null : Tests if the chain validation validates Null chain.
func Test_IsValidChain_Null(t *testing.T) {
	// declare chain
	var chain []*Block

	// validate block
	result, _, _ := IsValidChain(chain, time.Unix(int64(0), 0))

	// result should be negative as the chain is null
	if result {
		t.Errorf("chain validation succeeded even if chain is Null")
	}
}

//Test_ValidatesGenesisBlock : Tests if the chain validation validates valid genesis block.
func Test_IsValidChain_GenesisBlock(t *testing.T) {
	// create testing chain with genesis block
	chain := prepareChain()

	// validate block
	result, _, _ := IsValidChain(chain, time.Unix(int64(0), 0))

	// result should be positive
	if !result {
		t.Errorf("chain validation failed even if the blocks are valid")
	}
}

//Test_IsValidChain_InvalidGenesisBlock : Tests if the chain validation validates invalid genesis block.
func Test_IsValidChain_InvalidGenesisBlock(t *testing.T) {
	// create testing chain with genesis block
	chain := prepareChain()
	// break block
	chain[0].Hash = "fake hash"

	// validate block
	result, _, _ := IsValidChain(chain, time.Unix(int64(0), 0))

	// result should be false as the genesis block is not valid
	if result {
		t.Errorf("chain validation succeeded even if the genesis block is invalid")
	}
}

//Test_IsValidChain_SecondBlock : Tests if the chain validation validates second block.
func Test_IsValidChain_SecondBlock(t *testing.T) {
	// current time
	var currentTimestamp = time.Unix(15884366862500, 0)
	// create testing chain with genesis block
	chain := prepareChain()
	// create second block and add
	secondBlock := &Block{
		Index:        1,
		Message:      "My second block. HaHaHa",
		PreviousHash: chain[0].Hash,
		Hash:         "37eca9944ecb7da3c1cd07786bdfe089593cedd118c5a57911e0354543ae5d10",
		Timestamp:    currentTimestamp,
	}
	chain = append(chain, secondBlock)

	// validate block
	result, _, _ := IsValidChain(chain, currentTimestamp)

	// result should be positive
	if !result {
		t.Errorf("chain validation failed even if the blocks are valid")
	}
}

//Test_IsValidChain_ThirdInvalidBlock : Tests if the chain validation validates third invalid block.
func Test_IsValidChain_ThirdInvalidBlock(t *testing.T) {
	// create testing chain with genesis block
	chain := prepareChain()
	// create second block and add
	secondBlock := &Block{
		Index:        1,
		Message:      "My second block. HaHaHa",
		PreviousHash: chain[0].Hash,
		Hash:         "218b2719e02aa1a330e8ed84eafede61ea1dee1eee7852c0c593053a43ae2e65",
		Timestamp:    time.Unix(242334443, 0),
	}
	chain = append(chain, secondBlock)

	// create third and add
	thirdBlock := &Block{
		Index:        3,
		Message:      "My third block. HaHaHa",
		PreviousHash: secondBlock.Hash,
		Timestamp:    time.Now(),
	}
	thirdBlock.Hash = "invalid hash for šur"
	chain = append(chain, thirdBlock)

	// validate block
	result, _, _ := IsValidChain(chain, time.Unix(int64(0), 0))

	// result should be false
	if result {
		t.Errorf("chain validation succeeded even if the blocks are NOT valid")
	}
}

//Test_IsValidChain_UnspentTransactionOutputs : Tests if the unspent transaction outputs are returned correctly.
func Test_IsValidChain_UnspentTransactionOutputs(t *testing.T) {
	t.Errorf("implement")
}
