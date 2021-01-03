package blockchain

import (
	"naivecoin/transactions"
	"testing"
	"time"
)

// Test_MineBlock_MineBlock : Tests if hash is correctly mined.
func Test_MineBlock_MineBlock(t *testing.T) {
	var block = &Block{
		Index: 0,
		Hash:  "",
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
		Message:      "my genesis block!!",
		PreviousHash: "",
		Timestamp:    time.Unix(1465154705, 0),
		Difficulty:   5,
		Nonce:        0,
	}

	// hash block
	block.MineBlock()

	// result should be correct hash
	expectedHash := "01cae462faef5b5132df4a29cba801c620813bc7033ad8ae7d4ad8a8806bb7ca"
	if block.Hash != expectedHash {
		t.Errorf("mining result is incorrect, Actual: %s Expected: %s", block.Hash, expectedHash)
	}
	expectedNonce := 4
	if block.Nonce != expectedNonce {
		t.Errorf("mining result is incorrect, Actual: %d Expected: %d", block.Nonce, expectedNonce)
	}
}
