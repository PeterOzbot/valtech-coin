package blockchain

import "time"

//Block : Defines single block in the blockchain
type Block struct {
	Index        int
	Hash         string
	PreviousHash string
	Timestamp    time.Time
	Data         string
	Difficulty   int
	Nonce        int
}
