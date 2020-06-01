package blockchain

import (
	"testing"
)

//Test_SelectChain_NewChainNull : Tests if the new chain is ignored when null.
func Test_SelectChain_NewChainNull(t *testing.T) {
	// create empty chain for testing
	existingChain := make([]*Block, 1)

	// validate block
	result := SelectChain(nil, existingChain)

	// result should be false as new block is missing
	if result == nil || cap(result) != 1 {
		t.Errorf("when new chain is null the existing chain should be selected")
	}
}

//Test_SelectChain_ExistingChainNull : Tests if the new chain is selected when existing chain is null
func Test_SelectChain_ExistingChainNull(t *testing.T) {
	// create empty chain for testing
	newChain := make([]*Block, 1)

	// validate block
	result := SelectChain(newChain, nil)

	// result should be false as new block is missing
	if result == nil || cap(result) != 1 {
		t.Errorf("when existing chain is null the new chain should be selected")
	}
}
