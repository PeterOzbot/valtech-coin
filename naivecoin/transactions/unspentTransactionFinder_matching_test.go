package transactions

import "testing"

// Test_IsMatch_NilTransactionInput : Tests if nil transaction input is handled correctly.
func Test_IsMatch_NilTransactionInput(t *testing.T) {

	// create inputs
	var transactionInput *TransactionInput
	var unspentTransactionOutput = &UnspentTransactionOutput{
		OutputID:    "outputID-1",
		OutputIndex: 2,
		Address:     "address",
		Amount:      100,
	}

	// sign
	result := transactionInput.IsMatch(unspentTransactionOutput)

	// validate
	if result {
		t.Errorf("result should be false when transaction input is nil")
	}
}

// Test_IsMatch_NilUnspentTransactionOutput : Tests if nil unspent transaction output is handled correctly.
func Test_IsMatch_NilUnspentTransactionOutput(t *testing.T) {

	// create inputs
	var transactionInput = &TransactionInput{
		OutputID:    "outputID",
		OutputIndex: 5,
		Signature:   "signature",
	}
	var unspentTransactionOutput *UnspentTransactionOutput

	// sign
	result := transactionInput.IsMatch(unspentTransactionOutput)

	// validate
	if result {
		t.Errorf("result should be false when unspent transaction output is nil")
	}
}

// Test_IsMatch_UnspentTransactionOutputMatches : Tests that when unspent transaction output matches transaction input the result is true.
func Test_IsMatch_UnspentTransactionOutputMatches(t *testing.T) {

	// create inputs
	var transactionInput = &TransactionInput{
		OutputID:    "outputID",
		OutputIndex: 5,
		Signature:   "signature",
	}
	var unspentTransactionOutput = &UnspentTransactionOutput{
		OutputID:    "outputID",
		OutputIndex: 5,
		Address:     "address",
		Amount:      100,
	}

	// sign
	result := transactionInput.IsMatch(unspentTransactionOutput)

	// validate
	if !result {
		t.Errorf("result should be true when unspent transaction output matches")
	}
}

// Test_IsMatch_UnspentTransactionOutputNotMatches : Tests that when unspent transaction output does not matches transaction input the result is false.
func Test_IsMatch_UnspentTransactionOutputNotMatches(t *testing.T) {

	// create inputs
	var transactionInput = &TransactionInput{
		OutputID:    "outputID",
		OutputIndex: 5,
		Signature:   "signature",
	}
	var unspentTransactionOutput = &UnspentTransactionOutput{
		OutputID:    "outputID-1",
		OutputIndex: 10,
		Address:     "address",
		Amount:      100,
	}

	// sign
	result := transactionInput.IsMatch(unspentTransactionOutput)

	// validate
	if result {
		t.Errorf("result should be false when unspent transaction output does not matches")
	}
}
