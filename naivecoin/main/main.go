package main

import (
	"naivecoin/node"
	"naivecoin/p2p"

	"flag"

	"github.com/gin-gonic/gin"
)

var httpPort = flag.String("port", ":8080", "Port on which the http node runs.")
var p2pPort = flag.String("port", ":9090", "Port on which the p2p node runs.")

func main() {

	// read flags
	flag.Parse()

	// initialize servers
	initHTTP()
	initP2P()
}

func initHTTP() {
	// initialize routing
	router := gin.Default()

	// set the routs
	router.GET("/blocks", node.GetBlockChain)
	router.POST("/mineblock", node.Mineblocks)
	router.GET("/peers", p2p.GetPeers)
	router.POST("/addPeer", p2p.AddPeers)

	router.Run(*httpPort)
}

func initP2P() {

	//p2pPort
}
