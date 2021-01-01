package blockchain

import (
	"testing"
	"time"
)

// Test_CalculateHash_NilInpout : Tests if hashing handles nil input without failing.
func Test_CalculateHash_NilInput(t *testing.T) {

	// create nil block
	var block *Block

	// hash block
	result := block.CalculateHash()

	// result should be empty string
	if len(result) != 0 {
		t.Errorf("hashing result of nil block should be empty string.")
	}
}

// Test_CalculateHash_CorrectHash : Tests if hashing creates correct hash value.
func Test_CalculateHash_CorrectHash(t *testing.T) {

	// create example block
	var block *Block = &Block{
		Index:        15,
		Data:         "Dober dan gospod kamplan.",
		PreviousHash: "ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae",
		Timestamp:    time.Unix(1588430083866862500, 0),
		Difficulty:   5,
		Nonce:        32,
	}

	//  is manually generated hash from block
	expected := "d399b315465353ccf52ebe7b1cb6ac261368d300cddcd6de6f93dcac0190be20"

	// hash block
	result := block.CalculateHash()

	// result should be correct hash
	if result != expected {
		t.Errorf("hashing result is incorrect, Actual: %s Expected: %s", result, expected)
	}
}

// Test_MineBlock_MineBlock : Tests if hash is correctly mined.
func Test_MineBlock_MineBlock(t *testing.T) {
	var block = &Block{
		Index:        0,
		Hash:         "",
		Data:         "my genesis block!!",
		PreviousHash: "",
		Timestamp:    time.Unix(1465154705, 0),
		Difficulty:   5,
		Nonce:        0,
	}

	// hash block
	block.MineBlock()

	// result should be correct hash
	expectedHash := "07cd44e61db86462cd7727374f59a5c6cbe02c896746b44322e195d1f88b10f2"
	if block.Hash != expectedHash {
		t.Errorf("mining result is incorrect, Actual: %s Expected: %s", block.Hash, expectedHash)
	}
	expectedNonce := 32
	if block.Nonce != expectedNonce {
		t.Errorf("mining result is incorrect, Actual: %d Expected: %d", block.Nonce, expectedNonce)
	}
}
