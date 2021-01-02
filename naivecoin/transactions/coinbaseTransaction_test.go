package transactions

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Test_CoinbaseTransaction_Correct : Tests create coinbase transaction is correct.
func Test_CoinbaseTransaction_Correct(t *testing.T) {

	// create coinbase transaction
	var address = "coinbase-address"
	var blockIndex = 12
	var expectedTransaction = &Transaction{
		ID: "ebafa7518cac709e160f201a888bdf3c969c36993eefbf852cc30c9eb1a553b8",
		Inputs: []*TransactionInput{
			{
				OutputID:    "",
				Signature:   "",
				OutputIndex: blockIndex,
			},
		},
		Outputs: []*TransactionOutput{
			{
				Address: address,
				Amount:  CoinbaseAmount,
			},
		},
	}

	// create coinbase transaction
	transaction := CoinbaseTransaction(address, blockIndex, CoinbaseAmount)

	// validate expected
	if !cmp.Equal(transaction, expectedTransaction) {
		t.Errorf("coinbase transaction was not created correctly")
	}
}
