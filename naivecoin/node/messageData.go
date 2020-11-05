package node

//MessageData : Holds data for comunication between nodes.
type MessageData struct {
	Type string `json:"type" binding:"required"`
	Data string `json:"data" binding:"required"`
}

// socket data types
const (
	//GetLatestBlockType : node requests latest block
	GetLatestBlockType = "GetLatestBlock"
	//LatestBlockType : node receives latest block
	LatestBlockType = "LatestBlock"
	//GetBlockchainType :  node requests full block chain
	GetBlockchainType = "GetBlockchain"
	//BlockchainType: node receives full block chain
	BlockchainType = "Blockchain"
)
