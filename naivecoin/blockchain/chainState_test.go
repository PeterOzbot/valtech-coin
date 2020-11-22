package blockchain

import (
	"testing"
	"time"
)

//Test_SelectChain_NewChainNull : Tests if the new chain is ignored when null.
func Test_SelectChain_NewChainNull(t *testing.T) {
	// create empty chain for testing
	existingChain := make([]*Block, 1)

	// validate block
	result, newSelected := SelectChain(nil, existingChain)

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
	// create empty chain for testing
	newChain := make([]*Block, 1)

	// validate block
	result, newSelected := SelectChain(newChain, nil)

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
	// create testing chain with genesis block
	existingChain := prepareChain()

	// create new invalid chain
	newChain := prepareChain()
	secondBlock := &Block{
		Index:        10,
		Data:         "My second block. HaHaHa",
		PreviousHash: "random hash",
		Timestamp:    time.Now(),
	}
	newChain = append(newChain, secondBlock)

	// validate block
	result, newSelected := SelectChain(newChain, existingChain)

	// the selected chain should be existing which has length of 1
	if result == nil || len(result) > 1 {
		t.Errorf("when new chain is longer but invalid the existing chain should be selected")
	}
	if newSelected {
		t.Errorf("flag should indicate that existing was selected")
	}
}

//Test_SelectChain_ValidShorter : Tests if the new valid chain is not selected when it is shorter.
func Test_SelectChain_ValidShorter(t *testing.T) {
	// create testing chain with genesis block and add another
	existingChain := prepareChain()
	secondBlock := &Block{
		Index:        1,
		Data:         "My second block. HaHaHa",
		PreviousHash: existingChain[0].Hash,
		Timestamp:    time.Now(),
	}
	secondBlock.Hash = secondBlock.CalculateHash()
	existingChain = append(existingChain, secondBlock)

	// create new valid chain but shorter
	newChain := prepareChain()

	// validate block
	result, newSelected := SelectChain(newChain, existingChain)

	// result should be existing chain with length of 2
	if result == nil || len(result) <= 1 {
		t.Errorf("when new chain is longer but invalid the existing chain should be selected")
	}
	if newSelected {
		t.Errorf("flag should indicate that existing was selected")
	}
}

//Test_SelectChain_ValidLonger : Tests if the new valid longer chain is selected.
func Test_SelectChain_ValidLonger(t *testing.T) {
	// create testing chain with genesis block and add another
	existingChain := prepareChain()

	// create new valid chain but shorter
	newChain := prepareChain()
	secondBlock := &Block{
		Index:        1,
		Data:         "My second block. HaHaHa",
		PreviousHash: newChain[0].Hash,
		Timestamp:    time.Now(),
	}
	secondBlock.Hash = secondBlock.CalculateHash()
	newChain = append(newChain, secondBlock)

	// validate block
	result, newSelected := SelectChain(newChain, existingChain)

	// result should be new chain with length of 2
	if result == nil || len(result) < 2 {
		t.Errorf("when new chain is longer and valid it should be selected")
	}
	if !newSelected {
		t.Errorf("flag should indicate that new was selected")
	}
}
