package transactions

//IBlockTransactionValidator : Defines interface for transaction validation
type IBlockTransactionValidator interface {
	ValidateBlockTransactions(transactions []*Transaction, unspentTransactionOutputs []*UnspentTransactionOutput, blockIndex int) (bool, error)
}

//BlockTransactionValidator : Default transaction validator.
type BlockTransactionValidator struct{}
