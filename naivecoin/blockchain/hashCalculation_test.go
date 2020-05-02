package blockchain

import (
	"testing"
	"time"
)

// Test_NilInpout : Tests if hashing handles nil input without failing.
func Test_NilInput(t *testing.T) {

	// create nil block
	var block *Block

	// hash block
	result := block.CalculateHash()

	// result should be empty string
	if len(result) != 0 {
		t.Errorf("hashing result of nil block should be empty string.")
	}
}

// Test_CorrectHash : Tests if hashing creates correct hash value.
func Test_CorrectHash(t *testing.T) {

	// create example block
	var block *Block = &Block{
		index:        15,
		data:         "Dober dan gospod kamplan.",
		previousHash: "ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae",
		timestamp:    time.Unix(1588430083866862500, 0),
	}

	//  is manually generated hash from block
	// index + previousHash + timestamp + data
	expected := "6086832ab82d2ec069001023957746a3648e791819fc407ed99859c2753f6beb"

	// hash block
	result := block.CalculateHash()

	// result should be correct hash
	if result != expected {
		t.Errorf("hashing result is incorrect, Actual: %s Expected: %s", result, expected)
	}
}
