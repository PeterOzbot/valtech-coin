package node

import (
	"naivecoin/blockchain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//GetBlockChain : Loads current block chain and serializes it to the context which should be returned to the client.
func GetBlockChain(c *gin.Context) {
	currentBlockChain := blockchain.GetBlockChain()
	c.JSON(200, currentBlockChain)
}

//Mineblocks : Generates next block.
func Mineblocks(c *gin.Context) {
	// deserialize block data from body
	var blockData *BlockData = &BlockData{}
	if err := c.ShouldBindJSON(blockData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate
	if blockData == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Deserialization failed."})
		return
	}

	// get latest block
	var latestBlock = blockchain.GetBlockChain()[0]

	// create new block
	var newBlock = blockchain.GenerateBlock(blockData.Data, latestBlock, time.Now().Unix())

	// respond with new block
	c.JSON(http.StatusOK, newBlock)
}
