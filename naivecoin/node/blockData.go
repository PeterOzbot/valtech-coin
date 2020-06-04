package node

//BlockData : Defines data sent for block mining.
type BlockData struct {
	Data string `json:"data" binding:"required"`
}
