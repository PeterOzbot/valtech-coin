package node

import (
	"naivecoin/blockchain"
	"naivecoin/transactions"
	"naivecoin/wallet"
)

const (
	//BlockGenerationInterval : Defines how often a block should be found in seconds
	BlockGenerationInterval int = 100
	//DifficultyAdjustmentInterval : Defines how often the difficulty should adjust to the increasing or decreasing network hashrate
	DifficultyAdjustmentInterval int = 0
)

//CurrentBlockchain : Current block chain.
var CurrentBlockchain []*blockchain.Block

//UnspentTransactionOutputs : Collection of current unspent transaction outputs
var UnspentTransactionOutputs []*transactions.UnspentTransactionOutput = []*transactions.UnspentTransactionOutput{}

//Wallet : This node wallet.
var Wallet *wallet.Address
