package node

//BlockData : Defines data sent for block mining.
type BlockData struct {
	Message string   `json:"message" binding:"required"`
	Amount  *float64 `json:"amount" binding:"required"`
	Address *string  `json:"address" binding:"required"`
}
