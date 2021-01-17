package transactions

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Test_Sign_ : Tests if new transaction outputs are updated/added from input transactions.
func Test_UpdateUnspentTransactionOutputs_NewTransactionOutputs(t *testing.T) {
	var unspentTransactionManager = &UnspentTransactionManager{}
	// create inputs
	var newTransactions = []*Transaction{
		{
			ID: "transaction_ID-1",
			Outputs: []*TransactionOutput{
				{
					Address: "address-1",
					Amount:  1,
				},
				{
					Address: "address-2",
					Amount:  2,
				},
			},
		},
		{
			ID: "transaction_ID-2",
			Outputs: []*TransactionOutput{
				{
					Address: "address-3",
					Amount:  3,
				},
			},
		},
	}
	var unspentTransactionOutputs []*UnspentTransactionOutput

	// expected result
	expectedResult := []*UnspentTransactionOutput{
		{
			OutputID:    "transaction_ID-1",
			OutputIndex: 0,
			Address:     "address-1",
			Amount:      1,
		},
		{
			OutputID:    "transaction_ID-1",
			OutputIndex: 1,
			Address:     "address-2",
			Amount:      2,
		},
		{
			OutputID:    "transaction_ID-2",
			OutputIndex: 0,
			Address:     "address-3",
			Amount:      3,
		},
	}

	// update unspent transactions
	result := unspentTransactionManager.UpdateUnspentTransactionOutputs(newTransactions, unspentTransactionOutputs)

	// validate
	if len(result) != len(expectedResult) {
		t.Errorf("new unspent transactions should be %d not %d", len(expectedResult), len(result))
		return
	}
	// validate each
	for index, resultUnspent := range result {
		// find expected
		expectedResultUnspent := expectedResult[index]
		// validate expected
		if !cmp.Equal(resultUnspent, expectedResultUnspent) {
			t.Errorf("unspent transaction output does not match")
		}
	}
}

// Test_UpdateUnspentTransactionOutputs_FilteredTransactionOutputs : Tests if existing unspent transaction outputs are filtered with new transaction inputs.
func Test_UpdateUnspentTransactionOutputs_FilteredTransactionOutputs(t *testing.T) {
	var unspentTransactionManager = &UnspentTransactionManager{}
	// create inputs
	var newTransactions = []*Transaction{
		{
			Inputs: []*TransactionInput{
				{
					OutputID:    "transaction-out-id-1",
					OutputIndex: 0,
				},
			},
		},
		{
			Inputs: []*TransactionInput{
				{
					OutputID:    "transaction-out-id-2",
					OutputIndex: 0,
				},
				{
					OutputID:    "transaction-out-id-3",
					OutputIndex: 1,
				},
			},
		},
	}
	var unspentTransactionOutputs = []*UnspentTransactionOutput{
		{
			OutputID:    "transaction-out-id-3",
			OutputIndex: 0,
			Address:     "address-1",
			Amount:      1,
		},
		{
			OutputID:    "transaction-out-id-3",
			OutputIndex: 1,
			Address:     "address-2",
			Amount:      2,
		},
		{
			OutputID:    "transaction-out-id-1",
			OutputIndex: 0,
			Address:     "address-3",
			Amount:      3,
		},
		{
			OutputID:    "transaction-out-id-2",
			OutputIndex: 1,
			Address:     "address-4",
			Amount:      4,
		},
	}

	// expected result
	expectedResult := []*UnspentTransactionOutput{
		{
			OutputID:    "transaction-out-id-3",
			OutputIndex: 0,
			Address:     "address-1",
			Amount:      1,
		},
		{
			OutputID:    "transaction-out-id-2",
			OutputIndex: 1,
			Address:     "address-4",
			Amount:      4,
		},
	}

	// update unspent transactions
	result := unspentTransactionManager.UpdateUnspentTransactionOutputs(newTransactions, unspentTransactionOutputs)

	// validate
	if len(result) != len(expectedResult) {
		t.Errorf("new unspent transactions should be %d not %d", len(expectedResult), len(result))
		return
	}
	// validate each
	for index, resultUnspent := range result {
		// find expected
		expectedResultUnspent := expectedResult[index]
		// validate expected
		if !cmp.Equal(resultUnspent, expectedResultUnspent) {
			t.Errorf("unspent transaction output does not match")
		}
	}
}

