package transactions

//CoinbaseTransaction :  Generates new coinbase transaction.
func CoinbaseTransaction(address string, blockIndex int, coinbaseAmount float64) *Transaction {

	// create coinbase transaction
	var coinbaseTransaction = &Transaction{
		Inputs: []*TransactionInput{
			{
				OutputIndex: blockIndex,
			},
		},
		Outputs: []*TransactionOutput{
			{
				Address: address,
				Amount:  coinbaseAmount,
			},
		},
	}

	// calculate ID
	coinbaseTransaction.ID = coinbaseTransaction.CalculateID()

	// return
	return coinbaseTransaction
}
