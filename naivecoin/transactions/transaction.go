package transactions

//TransactionOutput : Transaction outputs specify where the coins are sent.
type TransactionOutput struct {
	Address string
	Amount  float64
}

//TransactionInput : Transaction inputs provide the information “where” the coins are coming from.
type TransactionInput struct {
	OutputID    string
	OutputIndex int
	Signature   string
}

// Transaction : Defines one transaction with inputs and outputs.
type Transaction struct {
	ID      string
	Inputs  []*TransactionInput
	Outputs []*TransactionOutput
}

// UnspentTransactionOutput : All inputs must refer to unspent transaction. Free coins. Can be always derived from current blockchain.
type UnspentTransactionOutput struct {
	OutputID    string
	OutputIndex int
	Address     string
	Amount      float64
}
