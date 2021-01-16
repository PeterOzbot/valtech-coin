package p2p

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

//InitializeServerPeer : Sets up web socket serving as Server, identifying this node.
func InitializeServerPeer(c *gin.Context, callerID string) (*SocketInfo, bool) {
	// initialize websocket upgrader
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

	// common initialization
	return InitializePeer(callerID, conn)
}
