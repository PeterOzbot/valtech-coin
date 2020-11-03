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
	if doesPeerExists(callerID) {
		conn.Close()
		return nil, true
	}

	// create new info
	socketInfo := &SocketInfo{
		Connection: conn,
		ID:         callerID,
	}

	// hook send message
	socketInfo.SendMessage = func(message string) {
		err = conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			fmt.Println("send message : ", err)
		}
	}

	// hook on message received
	socketInfo.OnMessageReceived = func(message string) {
		fmt.Println(message)
	}

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("error reading: ", err)
				break
			}
			if socketInfo.OnMessageReceived != nil {
				socketInfo.OnMessageReceived("FROM PEER : " + string(message))
			}
		}
	}()

	return socketInfo, false
}
