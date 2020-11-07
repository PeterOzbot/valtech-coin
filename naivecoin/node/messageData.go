package node

//MessageData : Holds data for comunication between nodes.
type MessageData struct {
	Type string `json:"type" binding:"required"`
	Data string `json:"data" binding:"required"`
}

// socket data types
const (
	//QueryLatestBlockType : node requests latest block
	QueryLatestBlockType = "QueryLatestBlockType"
	//ResponseLatestBlockType : node sends latest block
	ResponseLatestBlockType = "ResponseLatestBlockType"
	//QueryBlockchainType :  node requests full block chain
	QueryBlockchainType = "QueryBlockchainType"
	//ResponseBlockchainType: node sends full block chain
	ResponseBlockchainType = "ResponseBlockchainType"
)
