package transactions

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Test_FindUnspentTransactionOutput_EmptyUnspentTransactionOutputs : Tests if nil is returned when there are no unspent transaction outputs.
func Test_FindUnspentTransactionOutput_EmptyUnspentTransactionOutputs(t *testing.T) {

	// create inputs
	var transactionInput *TransactionInput
	var unspentTransactionOutputs []*UnspentTransactionOutput

	// sign
	result := FindUnspentTransactionOutput(unspentTransactionOutputs, transactionInput)

	// result should be nil
	if result != nil {
		t.Errorf("result is not nil when there are no unspent transaction outputs")
	}
}

// Test_FindUnspentTransactionOutput_TransactionInputNil : Tests if nil is returned when there is no transcation input.
func Test_FindUnspentTransactionOutput_TransactionInputNil(t *testing.T) {

	// create inputs
	var unspentTransactionOutputs []*UnspentTransactionOutput

	// sign
	result := FindUnspentTransactionOutput(unspentTransactionOutputs, nil)
	// result should be nil
	if result != nil {
		t.Errorf("result is not nil when there is no transaction input")
	}
}

// Test_FindUnspentTransactionOutput_NoMatchingUnspentTransactionOutput : Tests if nil is returned when there is no matching unspent output.
func Test_FindUnspentTransactionOutput_NoMatchingUnspentTransactionOutput(t *testing.T) {

	// create inputs
	var transactionInput = &TransactionInput{
		OutputID:    "outputID",
		OutputIndex: 5,
		Signature:   "signature",
	}
	var unspentTransactionOutputs = []*UnspentTransactionOutput{
		{
			OutputID:    "outputID-1",
			OutputIndex: 2,
			Address:     "address",
			Amount:      100,
		},
		{
			OutputID:    "outputID-2",
			OutputIndex: 4,
			Address:     "address-2",
			Amount:      10,
		},
		{
			OutputID:    "outputID-3",
			OutputIndex: 6,
			Address:     "address-3",
			Amount:      120,
		},
	}

	// sign
	result := FindUnspentTransactionOutput(unspentTransactionOutputs, transactionInput)

	// result should be nil
	if result != nil {
		t.Errorf("result is not nil when there is no mathing unspent transaction output")
	}
}

// Test_FindUnspentTransactionOutput_CorrectUnspentTransactionOutput : Tests if correct unspent output is returned.
func Test_FindUnspentTransactionOutput_CorrectUnspentTransactionOutput(t *testing.T) {

	// create inputs
	var transactionInput = &TransactionInput{
		OutputID:    "outputID",
		OutputIndex: 5,
		Signature:   "signature",
	}
	var unspentTransactionOutputs = []*UnspentTransactionOutput{
		{
			OutputID:    "outputID-1",
			OutputIndex: 2,
			Address:     "address",
			Amount:      100,
		},
		{
			OutputID:    "outputID",
			OutputIndex: 5,
			Address:     "address-2",
			Amount:      10,
		},
		{
			OutputID:    "outputID-3",
			OutputIndex: 6,
			Address:     "address-3",
			Amount:      120,
		},
	}
	var expectedResult = &UnspentTransactionOutput{
		OutputID:    "outputID",
		OutputIndex: 5,
		Address:     "address-2",
		Amount:      10,
	}

	// sign
	result := FindUnspentTransactionOutput(unspentTransactionOutputs, transactionInput)

	// validate expected
	if !cmp.Equal(result, expectedResult) {
		t.Errorf("unspent transaction output does not match")
	}
}
