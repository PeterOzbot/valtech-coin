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

	// new block was added notify peers
	notifyPeers(newBlock)

	// respond with new block
	c.JSON(http.StatusOK, newBlock)
}

//SelectChain :  Used to select new chain when received from other node.
func SelectChain(newBlockchain []*blockchain.Block) {
	// get current block chain
	currentBlockchain := blockchain.GetBlockchain()

	// select chain, current or new one
	selectedChain, newChainWasSelected := blockchain.SelectChain(newBlockchain, currentBlockchain)

	// set new chain
	blockchain.SetBlockchain(selectedChain)

	// if new chain was selected notify peers
	if newChainWasSelected {
		// get latest block and notify peers
		latestBlock := blockchain.GetLatestBlock()
		notifyPeers(latestBlock)
	}
}

//ReceivedBlock : When other peers send new block this method process it.
// if the whole chain must be queried then the result is true
func ReceivedBlock(newBlock *blockchain.Block) bool {
	// get this node's latest block
	latestBlock := blockchain.GetLatestBlock()

	// check if new block is relevant
	if newBlock.Index > latestBlock.Index {

		// if new block is not added and its index is greater then this node's latest block, the whole chain may be stale
		if !blockchain.AddBlockToChain(newBlock) {
			return true
		}

		// new block was added notify peers
		notifyPeers(newBlock)
	}

	// new block was ignored or added, either way the chain is not needed
	return false
}

// notifies all peers about new block
func notifyPeers(newBlock *blockchain.Block) {
	// notify peers
	newBlockMessage, err := OnNewBlock(newBlock)
	if err == nil {
		p2p.NotifyPeers(newBlockMessage)
	} else {
		fmt.Println("serializing message to notify peers failed : ", err)
	}
}
