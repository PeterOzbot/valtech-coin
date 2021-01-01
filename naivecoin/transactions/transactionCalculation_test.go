package transactions

import (
	"testing"
)

// Test_CalculateID_NilInput : Tests if nil input is handled without failing.
func Test_CalculateID_NilInput(t *testing.T) {

	// create nil transaction
	var transaction *Transaction

	// calculate ID
	result := transaction.CalculateID()

	// result should be empty string
	if len(result) != 0 {
		t.Errorf("result of nil transaction should be empty string.")
	}
}

// Test_CalculateID_Outputs : Tests if ID calculation includes transaction outputs.
func Test_CalculateID_Outputs(t *testing.T) {

	// create transaction
	var transaction = &Transaction{
		Outputs: []*TransactionOutput{
			{
				Address: "address1",
				Amount:  1.00000000,
			},
			{
				Address: "address2",
				Amount:  2.00000000,
			},
		},
	}

	// expected result
	expectedResult := "60fd3f3821177a548ea4b530c687d94860cf92b02430aff41882fcd04cd6a050"

	// calculate ID
	result := transaction.CalculateID()

	// check if result matches
	if result != expectedResult {
		t.Errorf("ID calculation failed. Actual: %s Expected: %s", result, expectedResult)
	}
}

// Test_CalculateID_Inputs : Tests if ID calculation includes transaction inputs.
func Test_CalculateID_Inputs(t *testing.T) {

	// create transaction
	var transaction = &Transaction{
		Inputs: []*TransactionInput{
			{
				OutputID:    "id1",
				OutputIndex: 1,
			},
			{
				OutputID:    "id2",
				OutputIndex: 2,
			},
		},
	}

	// expected result
	expectedResult := "c375c9c44cd9b753e83625c59a57947ac8b1cc9465f9cd90ed10255e40ec099d"

	// calculate ID
	result := transaction.CalculateID()

	// check if result matches
	if result != expectedResult {
		t.Errorf("ID calculation failed. Actual: %s Expected: %s", result, expectedResult)
	}
}

// Test_CalculateID_CorrectResult : Tests if ID calculation includes transaction inputs adn outputs.
func Test_CalculateID_CorrectResult(t *testing.T) {

	// create transaction
	var transaction = &Transaction{
		Inputs: []*TransactionInput{
			{
				OutputID:    "id1",
				OutputIndex: 1,
			},
		},
		Outputs: []*TransactionOutput{
			{
				Address: "address1",
				Amount:  1.00000000,
			},
		},
	}

	// expected result
	expectedResult := "7d89bc87859d24f03736ab81c618527b7e128d0e7f98754dec76bb231b5b533c"

	// calculate ID
	result := transaction.CalculateID()

	// check if result matches
	if result != expectedResult {
		t.Errorf("ID calculation failed. Actual: %s Expected: %s", result, expectedResult)
	}
}
