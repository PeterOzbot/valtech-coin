package p2p

import "github.com/gorilla/websocket"

//SocketInfo : Defines web socket info for p2p comunication.
type SocketInfo struct {
	//Connection : Connection used for connection to the peer
	Connection *websocket.Conn
	//ID : socket identifier
	ID string
	//OnMessageReceived : If set its executed when message is received.
	OnMessageReceived func(message string)
	//SendMessage : Used to send message through socket.
	SendMessage func(message string)
}
