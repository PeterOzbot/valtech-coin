package blockchain

import (
	"naivecoin/transactions"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

//Test_SelectChain_NewChainNull : Tests if the new chain is ignored when null.
func Test_SelectChain_NewChainNull(t *testing.T) {
	var chainState = getTestChainState(nil)

	// create empty chain for testing
	existingChain := make([]*Block, 1)

	// validate block
	result, newSelected, _, _ := chainState.SelectChain(nil, existingChain, time.Unix(int64(0), 0))

	// result should not null as existing chain should be selected
	if result == nil || cap(result) != 1 {
		t.Errorf("when new chain is null the existing chain should be selected")
	}
	if newSelected {
		t.Errorf("flag should indicate that existing was selected")
	}
}

//Test_SelectChain_ExistingChainNull : Tests if the new chain is selected when existing chain is null
func Test_SelectChain_ExistingChainNull(t *testing.T) {
	var chainState = getTestChainState(nil)

	// create empty chain for testing
	newChain := make([]*Block, 1)

	// validate block
	result, newSelected, _, _ := chainState.SelectChain(newChain, nil, time.Unix(int64(0), 0))

	// result should not null as new chain should be selected
	if result == nil || cap(result) != 1 {
		t.Errorf("when existing chain is null the new chain should be selected")
	}
	if !newSelected {
		t.Errorf("flag should indicate that new was selected")
	}
}

//Test_SelectChain_InvalidLonger : Tests if the new invalid chain is not selected even if its longer.
func Test_SelectChain_InvalidLonger(t *testing.T) {
	var chainState = getTestChainState(nil)

	// create testing chain with genesis block
	existingChain := prepareChain(false)

	// create new invalid chain
	newChain := prepareChain(false)
	secondBlock := &Block{
		Index:        10,
		Message:      "My second block. HaHaHa",
		Transactions: []*transactions.Transaction{},
		PreviousHash: "random hash",
		Timestamp:    time.Now(),
	}
	newChain = append(newChain, secondBlock)

	// validate block
	result, newSelected, _, _ := chainState.SelectChain(newChain, existingChain, time.Unix(int64(0), 0))

	// the selected chain should be existing which has length of 1
	if result == nil || len(result) != 1 {
		t.Errorf("when new chain is longer but invalid the existing chain should be selected")
	}
	if newSelected {
		t.Errorf("flag should indicate that existing was selected")
	}
}

//Test_SelectChain_ValidShorter : Tests if the new valid chain is not selected when it is shorter.
func Test_SelectChain_ValidShorter(t *testing.T) {
	var chainState = getTestChainState(nil)

	// create testing chain with genesis block and add another
	existingChain := prepareChain(false)
	secondBlock := &Block{
		Index:        1,
		Message:      "My second block. HaHaHa",
		Hash:         "2a41081647ed81811cc09c4f6dabfb088eefb346a48bc1963f74d93e93d05fa9",
		PreviousHash: existingChain[0].Hash,
		Timestamp:    time.Unix(543673, 0),
	}

	existingChain = append(existingChain, secondBlock)

	// create new valid chain but shorter
	newChain := prepareChain(false)

	// validate block
	result, newSelected, _, _ := chainState.SelectChain(newChain, existingChain, time.Unix(int64(0), 0))

	// result should be existing chain with length of 2
	if result == nil || len(result) != 2 {
		t.Errorf("when new chain is longer but invalid the existing chain should be selected")
	}
	if newSelected {
		t.Errorf("flag should indicate that existing was selected")
	}
}

//Test_SelectChain_ValidLonger : Tests if the new valid longer chain is selected.
func Test_SelectChain_ValidLonger(t *testing.T) {
	var chainState = getTestChainState(nil)

	// create testing chain with genesis block
	existingChain := prepareChain(false)

	// create new valid chain but shorter
	newChain := prepareChain(false)
	secondBlock := &Block{
		Index:        1,
		Message:      "My second block. HaHaHa",
		Hash:         "2991e4b1783c24f733c4ef152a456a9a712b15171335beb5949d781070b13496",
		PreviousHash: newChain[0].Hash,
		Timestamp:    newChain[0].Timestamp,
		Difficulty:   2,
		Nonce:        29,
	}
	newChain = append(newChain, secondBlock)

	// validate block
	result, newSelected, _, _ := chainState.SelectChain(newChain, existingChain, newChain[0].Timestamp)

	// result should be new chain with length of 2
	if result == nil || len(result) != 2 {
		t.Errorf("when new chain is longer and valid it should be selected")
	}
	if !newSelected {
		t.Errorf("flag should indicate that new was selected")
	}
}

//Test_SelectChain_AccumulatedDifficulty : Tests if the new valid longer chain is not selected when having lower accumulated difficulty.
func Test_SelectChain_AccumulatedDifficulty(t *testing.T) {
	var chainState = getTestChainState(nil)

	// create testing chain with genesis block
	existingChain := prepareChain(false)
	existingChain[0].Difficulty = 2

	// create new valid chain but shorter
	newChain := prepareChain(false)
	secondBlock := &Block{
		Index:        1,
		Message:      "My second block. HaHaHa",
		Hash:         "c9f757c9c98283e47fb265767dfa1af79cd15459003ae8798535ed495f996e6c",
		PreviousHash: newChain[0].Hash,
		Timestamp:    newChain[0].Timestamp,
		Difficulty:   0,
		Nonce:        0,
	}
	newChain = append(newChain, secondBlock)

	// validate block
	result, newSelected, _, _ := chainState.SelectChain(newChain, existingChain, newChain[0].Timestamp)

	// result should be existing chain with length of 1
	if result == nil || len(result) != 1 {
		t.Errorf("when existing chain has greater accumulated difficulty it should be selected")
	}
	if newSelected {
		t.Errorf("flag should indicate that existing was selected")
	}
}

//Test_SelectChain_AccumulatedDifficulty : Tests if the new valid longer chain is selected when having higher accumulated difficulty.
func Test_SelectChain_AccumulatedDifficultyNew(t *testing.T) {
	var chainState = getTestChainState(nil)

	// create testing chain with genesis block
	existingChain := prepareChain(false)

	// create new valid chain
	newChain := prepareChain(false)
	secondBlock := &Block{
		Index:        1,
		Message:      "My second block. HaHaHa",
		Hash:         "28653dbf45c423ceda4afad9b26375bb47a0f41521c4b1e72420664406cbf1bb",
		PreviousHash: newChain[0].Hash,
		Timestamp:    newChain[0].Timestamp,
		Difficulty:   2,
		Nonce:        0,
	}
	newChain = append(newChain, secondBlock)

	// select chain
	result, newSelected, _, _ := chainState.SelectChain(newChain, existingChain, newChain[0].Timestamp)

	// result should be new chain with length of 2
	if result == nil || len(result) != 2 {
		t.Errorf("when new chain has greater accumulated difficulty it should be selected")
	}
	if !newSelected {
		t.Errorf("flag should indicate that new was selected")
	}
}

//Test_SelectChain_UnspentTransactionOutputs : Tests that when new chain is selected the unspent transaction outputs are returned correctly.
func Test_SelectChain_UnspentTransactionOutputs(t *testing.T) {
	// mock unspent transaction outputs
	expectedUnspentOutputs := []*transactions.UnspentTransactionOutput{
		{
			OutputID: "1",
		},
		{
			OutputID: "2",
		},
	}
	var chainState = getTestChainState(expectedUnspentOutputs)

	// create testing chain with genesis block
	existingChain := prepareChain(false)

	// create new valid chain but shorter
	newChain := prepareChain(false)
	secondBlock := &Block{
		Index:        1,
		Message:      "My second block. HaHaHa",
		Hash:         "28653dbf45c423ceda4afad9b26375bb47a0f41521c4b1e72420664406cbf1bb",
		PreviousHash: newChain[0].Hash,
		Timestamp:    newChain[0].Timestamp,
		Difficulty:   2,
		Nonce:        0,
	}
	newChain = append(newChain, secondBlock)

	// select chain
	_, _, actualUnspentOutputs, _ := chainState.SelectChain(newChain, existingChain, newChain[0].Timestamp)

	// validate
	if len(actualUnspentOutputs) != len(expectedUnspentOutputs) {
		t.Errorf("unspent transactions should be %d not %d", len(expectedUnspentOutputs), len(actualUnspentOutputs))
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
