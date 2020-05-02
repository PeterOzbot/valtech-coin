package blockchain

import "time"

//Block : Defines single block in the blockchain
type Block struct {
	index        int
	hash         string
	previousHash string
	timestamp    time.Time
	data         string
}

//GenesisBlock : Create first block of the blockchain - Genesis block.
func GenesisBlock() *Block {
	return &Block{
		index:        0,
		hash:         "8d46e9fd83ead04cb38e6150130d556b97adf40cc865842395c1400ce48f724b",
		data:         "my genesis block!!",
		previousHash: "",
		timestamp:    time.Unix(1465154705, 0),
	}
}
