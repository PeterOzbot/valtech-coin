package node

import (
	"naivecoin/blockchain"
	"naivecoin/transactions"
	"naivecoin/wallet"
)

//CurrentBlockchain : Current block chain.
var CurrentBlockchain []*blockchain.Block

//UnspentTransactionOutputs : Collection of current unspent transaction outputs
var UnspentTransactionOutputs []*transactions.UnspentTransactionOutput

//Wallet : This node wallet.
var Wallet *wallet.Address
