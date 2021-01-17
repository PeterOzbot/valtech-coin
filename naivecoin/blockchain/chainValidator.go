package blockchain

import (
	"naivecoin/transactions"
	"time"
)

//IChainValidator : Defines interface for validating blockchain.
type IChainValidator interface {
	// Validates transactions that belong to a single block.
	ValidateBlockTransactions(transactions []*transactions.Transaction, unspentTransactionOutputs []*transactions.UnspentTransactionOutput, blockIndex int) (bool, error)

	//  Processes new transactions and generates new unspent transactions.
	UpdateUnspentTransactionOutputs(newTransactions []*transactions.Transaction, unspentTransactionOutputs []*transactions.UnspentTransactionOutput) []*transactions.UnspentTransactionOutput

	// Validates the whole block chain, including the first genesis block and all the rest.
	IsValidChain(blockchain []*Block, currentTimestamp time.Time) (bool, []*transactions.UnspentTransactionOutput, error)
}

//ChainValidator : Defines chain validator struct that has the dependencies for chain validation.
type ChainValidator struct {
	TransactionValidator      transactions.IBlockTransactionValidator
	UnspentTransactionManager transactions.IUnspentTransactionManager
}

//UpdateUnspentTransactionOutputs : Processes new transactions and generates new unspent transactions.
func (chainValidator *ChainValidator) UpdateUnspentTransactionOutputs(newTransactions []*transactions.Transaction, unspentTransactionOutputs []*transactions.UnspentTransactionOutput) []*transactions.UnspentTransactionOutput {
	return chainValidator.UnspentTransactionManager.UpdateUnspentTransactionOutputs(newTransactions, unspentTransactionOutputs)
}

//ValidateBlockTransactions : Validates transactions that belong to a single block.
func (chainValidator *ChainValidator) ValidateBlockTransactions(transactions []*transactions.Transaction, unspentTransactionOutputs []*transactions.UnspentTransactionOutput, blockIndex int) (bool, error) {
	return chainValidator.TransactionValidator.ValidateBlockTransactions(transactions, unspentTransactionOutputs, blockIndex)
}
