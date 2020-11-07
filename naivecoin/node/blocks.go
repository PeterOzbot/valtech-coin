package node

import (
	"fmt"
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
	newBlockMessage, err := OnMinedBlock(newBlock)
	if err == nil {
		p2p.NotifyPeers(newBlockMessage, c)
	} else {
		fmt.Println("serializing message to notify peers failed : ", err)
	}

	// respond with new block
	c.JSON(http.StatusOK, newBlock)
}

//SelectChain :  Used to select new chain when received from other node.
func SelectChain(newBlockchain []*blockchain.Block) {
	currentBlockchain := blockchain.GetBlockchain()
	blockchain.SetBlockchain(blockchain.SelectChain(newBlockchain, currentBlockchain))
}

//ReceivedBlock : When other peers send new block this method process it.
func ReceivedBlock(newBlock *blockchain.Block) bool {
	// get latest block
	latestBlock := blockchain.GetLatestBlock()

	// check if new block is relevant
	if newBlock.Index > latestBlock.Index {

		// check if new block is next for this node blockchain
		if newBlock.Hash == latestBlock.PreviousHash {
			return blockchain.AddBlockToChain(newBlock)
		}
	}

	// new block was ignored
	return false
}
