package main

import (
	"naivecoin/blockchain"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/blocks", getBlockChain)

	router.Run()
}

func getBlockChain(c *gin.Context) {
	genesisBlock := blockchain.GenesisBlock()
	c.JSON(200, genesisBlock)
}
