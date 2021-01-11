package main

import (
	"fmt"
	"naivecoin/node"
	"naivecoin/p2p"

	"flag"

	"github.com/gin-gonic/gin"
)

var port = flag.String("port", "8080", "Port on which the http node runs.")

func main() {

	// read flags
	flag.Parse()

	// initialize routing
	router := gin.Default()

	// initialize http and p2p
	initHTTP(router)
	initP2P(router)

	// initialize this node
	if !node.InitializeNode(*port) {
		return
	}

	// run
	router.Run(fmt.Sprintf(":%s", *port))
}

func initHTTP(router *gin.Engine) {

	// set the routs
	router.GET("/blocks", node.GetBlockchain)
	router.POST("/mineblock", node.MineBlock)
	router.GET("/peers", p2p.GetPeers)
	router.POST("/addPeer", p2p.AddPeer)
}

func initP2P(router *gin.Engine) {
	router.GET("/ws", p2p.AddServerPeer)
}
