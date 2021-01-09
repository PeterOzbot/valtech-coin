package wallet

import (
	"naivecoin/transactions"
	"testing"
)

// Test_GetBalance_NoAddress : Tests if getting balance does not fail if there is no address.
func Test_GetBalance_NoAddress(t *testing.T) {
	// inputs
	var address *Address
	var unspentTransactionOutputs = []*transactions.UnspentTransactionOutput{
		{
			Address: "address-1",
			Amount:  1,
		},
		{
			Address: "address-2",
			Amount:  2,
		},
	}

	// get balance
	balance := GetBalance(address, unspentTransactionOutputs)

	// check result
	if balance != 0 {
		t.Errorf("when there is no address, balance should be zero")
	}
}

// Test_GetBalance_NoAddress :  Tests if getting balance returns correct amount if there are no unspent amounts.
func Test_GetBalance_NoUnspentOutputs(t *testing.T) {
	// inputs
	var address = &Address{
		PrivateKey: "",
		PublicKey:  "address-3",
	}
	var unspentTransactionOutputs = []*transactions.UnspentTransactionOutput{
		{
			Address: "address-1",
			Amount:  1,
		},
		{
			Address: "address-2",
			Amount:  2,
		},
	}

	// get balance
	balance := GetBalance(address, unspentTransactionOutputs)

	// check result
	if balance != 0 {
		t.Errorf("when there is no unspent outputs, balance should be zero")
	}
}

// Test_GetBalance_CorrectAmount : Tests if getting balance returns correct amount when there are some unspent outputs.
func Test_GetBalance_CorrectAmount(t *testing.T) {
	// inputs
	var address = &Address{
		PrivateKey: "",
		PublicKey:  "address-1",
	}
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
	}

	// get balance
	balance := GetBalance(address, unspentTransactionOutputs)

	// check result
	if balance != 4 {
		t.Errorf("sum amount of unspent outputs should be 4")
	}
}
