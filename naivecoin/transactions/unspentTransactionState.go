package transactions

// UpdateUnspentTransactionOutputs : Processes new transactions and generates new unspent transactions.
func UpdateUnspentTransactionOutputs(newTransactions []*Transaction, unspentTransactionOutputs []*UnspentTransactionOutput) []*UnspentTransactionOutput {

	// define array of unspent transactions
	var newUnspentTransactionOutputs []*UnspentTransactionOutput
	// counter to remember how many unspent transaction outputs were removed
	var removedUnspentOutputsCount = 0

	// go through transactions
	for _, newTransaction := range newTransactions {
		if newTransaction == nil {
			continue
		}
		// this transacton unspent outputs
		var currentUnspentTransactionOutputs = make([]*UnspentTransactionOutput, 0, len(newTransaction.Outputs))

		// go through transaction outputs and add new unspent transaction outputs
		for transactionOutputIndex, transactionOutput := range newTransaction.Outputs {

			// create unspent output
			currentUnspentTransactionOutputs = append(currentUnspentTransactionOutputs, &UnspentTransactionOutput{
				OutputID:    newTransaction.ID,
				OutputIndex: transactionOutputIndex,
				Address:     transactionOutput.Address,
				Amount:      transactionOutput.Amount,
			})
		}

		// append to the outputs from other transactions
		newUnspentTransactionOutputs = append(newUnspentTransactionOutputs, currentUnspentTransactionOutputs...)

		// go through transaction inputs and remove existing transaction outputs
		for _, input := range newTransaction.Inputs {

			// if we already removed some we do not need to check them so we reduce the count with removedUnspentOutputsCount
			unspentTransactionOutputCount := len(unspentTransactionOutputs) - removedUnspentOutputsCount

			for unspentOutputIndex := 0; unspentOutputIndex < unspentTransactionOutputCount; unspentOutputIndex++ {
				// get current output beeing checked
				unspentOutput := unspentTransactionOutputs[unspentOutputIndex]

				// check if unspent transaction output matches transaction input and 'remove' it
				if input.IsMatch(unspentOutput) {
					// get the index of transaction to swap with
					// if last was already swapped then the transaction to swap should be second to last
					lastIndex := len(unspentTransactionOutputs) - 1 - removedUnspentOutputsCount

					// swap with last
					unspentTransactionOutputs[unspentOutputIndex] = unspentTransactionOutputs[lastIndex]
					// to prevent memory leak
					unspentTransactionOutputs[lastIndex] = nil

					// decrease index to check swapped unspent transaction, next iteration we check the element which was swapped
					unspentOutputIndex--
					// remember that we deleted one so that we can cut deleted transactions out later
					removedUnspentOutputsCount++
				}
			}
		}
	}

	// slice removed out of the array
	var lastRemovedIndex = len(unspentTransactionOutputs) - removedUnspentOutputsCount
	if lastRemovedIndex > -1 {
		unspentTransactionOutputs = unspentTransactionOutputs[:lastRemovedIndex]
	}

	// return combined constructed list
	return append(newUnspentTransactionOutputs, unspentTransactionOutputs...)
}
