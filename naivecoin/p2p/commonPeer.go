package p2p

//OnMessageReceived : Handles message when its received from peers.
var OnMessageReceived func(requestMessage string, socketInfo *SocketInfo)

//OnPeerConnected : Executes actions that are required when peer gets conencted to this node.
var OnPeerConnected func(socketInfo *SocketInfo)
