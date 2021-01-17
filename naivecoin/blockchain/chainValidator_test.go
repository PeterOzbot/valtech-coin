package blockchain

import (
	"naivecoin/transactions"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

//Test_IsValidChain_Empty : Tests if the chain validation validates empty chain.
func Test_IsValidChain_Empty(t *testing.T) {
	var chainValidator = createValidator(true, nil)
	// create null chain
	chain := make([]*Block, 0)

	// validate block
	result, _, _ := chainValidator.IsValidChain(chain, time.Unix(int64(0), 0))

	// result should be negative as the chain is empty
	if result {
		t.Errorf("chain validation succeeded even if chain is empty")
	}
}

//Test_IsValidChain_Null : Tests if the chain validation validates Null chain.
func Test_IsValidChain_Null(t *testing.T) {
	var chainValidator = createValidator(true, nil)

	// declare chain
	var chain []*Block

	// validate block
	result, _, _ := chainValidator.IsValidChain(chain, time.Unix(int64(0), 0))

	// result should be negative as the chain is null
	if result {
		t.Errorf("chain validation succeeded even if chain is Null")
	}
}

//Test_ValidatesGenesisBlock : Tests if the chain validation validates valid genesis block.
func Test_IsValidChain_GenesisBlock(t *testing.T) {
	var chainValidator = createValidator(true, nil)

	// create testing chain with genesis block
	chain := prepareChain(false)

	// validate block
	result, _, _ := chainValidator.IsValidChain(chain, time.Unix(int64(0), 0))

	// result should be positive
	if !result {
		t.Errorf("chain validation failed even if the blocks are valid")
	}
}

//Test_IsValidChain_InvalidGenesisBlock : Tests if the chain validation validates invalid genesis block.
func Test_IsValidChain_InvalidGenesisBlock(t *testing.T) {
	var chainValidator = createValidator(true, nil)

	// create testing chain with genesis block
	chain := prepareChain(false)
	// break block
	chain[0].Hash = "fake hash"

	// validate block
	result, _, _ := chainValidator.IsValidChain(chain, time.Unix(int64(0), 0))

	// result should be false as the genesis block is not valid
	if result {
		t.Errorf("chain validation succeeded even if the genesis block is invalid")
	}
}

//Test_IsValidChain_SecondBlock : Tests if the chain validation validates second block.
func Test_IsValidChain_SecondBlock(t *testing.T) {
	var chainValidator = createValidator(true, nil)

	// current time
	var currentTimestamp = time.Unix(15884366862500, 0)
	// create testing chain with genesis block
	chain := prepareChain(false)
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
	result, _, _ := chainValidator.IsValidChain(chain, currentTimestamp)

	// result should be positive
	if !result {
		t.Errorf("chain validation failed even if the blocks are valid")
	}
}

//Test_IsValidChain_ThirdInvalidBlock : Tests if the chain validation validates third invalid block.
func Test_IsValidChain_ThirdInvalidBlock(t *testing.T) {
	var chainValidator = createValidator(true, nil)

	// create testing chain with genesis block
	chain := prepareChain(false)
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
	result, _, _ := chainValidator.IsValidChain(chain, time.Unix(int64(0), 0))

	// result should be false
	if result {
		t.Errorf("chain validation succeeded even if the blocks are NOT valid")
	}
}

//Test_IsValidChain_UnspentTransactionOutputsNotValid : Tests if chain is invalid when block transactions are invalid.
func Test_IsValidChain_UnspentTransactionOutputsNotValid(t *testing.T) {
	var chainValidator = createValidator(false, nil)

	// create testing chain with genesis block
	chain := prepareChain(true)

	// validate block
	result, _, _ := chainValidator.IsValidChain(chain, time.Unix(15884366862500, 0))

	// result should be false
	if result {
		t.Errorf("chain validation succeeded even if the block transactions are NOT valid")
	}
}

//Test_IsValidChain_UnspentTransactionOutputs : Tests if the unspent transaction outputs are returned correctly.
func Test_IsValidChain_UnspentTransactionOutputs(t *testing.T) {
	// create expected outputs
	var expectedUnspentOutputs = []*transactions.UnspentTransactionOutput{
		{
			OutputID: "1",
		},
		{
			OutputID: "2",
		}}
	var chainValidator = createValidator(true, expectedUnspentOutputs)

	// create testing chain with genesis block
	chain := prepareChain(true)

	// validate block
	_, actualUnspentOutputs, _ := chainValidator.IsValidChain(chain, time.Unix(15884366862500, 0))

	// validate
	if len(actualUnspentOutputs) != len(expectedUnspentOutputs) {
		t.Errorf("new unspent transactions should be %d not %d", len(expectedUnspentOutputs), len(actualUnspentOutputs))
		return
	}
	// validate each
	for index, resultUnspent := range actualUnspentOutputs {
		// find expected
		expectedResultUnspent := expectedUnspentOutputs[index]
		// validate expected
		if !cmp.Equal(resultUnspent, expectedResultUnspent) {
			t.Errorf("unspent transaction output does not match")
		}
	}
}
