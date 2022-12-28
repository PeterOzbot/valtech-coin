package mempool

import "naivecoin/transactions"

//CanAddTransaction : Validates transaction if it can be added to the transaction pool.
func CanAddTransaction(transaction *transactions.Transaction) bool {
	return false
}
