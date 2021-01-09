package wallet

import (
	"naivecoin/transactions"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Test_GetUnspentOutputCombination_NoLeftOver : Tests if returned outputs match the amount the most.
func Test_GetUnspentOutputCombination_NoLeftOver(t *testing.T) {
	// inputs
	amount := 5.0
	var unspentTransactionOutputs = []*transactions.UnspentTransactionOutput{
		{
			Amount: 1,
		},
		{
			Amount: 2,
		},
		{
			Amount: 3,
		},
		{
			Amount: 4,
		},
	}
	var expectedUnspentOutputs = []*transactions.UnspentTransactionOutput{
		{
			Amount: 1,
		},
		{
			Amount: 4,
		},
	}

	// get balance
	actualUnspentOutputs, leftover := GetUnspentOutputCombination(unspentTransactionOutputs, amount)

	// check result
	if leftover != 0 {
		t.Errorf("left over amount should be zero")
	}
	// check unspent outputs
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

// Test_GetUnspentOutputCombination_LeftOver : Tests if returned outputs match the amount the most.
func Test_GetUnspentOutputCombination_LeftOver(t *testing.T) {
	// inputs
	amount := 5.0
	var unspentTransactionOutputs = []*transactions.UnspentTransactionOutput{
		{
			Amount: 8,
		},
		{
			Amount: 3,
		},
	}
	var expectedUnspentOutputs = []*transactions.UnspentTransactionOutput{
		{
			Amount: 8,
		},
	}

	// get balance
	actualUnspentOutputs, leftover := GetUnspentOutputCombination(unspentTransactionOutputs, amount)

	// check result
	if leftover != 3 {
		t.Errorf("left over amount should be 3")
	}
	// check unspent outputs
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

// Test_GetUnspentOutputCombination_AmountDoesNotSurpass : Tests if returned outputs total amount is not too small. Even if selecting outputs would be 'closer' to target amount, we must always select more than target amount.
func Test_GetUnspentOutputCombination_AmountDoesNotSurpass(t *testing.T) {
	// inputs
	amount := 5.0
	var unspentTransactionOutputs = []*transactions.UnspentTransactionOutput{
		{
			Amount: 8,
		},
		{
			Amount: 4,
		},
		{
			Amount: 7,
		},
	}
	var expectedUnspentOutputs = []*transactions.UnspentTransactionOutput{
		{
			Amount: 7,
		},
	}

	// get balance
	actualUnspentOutputs, leftover := GetUnspentOutputCombination(unspentTransactionOutputs, amount)

	// check result
	if leftover != 2 {
		t.Errorf("left over amount should be 2")
	}
	// check unspent outputs
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

// Test_GetUnspentOutputCombination_AmountDoesNotSurpass : Tests if result has the most unspent outputs possible.
func Test_GetUnspentOutputCombination_Longer(t *testing.T) {
	// inputs
	amount := 5.0
	var unspentTransactionOutputs = []*transactions.UnspentTransactionOutput{
		{
			Amount: 1.5,
		},
		{
			Amount: 2.5,
		},
		{
			Amount: 4,
		},
		{
			Amount: 1,
		},
		{
			Amount: 0.5,
		},
		{
			Amount: 0.5,
		},
	}
	var expectedUnspentOutputs = []*transactions.UnspentTransactionOutput{
		{
			Amount: 1.5,
		},
		{
			Amount: 2.5,
		},
		{
			Amount: 0.5,
		},
		{
			Amount: 0.5,
		},
	}

	// get balance
	actualUnspentOutputs, leftover := GetUnspentOutputCombination(unspentTransactionOutputs, amount)

	// check result
	if leftover != 0 {
		t.Errorf("left over amount should be 0")
	}
	// check unspent outputs
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