// Test_UpdateUnspentTransactionOutputs_CombinedTransactionOutputs : Tests if result contains both new and existing unspent transaction outputs.
func Test_UpdateUnspentTransactionOutputs_CombinedTransactionOutputs(t *testing.T) {
	var unspentTransactionManager = &UnspentTransactionManager{}
	// create inputs
	var newTransactions = []*Transaction{
		{
			ID: "transaction_ID-1",
			Outputs: []*TransactionOutput{
				{
					Address: "address-1",
					Amount:  1,
				},
				{
					Address: "address-2",
					Amount:  2,
				},
			},
			Inputs: []*TransactionInput{
				{
					OutputID:    "transaction-out-id-1",
					OutputIndex: 0,
				},
			},
		},
		{
			ID: "transaction_ID-2",
			Outputs: []*TransactionOutput{
				{
					Address: "address-3",
					Amount:  3,
				},
			},
			Inputs: []*TransactionInput{
				{
					OutputID:    "transaction-out-id-2",
					OutputIndex: 0,
				},
				{
					OutputID:    "transaction-out-id-3",
					OutputIndex: 1,
				},
			},
		},
	}
	var unspentTransactionOutputs = []*UnspentTransactionOutput{
		{
			OutputID:    "transaction-out-id-3",
			OutputIndex: 0,
			Address:     "address-1",
			Amount:      1,
		},
		{
			OutputID:    "transaction-out-id-3",
			OutputIndex: 1,
			Address:     "address-2",
			Amount:      2,
		},
		{
			OutputID:    "transaction-out-id-1",
			OutputIndex: 0,
			Address:     "address-3",
			Amount:      3,
		},
		{
			OutputID:    "transaction-out-id-2",
			OutputIndex: 1,
			Address:     "address-4",
			Amount:      4,
		},
	}

	// expected result
	expectedResult := []*UnspentTransactionOutput{
		{
			OutputID:    "transaction_ID-1",
			OutputIndex: 0,
			Address:     "address-1",
			Amount:      1,
		},
		{
			OutputID:    "transaction_ID-1",
			OutputIndex: 1,
			Address:     "address-2",
			Amount:      2,
		},
		{
			OutputID:    "transaction_ID-2",
			OutputIndex: 0,
			Address:     "address-3",
			Amount:      3,
		},
		{
			OutputID:    "transaction-out-id-3",
			OutputIndex: 0,
			Address:     "address-1",
			Amount:      1,
		},
		{
			OutputID:    "transaction-out-id-2",
			OutputIndex: 1,
			Address:     "address-4",
			Amount:      4,
		},
	}

	// update unspent transactions
	result := unspentTransactionManager.UpdateUnspentTransactionOutputs(newTransactions, unspentTransactionOutputs)

	// validate
	if len(result) != len(expectedResult) {
		t.Errorf("new unspent transactions should be %d not %d", len(expectedResult), len(result))
		return
	}
	// validate each
	for index, resultUnspent := range result {
		// find expected
		expectedResultUnspent := expectedResult[index]
		// validate expected
		if !cmp.Equal(resultUnspent, expectedResultUnspent) {
			t.Errorf("unspent transaction output does not match")
		}
	}
}

// Test_UpdateUnspentTransactionOutputsMultipleInputForTransactionOutputs : Tests if existing unspent transaction outputs are filtered when there are multiple inputs that match to single unspent output.
func Test_UpdateUnspentTransactionOutputsMultipleInputForTransactionOutputs(t *testing.T) {
	var unspentTransactionManager = &UnspentTransactionManager{}
	// create inputs
	var newTransactions = []*Transaction{
		{
			Inputs: []*TransactionInput{
				{
					OutputID:    "transaction-out-id-1",
					OutputIndex: 0,
				},
			},
		},
	}
	var unspentTransactionOutputs = []*UnspentTransactionOutput{
		{
			OutputID:    "transaction-out-id-1",
			OutputIndex: 0,
			Address:     "address-1",
			Amount:      1,
		},
		{
			OutputID:    "transaction-out-id-1",
			OutputIndex: 0,
			Address:     "address-2",
			Amount:      2,
		},
		{
			OutputID:    "transaction-out-id-2",
			OutputIndex: 1,
			Address:     "address-3",
			Amount:      3,
		},
	}

	// expected result
	expectedResult := []*UnspentTransactionOutput{
		{
			OutputID:    "transaction-out-id-2",
			OutputIndex: 1,
			Address:     "address-3",
			Amount:      3,
		},
	}

	// update unspent transactions
	result := unspentTransactionManager.UpdateUnspentTransactionOutputs(newTransactions, unspentTransactionOutputs)

	// validate
	if len(result) != len(expectedResult) {
		t.Errorf("new unspent transactions should be %d not %d", len(expectedResult), len(result))
		return
	}
	// validate each
	for index, resultUnspent := range result {
		// find expected
		expectedResultUnspent := expectedResult[index]
		// validate expected
		if !cmp.Equal(resultUnspent, expectedResultUnspent) {
			t.Errorf("unspent transaction output does not match")
		}
	}
}

// Test_UpdateUnspentTransactionOutputs_FilteredAllTransactionOutputs : Tests if all existing unspent transaction outputs are filtered correctly.
func Test_UpdateUnspentTransactionOutputs_FilteredAllTransactionOutputs(t *testing.T) {
	var unspentTransactionManager = &UnspentTransactionManager{}
	// create inputs
	var newTransactions = []*Transaction{
		{
			Inputs: []*TransactionInput{
				{
					OutputID:    "transaction-out-id-1",
					OutputIndex: 0,
				},
			},
		},
	}
	var unspentTransactionOutputs = []*UnspentTransactionOutput{
		{
			OutputID:    "transaction-out-id-1",
			OutputIndex: 0,
			Address:     "address-1",
			Amount:      1,
		},
		{
			OutputID:    "transaction-out-id-1",
			OutputIndex: 0,
			Address:     "address-1",
			Amount:      1,
		},
	}

	// expected result
	expectedResult := []*UnspentTransactionOutput{}

	// update unspent transactions
	result := unspentTransactionManager.UpdateUnspentTransactionOutputs(newTransactions, unspentTransactionOutputs)

	// validate
	if len(result) != len(expectedResult) {
		t.Errorf("new unspent transactions should be %d not %d", len(expectedResult), len(result))
		return
	}
	// validate each
	for index, resultUnspent := range result {
		// find expected
		expectedResultUnspent := expectedResult[index]
		// validate expected
		if !cmp.Equal(resultUnspent, expectedResultUnspent) {
			t.Errorf("unspent transaction output does not match")
		}
	}
}
