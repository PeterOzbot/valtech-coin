package wallet

import "naivecoin/transactions"

//FilterUnspentTransactionOutput : Returns unspent transaction outputs that matches to the address input.
func FilterUnspentTransactionOutput(unspentTransactionOutputs []*transactions.UnspentTransactionOutput, address string) []*transactions.UnspentTransactionOutput {

	// initialize result
	var filteredUnspentOutputs []*transactions.UnspentTransactionOutput

	// loop through and add matching outputs
	for _, unspentTransaction := range unspentTransactionOutputs {

		// check address
		if unspentTransaction.Address == address {
			filteredUnspentOutputs = append(filteredUnspentOutputs, unspentTransaction)
		}
	}

	// return what was found
	return filteredUnspentOutputs
}
