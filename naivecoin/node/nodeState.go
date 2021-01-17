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

//UnspentTransactionManager : Logic for managing unspent transaction outputs.
var UnspentTransactionManager = &transactions.UnspentTransactionManager{}

//TransactionValidator : Logic for validating block transactions
var TransactionValidator = &transactions.BlockTransactionValidator{}

//ChainState : Logic for chain state handling.
var ChainState = &blockchain.ChainState{TransactionValidator: TransactionValidator, ChainValidator: &blockchain.ChainValidator{TransactionValidator: TransactionValidator, UnspentTransactionManager: UnspentTransactionManager}}
