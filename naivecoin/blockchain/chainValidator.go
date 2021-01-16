package blockchain

import (
	"naivecoin/transactions"
	"time"
)

//IsValidChain : Validates the whole block chain, including the first genesis block and all the rest.
func IsValidChain(blockchain []*Block, currentTimestamp time.Time) (bool, []*transactions.UnspentTransactionOutput, error) {
	if blockchain == nil || len(blockchain) == 0 {
		return false, []*transactions.UnspentTransactionOutput{}, nil
	}

	// lets validate genesis block first
	if !isGenesisBlockValid(blockchain[0]) {
		return false, []*transactions.UnspentTransactionOutput{}, nil
	}

	// define unspent transaction outputs which get updated and used in validation
	var unspentTransactionOutputs []*transactions.UnspentTransactionOutput

	// loop through chain and validate all blocks with each other
	for index := 1; index < len(blockchain); index++ {

		currentBlock := blockchain[index]
		previousBlock := blockchain[index-1]

		isBlockValid, err := IsValidNewBlock(currentBlock, previousBlock, currentTimestamp, unspentTransactionOutputs)

		if !isBlockValid || err != nil {
			return false, []*transactions.UnspentTransactionOutput{}, err
		}

		// update current node unspent outputs
		unspentTransactionOutputs = transactions.UpdateUnspentTransactionOutputs(currentBlock.Transactions, unspentTransactionOutputs)
	}

	return true, unspentTransactionOutputs, nil
}
