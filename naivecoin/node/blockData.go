package node

import "naivecoin/transactions"

//BlockData : Defines data sent for block mining.
type BlockData struct {
	Message      string                      `json:"message" binding:"required"`
	Transactions []*transactions.Transaction `json:"transactions" binding:"required"`
}
