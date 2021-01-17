package blockchain

import "testing"

//Test_GetBlockchain_Nil : GetBlockchain should return at-least one element even if the blockchain is nil.
func Test_GetBlockchain_Nil(t *testing.T) {
	var chainState = &ChainState{TransactionValidator: testBlockTransactionValidator{}}
	// init input
	var currentBlockchain []*Block

	// get blockchain
	result := chainState.GetBlockchain(currentBlockchain)

	// validate returned blockchain
	if len(result) == 0 {
		t.Errorf("the blockchain should have one block even if the input is nil")
	}
}

//Test_GetBlockchain_Empty : GetBlockchain should return at-least one element even if the blockchain is empty.
func Test_GetBlockchain_Empty(t *testing.T) {
	var chainState = &ChainState{TransactionValidator: testBlockTransactionValidator{}}
	// init input
	var currentBlockchain = []*Block{}

	// get blockchain
	result := chainState.GetBlockchain(currentBlockchain)

	// validate returned blockchain
	if len(result) == 0 {
		t.Errorf("the blockchain should have one block even if the input is empty")
	}
}

//Test_GetBlockchain_NotChanged : GetBlockchain should return the input if there are more elements in it already.
func Test_GetBlockchain_NotChanged(t *testing.T) {
	var chainState = &ChainState{TransactionValidator: testBlockTransactionValidator{}}
	// init input
	var currentBlockchain = []*Block{
		{
			Index: 1,
			Hash:  "",
		},
	}

	// get blockchain
	result := chainState.GetBlockchain(currentBlockchain)

	// validate returned blockchain
	if len(result) != 1 {
		t.Errorf("the blockchain must contain the original block and nothing more")
	}
}
