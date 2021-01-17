package blockchain

import (
	"naivecoin/transactions"
)

//ChainState :  Defines chain state struct that has the dependencies for chain state handling.
type ChainState struct {
	TransactionValidator transactions.IBlockTransactionValidator
	ChainValidator       IChainValidator
}
