package p2p

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

//TryInitializeCallerPeer : Initializes peer connected to this node if the peer is not already connected.
func TryInitializeCallerPeer(peerData *PeerData) (*SocketInfo, bool) {

	// combine server websocket url
	url := url.URL{Scheme: "ws", Host: peerData.Address, Path: "/ws"}
	fmt.Println("connecting to: ", url.String())

	// construct header with url to notify server with caller URL
	requestHeader := http.Header{}
	requestHeader.Add(IdentifierHeader, Identifier.ID)

	// create connection
	conn, response, err := websocket.DefaultDialer.Dial(url.String(), requestHeader)
	if err != nil {

		fmt.Println("websocket dial: ", err)
		return nil, false
	}

	// get caller id
	callerID := response.Header.Get(IdentifierHeader)

	// common initialization
	return InitializePeer(callerID, conn)
}
