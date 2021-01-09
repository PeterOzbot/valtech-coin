package wallet

import (
	"naivecoin/transactions"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Test_GetBalance_CorrectAmount : Tests if getting balance returns correct amount when there are some unspent outputs.
func Test_GetUnspent(t *testing.T) {

	var address = "address-1"
	var unspentTransactionOutputs = []*transactions.UnspentTransactionOutput{
		{
			Address: "address-1",
			Amount:  1,
		},
		{
			Address: "address-2",
			Amount:  2,
		},
		{
			Address: "address-1",
			Amount:  3,
		},
		{
			Address: "address-3",
			Amount:  4,
		},
	}
	var expectedUnspentOutputs = []*transactions.UnspentTransactionOutput{
		{
			Address: "address-1",
			Amount:  1,
		},
		{
			Address: "address-1",
			Amount:  3,
		},
	}

	// get unspent outputs
	actualUnspentOutputs := FilterUnspentTransactionOutput(unspentTransactionOutputs, address)

	// check result
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
