package transactions

//IUnspentTransactionManager : Defines interface for updating unspent transaction outputs
type IUnspentTransactionManager interface {
	// Processes new transactions and generates new unspent transactions.
	UpdateUnspentTransactionOutputs(newTransactions []*Transaction, unspentTransactionOutputs []*UnspentTransactionOutput) []*UnspentTransactionOutput
}

//UnspentTransactionManager : Default implementation of unspent transaction manager.
type UnspentTransactionManager struct{}
