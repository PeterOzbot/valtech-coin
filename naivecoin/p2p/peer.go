package p2p

//Peer : Defines single peer connected to the network.
type Peer struct {
	Address string `json:"address" binding:"required"`
	Port    int    `json:"port" binding:"required"`
}
