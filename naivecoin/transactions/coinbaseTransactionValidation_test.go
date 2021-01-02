package transactions

import "testing"

// Test_IsValidCoinbase_NilInput : Tests if nil input is handled without failing.
func Test_IsValidCoinbase_NilInput(t *testing.T) {

	// create coinbase transaction
	var transaction *Transaction

	// calculate ID
	result := transaction.IsValidCoinbase(0)

	// result should be false
	if result {
		t.Errorf("result of nil transaction should be false")
	}
}

// Test_IsValidCoinbase_ID : Tests validation checks transaction ID.
func Test_IsValidCoinbase_ID(t *testing.T) {

	// create inputs
	var transaction = &Transaction{
		ID: "invalid-ID",
		Inputs: []*TransactionInput{
			{
				OutputID:    "OutputID",
				OutputIndex: 0,
				Signature:   "",
			},
		},
		Outputs: []*TransactionOutput{
			{
				Address: "address",
				Amount:  10,
			},
		},
	}
	var blockIndex = 0

	// validate
	result := transaction.IsValidCoinbase(blockIndex)

	// result should false
	if result {
		t.Errorf("transaction is invalid so result should be false")
	}
}

// Test_IsValidCoinbase_TransactionInputCount : Tests validation checks that there is only one transaction input.
func Test_IsValidCoinbase_TransactionInputCount(t *testing.T) {

	// create inputs
	var transaction = &Transaction{
		ID: "bcad4f2936bbc94f0632fdf586a54d1fbb0f719bf6d1c4d2bd3021b6b485fc43",
		Inputs: []*TransactionInput{
			{
				OutputID:    "OutputID-1",
				OutputIndex: 0,
				Signature:   "",
			},
			{
				OutputID:    "OutputID-2",
				OutputIndex: 1,
				Signature:   "",
			},
		},
		Outputs: []*TransactionOutput{},
	}
	var blockIndex = 0

	// validate
	result := transaction.IsValidCoinbase(blockIndex)

	// result should false
	if result {
		t.Errorf("transaction is invalid so result should be false")
	}
}

// Test_IsValidCoinbase_TransactionInput : Tests validation checks that the transaction input's output index is the same as blockindex.
func Test_IsValidCoinbase_TransactionInput(t *testing.T) {

	// create inputs
	var transaction = &Transaction{
		ID: "d3ba8a701c8982d18fa6d878a50cd7241234d10299797a098b4cf77bff668257",
		Inputs: []*TransactionInput{
			{
				OutputID:    "OutputID-1",
				OutputIndex: 5,
				Signature:   "",
			},
		},
		Outputs: []*TransactionOutput{},
	}
	var blockIndex = 23

	// validate
	result := transaction.IsValidCoinbase(blockIndex)

	// result should false
	if result {
		t.Errorf("transaction is invalid so result should be false")
	}
}

// Test_IsValidCoinbase_TransactionOutputCount : Tests validation checks that there is only one transaction output.
func Test_IsValidCoinbase_TransactionOutputCount(t *testing.T) {

	// create inputs
	var transaction = &Transaction{
		ID: "41c149542b19739a94dd4528918c4f1a9e6901a402ae94d14e2c3b2dc4bea866",
		Inputs: []*TransactionInput{
			{
				OutputID:    "OutputID-1",
				OutputIndex: 0,
				Signature:   "",
			},
		},
		Outputs: []*TransactionOutput{
			{
				Address: "address1",
				Amount:  10,
			},
			{
				Address: "address2",
				Amount:  20,
			},
		},
	}
	var blockIndex = 0

	// validate
	result := transaction.IsValidCoinbase(blockIndex)

	// result should false
	if result {
		t.Errorf("transaction is invalid so result should be false")
	}
}

// Test_IsValidCoinbase_TransactionOutput : Tests validation checks that the transaction output's amount is the same as coinbase amount.
func Test_IsValidCoinbase_TransactionOutput(t *testing.T) {

	// create inputs
	var transaction = &Transaction{
		ID: "c8f4ea80717166b96b772d5f8b4d99075b6a2f4fd3d0fb5a4be49c269cddb3fd",
		Inputs: []*TransactionInput{
			{
				OutputID:    "OutputID-1",
				OutputIndex: 23,
				Signature:   "",
			},
		},
		Outputs: []*TransactionOutput{
			{
				Address: "address1",
				Amount:  10,
			},
		},
	}
	var blockIndex = 23

	// validate
	result := transaction.IsValidCoinbase(blockIndex)

	// result should false
	if result {
		t.Errorf("transaction is invalid so result should be false")
	}
}

// Test_IsValidCoinbase_IsValid : Tests validation checks that validation returns true when transaction is valid.
func Test_IsValidCoinbase_IsValid(t *testing.T) {

	// create inputs
	var transaction = &Transaction{
		ID: "c8ecb5a9032c0a888896f3f4d90c9ffef257c23c31bc87aec9597225e694c217",
		Inputs: []*TransactionInput{
			{
				OutputID:    "OutputID-1",
				OutputIndex: 23,
				Signature:   "",
			},
		},
		Outputs: []*TransactionOutput{
			{
				Address: "address1",
				Amount:  50,
			},
		},
	}
	var blockIndex = 23

	// validate
	result := transaction.IsValidCoinbase(blockIndex)

	// transaction is valid
	if !result {
		t.Errorf("transaction is valid so result should be true")
	}
}
