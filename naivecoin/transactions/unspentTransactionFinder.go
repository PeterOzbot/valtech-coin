package transactions

//FindUnspentTransactionOutput : Locates unspent transaction output that matches to the transaction input.
func FindUnspentTransactionOutput(unspentTransactionOutputs []*UnspentTransactionOutput, transactionInput *TransactionInput) *UnspentTransactionOutput {
	if transactionInput == nil {
		return nil
	}

	// find the unspent transaction that matched transaction input
	var matchingUnspentTransaction *UnspentTransactionOutput
	for _, unspentTransaction := range unspentTransactionOutputs {

		// match to transaction input
		if transactionInput.IsMatch(unspentTransaction) {
			matchingUnspentTransaction = unspentTransaction
			break
		}
	}

	// return what was found
	return matchingUnspentTransaction
}

//IsMatch : Checks if unspent transaction output matches to the transaction input.
func (transactionInput *TransactionInput) IsMatch(unspentTransactionOutput *UnspentTransactionOutput) bool {
	if transactionInput == nil || unspentTransactionOutput == nil {
		return false
	}

	return unspentTransactionOutput.OutputID == transactionInput.OutputID && unspentTransactionOutput.OutputIndex == transactionInput.OutputIndex
}
