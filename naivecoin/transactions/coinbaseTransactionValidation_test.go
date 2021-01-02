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
		ID: "invalid-id",
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

	// result should false
	if result {
		t.Errorf("transaction is invalid so result should be false")
	}
}

// Test_IsValidCoinbase_TransactionInputCount : Tests validation checks that there is only one transaction input.
func Test_IsValidCoinbase_TransactionInputCount(t *testing.T) {

	// create inputs
	var transaction = &Transaction{
		ID: "f706c17df1e2ef4e8ac4e2e27f6da52d009fa65b67b6f19a41a1de0cab8b7390",
		Inputs: []*TransactionInput{
			{
				OutputID:    "OutputID-1",
				OutputIndex: 23,
				Signature:   "",
			},
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

	// result should false
	if result {
		t.Errorf("transaction is invalid so result should be false")
	}
}

// Test_IsValidCoinbase_TransactionInput : Tests validation checks that the transaction input's output index is the same as blockindex.
func Test_IsValidCoinbase_TransactionInput(t *testing.T) {

	// create inputs
	var transaction = &Transaction{
		ID: "37925c37799dce925533cdc785d08e4eaded9c0ee9df52dfcdb388c4b908cc2a",
		Inputs: []*TransactionInput{
			{
				OutputID:    "OutputID-1",
				OutputIndex: 4,
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

	// result should false
	if result {
		t.Errorf("transaction is invalid so result should be false")
	}
}

// Test_IsValidCoinbase_TransactionOutputCount : Tests validation checks that there is only one transaction output.
func Test_IsValidCoinbase_TransactionOutputCount(t *testing.T) {

	// create inputs
	var transaction = &Transaction{
		ID: "322ba82b33d3eaef2be54f97f8696c299ca1494fe2c797523660a18e786cdb4c",
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
			{
				Address: "address1",
				Amount:  50,
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
