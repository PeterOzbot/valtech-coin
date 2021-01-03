package blockchain

import (
	"naivecoin/transactions"
	"time"
)

//GenesisBlock : Create first block of the blockchain - Genesis block.
func GenesisBlock() *Block {
	return &Block{
		Index:        0,
		Hash:         "7b9a528eda2b124b6555ab24f23a493805b7a0851e065bb3b0f749aa73e4c3e3",
		Transactions: []*transactions.Transaction{},
		Message:      "my genesis block!!",
		PreviousHash: "",
		Timestamp:    time.Unix(1465154705, 0),
		Nonce:        32,
		Difficulty:   0,
	}
}
