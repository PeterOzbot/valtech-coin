package blockchain

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

// GeneratesCorrectBlock : Tests if the generator generates correct block from data and previous block.
func Test_GeneratesCorrectBlock(t *testing.T) {
	// create example block
	var latestBlock *Block = &Block{
		index:        15,
		data:         "Dober dan gospod kamplan.",
		previousHash: "ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae",
		timestamp:    time.Unix(1588430083866862500, 0),
		hash:         "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb",
	}
	newBlockData := "My new block data"
	// create expected block
	var expectedBlock *Block = &Block{
		index:        16,
		data:         newBlockData,
		previousHash: "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb",
		timestamp:    time.Unix(1588430083866862555, 0),
		hash:         "079f8214cc7e47783b8ce6656de68002040c893afccebb99e7dfdc6df3729ccc",
	}

	// hash block
	result := GenerateBlock(newBlockData, latestBlock)

	// result should be nil
	if !cmp.Equal(result, expectedBlock) {
		t.Errorf("new block generator did not generate correct block")
	}
}
