package transactions

//IsValidCoinbase : Coinbase transaction must be validated differently than 'normal' transaction.
func (transaction *Transaction) IsValidCoinbase(blockIndex int) bool {
	if transaction == nil {
		return false
	}

	// get the proper transaction ID
	var calculatedTransactionID = transaction.CalculateID()
	// check if transaction has proper ID
	if transaction.ID != calculatedTransactionID {
		return false
	}

	// validate transaction inputs
	if len(transaction.Inputs) != 1 {
		return false
	}
	// check the output index
	if transaction.Inputs[0].OutputIndex != blockIndex {
		return false
	}

	// validate transaction outputs
	if len(transaction.Outputs) != 1 {
		return false
	}
	// check amount
	if transaction.Outputs[0].Amount != CoinbaseAmount {
		return false
	}

	return true
}
