package transactions

//ValidateBlockTransactions : Validates transactions that belong to a single block.
func ValidateBlockTransactions(transactions []*Transaction, unspentTransactionOutputs []*UnspentTransactionOutput, blockIndex int) (bool, error) {

	// if there are no transactions its not valid, there must be at least one
	if len(transactions) < 1 {
		return false, nil
	}

	// get coinbase transaction and validate
	var coinbaseTransaction = transactions[0]
	if !coinbaseTransaction.IsValidCoinbase(blockIndex) {
		return false, nil
	}

	// get the rest of transactions and validate each
	var normalTransactions = transactions[1:]
	for _, transaction := range normalTransactions {

		// validate transaction and break if invalid is found
		isValid, err := transaction.IsValid(unspentTransactionOutputs)
		if !isValid || err != nil {
			return false, err
		}
	}

	// transactions are valid
	return true, nil
}
