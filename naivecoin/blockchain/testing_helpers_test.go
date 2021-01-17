package blockchain

import (
	"naivecoin/transactions"
	"time"
)

// creates chain state for testing purposes
func getTestChainState(unspentTransactionOutputs []*transactions.UnspentTransactionOutput) *ChainState {
	return &ChainState{
		TransactionValidator: &testBlockTransactionValidator{},
		ChainValidator: &testChainValidator{
			TransactionValidator: &testBlockTransactionValidator{},
			UnspentTransactionManager: &testUnspentTransactionManager{
				UnspentTransactionOutputs: unspentTransactionOutputs,
			},
			UnspentTransactionOutputs: unspentTransactionOutputs,
		},
	}
}

// mock validator struct
type testChainValidator struct {
	TransactionValidator      transactions.IBlockTransactionValidator
	UnspentTransactionManager transactions.IUnspentTransactionManager
	UnspentTransactionOutputs []*transactions.UnspentTransactionOutput
}

// mock validator methods
func (test *testChainValidator) IsValidChain(blockchain []*Block, currentTimestamp time.Time) (bool, []*transactions.UnspentTransactionOutput, error) {
	return true, test.UnspentTransactionOutputs, nil
}

func (test *testChainValidator) UpdateUnspentTransactionOutputs(newTransactions []*transactions.Transaction, unspentTransactionOutputs []*transactions.UnspentTransactionOutput) []*transactions.UnspentTransactionOutput {
	return test.UnspentTransactionManager.UpdateUnspentTransactionOutputs(newTransactions, unspentTransactionOutputs)
}

func (test *testChainValidator) ValidateBlockTransactions(transactions []*transactions.Transaction, unspentTransactionOutputs []*transactions.UnspentTransactionOutput, blockIndex int) (bool, error) {
	return test.TransactionValidator.ValidateBlockTransactions(transactions, unspentTransactionOutputs, blockIndex)
}

// mocked transaction validator
type testBlockTransactionValidator struct {
	Result bool
}

// mocked transaction validation which return success
func (r testBlockTransactionValidator) ValidateBlockTransactions(transactions []*transactions.Transaction, unspentTransactionOutputs []*transactions.UnspentTransactionOutput, blockIndex int) (bool, error) {
	return r.Result, nil
}

// mocked transaction manager
type testUnspentTransactionManager struct {
	UnspentTransactionOutputs []*transactions.UnspentTransactionOutput
}

// mocked update unspent transaction output that return whatever mocked manager holds
func (r testUnspentTransactionManager) UpdateUnspentTransactionOutputs(newTransactions []*transactions.Transaction, unspentTransactionOutputs []*transactions.UnspentTransactionOutput) []*transactions.UnspentTransactionOutput {
	return r.UnspentTransactionOutputs
}

// creates test validator
func createValidator(result bool, unspentTransactionOutputs []*transactions.UnspentTransactionOutput) *ChainValidator {
	return &ChainValidator{
		TransactionValidator: &testBlockTransactionValidator{
			Result: result,
		},
		UnspentTransactionManager: &testUnspentTransactionManager{
			UnspentTransactionOutputs: unspentTransactionOutputs,
		},
	}
}

// creates chain for testing purposes
func prepareChain(addSecond bool) []*Block {
	// create genesis block
	genesisBlock := GenesisBlock()

	// create chain and add genesis block
	var chain []*Block
	chain = append(chain, genesisBlock)

	// create second block and add
	if addSecond {
		secondBlock := &Block{
			Index:        1,
			Message:      "My second block. HaHaHa",
			PreviousHash: chain[0].Hash,
			Hash:         "37eca9944ecb7da3c1cd07786bdfe089593cedd118c5a57911e0354543ae5d10",
			Timestamp:    time.Unix(15884366862500, 0),
		}
		chain = append(chain, secondBlock)
	}

	// return chain
	return chain
}
