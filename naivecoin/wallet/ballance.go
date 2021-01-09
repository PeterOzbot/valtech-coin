package wallet

import (
	"naivecoin/transactions"
)

//GetBalance :  Returns ballance from all unspent transaction outputs for specific address.
func GetBalance(address *Address, unspentTransactionOutputs []*transactions.UnspentTransactionOutput) float64 {
	if address == nil {
		return 0
	}

	// filter by address
	filteredUnspentOutputs := FilterUnspentTransactionOutput(unspentTransactionOutputs, address.PublicKey)

	// init balance amount
	var addressBalance float64 = 0

	// go through unspent transaction outputs and calculate amounts from those that match the public key
	for _, unspentOutput := range filteredUnspentOutputs {
		addressBalance += unspentOutput.Amount
	}

	// return balance
	return addressBalance
}
