package p2p

//PeerData : Defines single peer connected to the network.
type PeerData struct {
	ID      string `json:"id"`
	Address string `json:"address" binding:"required"`
}
