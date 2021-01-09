package wallet

import "testing"

// Test_GenerateAddress_GeneratesAddress : Tests address gets generated without failing.
func Test_GenerateAddress_GeneratesAddress(t *testing.T) {

	// get address
	address, err := GenerateAddress()

	// check result
	if err != nil {
		panic(err)
	}
	if address == nil {
		t.Errorf("address was not generated")
	}
}
