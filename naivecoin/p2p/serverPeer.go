package p2p

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

//InitializeServerPeer : Sets up web socket serving as Server, identifying this node.
func InitializeServerPeer(c *gin.Context, callerID string) (*SocketInfo, bool) {
	// initialize upgrader
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// this is short-cut to make it work with browser
		// CheckOrigin: func(r *http.Request) bool {
		// 	return true
		// },
	}

	// construct header with url to notify caller with server URL
	requestHeader := http.Header{}
	requestHeader.Add(IdentifierHeader, Identifier.ID)

	// hook to get connection
	conn, err := upgrader.Upgrade(c.Writer, c.Request, requestHeader)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: ", err)
		return nil, false
	}

	// check if already connected
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
	socketInfo.OnMessageReceived = func(requestMessage string) {
		OnMessageReceived(requestMessage, socketInfo)
	}

	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Failed to read message: ", err)
				break
			}
			if socketInfo.OnMessageReceived != nil {
				socketInfo.OnMessageReceived(string(msg))
			}
		}
	}()

	// return
	return socketInfo, false
}
