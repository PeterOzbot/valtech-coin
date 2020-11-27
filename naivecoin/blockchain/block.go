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

//GenesisBlock : Create first block of the blockchain - Genesis block.
func GenesisBlock() *Block {
	return &Block{
		Index:        0,
		Hash:         "07cd44e61db86462cd7727374f59a5c6cbe02c896746b44322e195d1f88b10f2",
		Data:         "my genesis block!!",
		PreviousHash: "",
		Timestamp:    time.Unix(1465154705, 0),
		Nonce:        32,
	}
}
