package blockchain

import (
	"naivecoin/transactions"
	"testing"
	"time"
)

// Test_CalculateHash_NilInpout : Tests if hashing handles nil input without failing.
func Test_CalculateHash_NilInput(t *testing.T) {

	// create nil block
	var block *Block

	// hash block
	result, _ := block.CalculateHash()

	// result should be empty string
	if len(result) != 0 {
		t.Errorf("hashing result of nil block should be empty string.")
	}
}

// Test_CalculateHash_CorrectHash : Tests if hashing creates correct hash value.
func Test_CalculateHash_CorrectHash(t *testing.T) {

	// create example block
	var block *Block = &Block{
		Index: 15,
		Transactions: []*transactions.Transaction{
			{
				ID: "transactionID",
				Inputs: []*transactions.TransactionInput{
					{
						OutputID:    "output-ID",
						OutputIndex: 1,
						Signature:   "signature",
					},
				},
				Outputs: []*transactions.TransactionOutput{
					{
						Address: "address",
						Amount:  25,
					},
				},
			},
		},
		Message:      "Dober dan gospod kamplan.",
		PreviousHash: "ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae",
		Timestamp:    time.Unix(1588430083866862500, 0),
		Difficulty:   5,
		Nonce:        32,
	}

	//  is manually generated hash from block
	expected := "50d35d6ece26625cd28aa61cb8535ccd25097b7d1485f834e168045aca831309"

	// hash block
	result, _ := block.CalculateHash()

	// result should be correct hash
	if result != expected {
		t.Errorf("hashing result is incorrect, Actual: %s Expected: %s", result, expected)
	}
}

// Test_CalculateHash_NoTransactions : Tests if hashing handles no transactions without failing.
func Test_CalculateHash_NoTransactions(t *testing.T) {

	// create example block
	var block *Block = &Block{
		Index:        15,
		Transactions: []*transactions.Transaction{},
		Message:      "Dober dan gospod kamplan.",
		PreviousHash: "ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae",
		Timestamp:    time.Unix(1588430083866862500, 0),
		Difficulty:   5,
		Nonce:        32,
	}

	//  is manually generated hash from block
	expected := "73a20e74f447c0b2a1918284000605b645a911e667500d8eb42f1f874d0d7974"

	// hash block
	result, _ := block.CalculateHash()

	// result should be correct hash
	if result != expected {
		t.Errorf("hashing result is incorrect, Actual: %s Expected: %s", result, expected)
	}
}
