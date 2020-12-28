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

	// generate and mine block
	newBlock := generateNewBlock(blockData)

	// new block was added notify peers
	notifyPeers(newBlock)

	// respond with new block
	c.JSON(http.StatusOK, newBlock)
}

//SelectChain :  Used to select new chain when received from other node.
func SelectChain(newBlockchain []*blockchain.Block) {
	// get current block chain
	currentBlockchain := blockchain.GetBlockchain()

	// get current time
	var currentTimestamp = time.Now()

	// select chain, current or new one
	selectedChain, newChainWasSelected := blockchain.SelectChain(newBlockchain, currentBlockchain, currentTimestamp)

	// set new chain
	blockchain.SetBlockchain(selectedChain)

	// if new chain was selected notify peers
	if newChainWasSelected {
		// get latest block and notify peers
		latestBlock := currentBlockchain[len(currentBlockchain)-1]
		notifyPeers(latestBlock)
	}
}

//ReceivedBlock : When other peers send new block this method process it.
// if the whole chain must be queried then the result is true
func ReceivedBlock(newBlock *blockchain.Block) bool {
	// get this node's latest block
	currentBlockchain := blockchain.GetBlockchain()
	latestBlock := currentBlockchain[len(currentBlockchain)-1]

	// get current timestamp
	var currentTimestamp = time.Now()

	// check if new block is relevant
	if newBlock.Index > latestBlock.Index {

		// if new block is not added and its index is greater then this node's latest block, the whole chain may be stale
		if !blockchain.AddBlockToChain(latestBlock, newBlock, currentBlockchain, currentTimestamp) {
			return true
		}

		// new block was added notify peers
		notifyPeers(newBlock)
	}

	// new block was ignored or added, either way the chain is not needed
	return false
}

//generates/mines new block and adds it to the internal blockchain
func generateNewBlock(blockData *BlockData) *blockchain.Block {
	// get current block chain
	var currentBlockchain = blockchain.GetBlockchain()
	var latestBlock = currentBlockchain[len(currentBlockchain)-1]

	// determine current difficulty
	var difficulty = blockchain.GetDifficulty(latestBlock, currentBlockchain, blockchain.BlockGenerationInterval, blockchain.DifficultyAdjustmentInterval)

	// current time
	var currentTimestamp = time.Now()

	// create new block
	var newBlock = blockchain.GenerateBlock(blockData.Data, latestBlock, currentTimestamp, difficulty)

	// mine this block
	newBlock.MineBlock()

	// add new block
	blockchain.AddBlockToChain(latestBlock, newBlock, currentBlockchain, currentTimestamp)

	// return generated and mined block
	return newBlock
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
