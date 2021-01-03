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
	currentBlockchain := blockchain.GetBlockchain(blockchain.CurrentBlockchain)
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
	newBlock, blockAddedToChain, err := generateNewBlock(blockData)

	// notify result
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error: Generatin new block failed. ": err.Error()})
		return
	}

	// new block was added notify peers
	if blockAddedToChain {
		notifyPeers(newBlock)
	}

	// respond with new block
	c.JSON(http.StatusOK, newBlock)
}

//SelectChain :  Used to select new chain when received from other node.
func SelectChain(newBlockchain []*blockchain.Block) error {
	// get current block chain
	currentBlockchain := blockchain.GetBlockchain(blockchain.CurrentBlockchain)

	// get current time
	var currentTimestamp = time.Now()

	// select chain, current or new one
	selectedChain, newChainWasSelected, err := blockchain.SelectChain(newBlockchain, currentBlockchain, currentTimestamp)

	// if there was no error set chain and notify peers
	if err == nil {

		// set new chain, selectedChain will be the chain that was selected
		blockchain.CurrentBlockchain = selectedChain

		// if new chain was selected notify peers
		if newChainWasSelected {
			// get latest block and notify peers
			latestBlock := blockchain.CurrentBlockchain[len(currentBlockchain)-1]
			notifyPeers(latestBlock)
		}
	}

	return err
}

//ReceivedBlock : When other peers send new block this method process it.
// if the whole chain must be queried then the result is true
func ReceivedBlock(newBlock *blockchain.Block) (bool, error) {
	// get this node's latest block
	currentBlockchain := blockchain.GetBlockchain(blockchain.CurrentBlockchain)
	latestBlock := currentBlockchain[len(currentBlockchain)-1]

	// get current timestamp
	var currentTimestamp = time.Now()

	// check if new block is relevant
	var err error
	if newBlock.Index > latestBlock.Index {

		// if new block is not added and its index is greater then this node's latest block, the whole chain may be stale
		var blockAddedToChain bool
		blockAddedToChain, blockchain.CurrentBlockchain, err = blockchain.AddBlockToChain(latestBlock, newBlock, currentBlockchain, currentTimestamp)
		if blockAddedToChain && err != nil {

			// new block was added notify peers
			notifyPeers(newBlock)

			// return that the block was added
			return true, nil
		}
	}

	// new block was ignored or added, either way the chain is not needed
	return false, err
}

//generates/mines new block and adds it to the internal blockchain
func generateNewBlock(blockData *BlockData) (*blockchain.Block, bool, error) {
	// get current block chain
	var currentBlockchain = blockchain.GetBlockchain(blockchain.CurrentBlockchain)
	var latestBlock = currentBlockchain[len(currentBlockchain)-1]

	// determine current difficulty
	var difficulty = blockchain.GetDifficulty(latestBlock, currentBlockchain, blockchain.BlockGenerationInterval, blockchain.DifficultyAdjustmentInterval)

	// current time
	var currentTimestamp = time.Now()

	// create new block
	var newBlock = &blockchain.Block{
		Index:        latestBlock.Index + 1,
		PreviousHash: latestBlock.Hash,
		Timestamp:    currentTimestamp,
		Transactions: blockData.Transactions,
		Message:      blockData.Message,
		Difficulty:   difficulty,
	}

	// mine this block
	var err = newBlock.MineBlock()

	// if there is no error from mining block then add to the blockchain
	if err == nil {

		// add new block
		var blockAddedToChain bool
		blockAddedToChain, blockchain.CurrentBlockchain, err = blockchain.AddBlockToChain(latestBlock, newBlock, currentBlockchain, currentTimestamp)

		// return generated and mined block
		return newBlock, blockAddedToChain, err
	}

	// return generated and mined block
	return newBlock, false, err
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
