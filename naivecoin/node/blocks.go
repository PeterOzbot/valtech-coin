package node

import (
	"fmt"
	"naivecoin/blockchain"
	"naivecoin/p2p"
	"naivecoin/transactions"
	"naivecoin/wallet"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//GetBlockchain : Loads current block chain and serializes it to the context which is returned to the client.
func GetBlockchain(c *gin.Context) {
	currentBlockchain := blockchain.GetBlockchain(CurrentBlockchain)
	c.JSON(200, currentBlockchain)
}

//GetUnspentTransactionOutputs : Returns current unspent transaction outputs.
func GetUnspentTransactionOutputs(c *gin.Context) {
	c.JSON(200, UnspentTransactionOutputs)
}

//MineBlock : Generates next block.
func MineBlock(c *gin.Context) {
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
	currentBlockchain := blockchain.GetBlockchain(CurrentBlockchain)

	// get current time
	var currentTimestamp = time.Now()

	// select chain, current or new one
	selectedChain, newChainWasSelected, newUnspentOutputs, err := blockchain.SelectChain(newBlockchain, currentBlockchain, currentTimestamp)

	// if there was no error set chain and notify peers
	if err == nil {

		// set new chain, selectedChain will be the chain that was selected
		CurrentBlockchain = selectedChain

		// if new chain was selected notify peers
		if newChainWasSelected {

			// update unspent transaction outputs
			UnspentTransactionOutputs = newUnspentOutputs

			// get latest block and notify peers
			latestBlock := CurrentBlockchain[len(currentBlockchain)-1]
			notifyPeers(latestBlock)
		}
	}

	return err
}

//ReceivedBlock : When other peers send new block this method process it.
// if the whole chain must be queried then the result is true
func ReceivedBlock(newBlock *blockchain.Block) (bool, error) {
	// get this node's latest block
	currentBlockchain := blockchain.GetBlockchain(CurrentBlockchain)
	latestBlock := currentBlockchain[len(currentBlockchain)-1]

	// get current timestamp
	var currentTimestamp = time.Now()

	// check if new block is relevant
	var err error
	if newBlock.Index > latestBlock.Index {

		// if new block is not added and its index is greater then this node's latest block, the whole chain may be stale
		var blockAddedToChain bool
		blockAddedToChain, CurrentBlockchain, err = addBlockToChain(latestBlock, newBlock, currentBlockchain, currentTimestamp)
		if blockAddedToChain && err != nil {

			// new block was added notify peers
			notifyPeers(newBlock)

			// return that the block was added
			return true, nil
		}

		// the new block was not added query the chain as current may be stale
		return true, nil
	}

	// new block was ignored or added, either way the chain is not needed
	return false, err
}

// Generates/mines new block and adds it to the internal blockchain
func generateNewBlock(blockData *BlockData) (*blockchain.Block, bool, error) {
	// get current block chain
	var currentBlockchain = blockchain.GetBlockchain(CurrentBlockchain)
	var latestBlock = currentBlockchain[len(currentBlockchain)-1]

	// determine current difficulty
	var difficulty = blockchain.GetDifficulty(latestBlock, currentBlockchain, BlockGenerationInterval, DifficultyAdjustmentInterval)

	// current time
	var currentTimestamp = time.Now()

	// generate coinbase transaction
	coinbaseTransaction := transactions.CoinbaseTransaction(Wallet.PublicKey, latestBlock.Index+1, transactions.CoinbaseAmount)

	// generate transaction from block data
	transaction, transactionErr := wallet.GenerateTransaction(*blockData.Address, Wallet, *blockData.Amount, UnspentTransactionOutputs)
	if transactionErr != nil {
		return nil, false, transactionErr
	}

	// create new block
	var newBlock = &blockchain.Block{
		Index:        latestBlock.Index + 1,
		PreviousHash: latestBlock.Hash,
		Timestamp:    currentTimestamp,
		Transactions: []*transactions.Transaction{coinbaseTransaction, transaction},
		Message:      blockData.Message,
		Difficulty:   difficulty,
	}

	// mine this block
	var err = newBlock.MineBlock()

	// if there is no error from mining block then add to the blockchain
	if err == nil {

		// add new block
		var blockAddedToChain bool
		blockAddedToChain, CurrentBlockchain, err = addBlockToChain(latestBlock, newBlock, currentBlockchain, currentTimestamp)

		// return generated and mined block
		return newBlock, blockAddedToChain, err
	}

	// return generated and mined block
	return newBlock, false, err
}

// Adds block to the blockchain.
func addBlockToChain(latestBlock *blockchain.Block, newBlock *blockchain.Block, currentBlockchain []*blockchain.Block, currentTimestamp time.Time) (bool, []*blockchain.Block, error) {

	// check if block is valid
	isValidNewBlock, err := blockchain.IsValidNewBlock(newBlock, latestBlock, currentTimestamp, UnspentTransactionOutputs)
	if err != nil {
		return false, currentBlockchain, err
	}

	// if it is valid add it
	if isValidNewBlock {

		// add to block chain
		newBlockChain := append(currentBlockchain, newBlock)

		// update current node unspent outputs
		UnspentTransactionOutputs = transactions.UpdateUnspentTransactionOutputs(newBlock.Transactions, UnspentTransactionOutputs)

		// return success
		return true, newBlockChain, nil
	}

	return false, currentBlockchain, nil
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
