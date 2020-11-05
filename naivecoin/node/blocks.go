package node

import (
	"naivecoin/blockchain"
	"naivecoin/p2p"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//GetBlockchain : Loads current block chain and serializes it to the context which is returned to the client.
func GetBlockchain(c *gin.Context) {
	currentBlockchain := blockchain.GetBlockchain()
	c.JSON(200, currentBlockchain)
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
	var latestBlock = blockchain.GetLatestBlock()

	// create new block
	var newBlock = blockchain.GenerateBlock(blockData.Data, latestBlock, time.Now().Unix())

	// add new block
	blockchain.AddBlockToChain(newBlock)

	// notify peers
	newBlockMessage := "" //TODO: implement
	p2p.NotifyPeers(newBlockMessage, c)

	// respond with new block
	c.JSON(http.StatusOK, newBlock)
}

//SelectChain :  Used to select new chain when received from other node.
func SelectChain(newBlockchain []*blockchain.Block) {
	currentBlockchain := blockchain.GetBlockchain()
	blockchain.SetBlockchain(blockchain.SelectChain(newBlockchain, currentBlockchain))
}

//AddBlock : When other peers send new block this process it.
func AddBlock(newBlock *blockchain.Block) {
	blockchain.AddBlockToChain(newBlock)
}
