package blockchain

import "time"

//Block : Defines single block in the blockchain
type Block struct {
	Index        int
	Hash         string
	PreviousHash string
	Timestamp    time.Time
	Data         string
}

//GenesisBlock : Create first block of the blockchain - Genesis block.
func GenesisBlock() *Block {
	return &Block{
		Index:        0,
		Hash:         "8d46e9fd83ead04cb38e6150130d556b97adf40cc865842395c1400ce48f724b",
		Data:         "my genesis block!!",
		PreviousHash: "",
		Timestamp:    time.Unix(1465154705, 0),
	}
}
