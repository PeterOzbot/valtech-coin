package blockchain

import (
	"naivecoin/transactions"
	"time"
)

//Block : Defines single block in the blockchain
type Block struct {
	Index        int
	Hash         string
	PreviousHash string
	Timestamp    time.Time
	Transactions []*transactions.Transaction
	Message      string
	Difficulty   int
	Nonce        int
}
