package transactions

import "testing"

// Test_CanSign_NilInput : Tests if nil input is handled without failing.
func Test_CanSign_NilInput(t *testing.T) {

	// prepare input
	var transactionInput *TransactionInput
	var unspentTransactions []*UnspentTransactionOutput
	var publicKey = ""

	// call can sign
	result := CanSign(unspentTransactions, transactionInput, publicKey)

	// result should false
	if result {
		t.Errorf("result of nil transaction should be false.")
	}
}

// Test_CanSign_NoUnspentTransactions : Tests if no unspent transations is handled without failing.
func Test_CanSign_NoUnspentTransactions(t *testing.T) {

	// prepare input
	var transactionInput *TransactionInput = &TransactionInput{
		OutputID:    "outID",
		OutputIndex: 10,
	}
	var unspentTransactions []*UnspentTransactionOutput
	var publicKey = ""

	// call can sign
	result := CanSign(unspentTransactions, transactionInput, publicKey)

	// result should false
	if result {
		t.Errorf("result when there are no unspent transactions should be false.")
	}
}

// Test_CanSign_NoUnspentTransactionMatches : Tests if unspent transaction can not be found.
func Test_CanSign_NoUnspentTransactionMatches(t *testing.T) {

	// prepare input
	var transactionInput *TransactionInput = &TransactionInput{
		OutputID:    "outID",
		OutputIndex: 10,
	}
	var unspentTransactions = []*UnspentTransactionOutput{
		{
			OutputID:    "outID1",
			OutputIndex: 100,
			Address:     "addressX",
		},
	}
	var publicKey = "public_key"

	// call can sign
	result := CanSign(unspentTransactions, transactionInput, publicKey)

	// result should false
	if result {
		t.Errorf("when no unspent transaction matches the result should be false")
	}
}

// Test_CanSign_UnspentTransactionWrongAddress : Tests if unspent transaction address does not equal public key.
func Test_CanSign_UnspentTransactionWrongAddress(t *testing.T) {

	// prepare input
	var transactionInput *TransactionInput = &TransactionInput{
		OutputID:    "outID",
		OutputIndex: 10,
	}
	var unspentTransactions = []*UnspentTransactionOutput{
		{
			OutputID:    "outID",
			OutputIndex: 10,
			Address:     "addressX",
		},
	}
	var publicKey = "public_key"

	// call can sign
	result := CanSign(unspentTransactions, transactionInput, publicKey)

	// result should false
	if result {
		t.Errorf("when unspent transaction address is not the same as public key the result should be false")
	}
}

// Test_CanSign_Correct : Tests if unspent transaction is found and address is correct.
func Test_CanSign_Correct(t *testing.T) {

	// prepare input
	var transactionInput *TransactionInput = &TransactionInput{
		OutputID:    "outID",
		OutputIndex: 10,
	}
	var unspentTransactions = []*UnspentTransactionOutput{
		{
			OutputID:    "outID-1",
			OutputIndex: 1000,
			Address:     "addressX",
		},
		{
			OutputID:    "outID",
			OutputIndex: 10,
			Address:     "public_key",
		},
	}
	var publicKey = "public_key"

	// call can sign
	result := CanSign(unspentTransactions, transactionInput, publicKey)

	// result should false
	if !result {
		t.Errorf("The result must be true when there is unspent transaction and adddress is the same the public key")
	}
}
