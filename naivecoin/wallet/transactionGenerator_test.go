package wallet

import (
	"naivecoin/transactions"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Test_GenerateTransaction_OwnerAddresNil : Test the result when owner address is nil.
func Test_GenerateTransaction_OwnerAddresNil(t *testing.T) {
	// inputs
	var receiverAddress = "receiver address"
	var ownerAddress *Address
	var amount = 2.5
	var unspentTransactionOutputs = []*transactions.UnspentTransactionOutput{
		{
			OutputID:    "1",
			OutputIndex: 1,
			Address:     "address",
			Amount:      2.5,
		},
	}

	// get transaction
	transaction, _ := GenerateTransaction(receiverAddress, ownerAddress, amount, unspentTransactionOutputs)

	// check result
	if transaction != nil {
		t.Errorf("generated transaction should be nil when no owner address")
	}
}

// Test_GenerateTransaction_ZeroAmount : Tests that empty transaction is generated when there is no amount.
func Test_GenerateTransaction_ZeroAmount(t *testing.T) {
	// inputs
	var receiverAddress = "receiver address"
	var ownerAddress = &Address{
		PublicKey:  testPublicKey,
		PrivateKey: testPrivateKey,
	}
	var amount = 0.0
	var unspentTransactionOutputs = []*transactions.UnspentTransactionOutput{
		{
			OutputID:    "1",
			OutputIndex: 1,
			Address:     "address",
			Amount:      2.5,
		},
	}

	// expected result
	expectedTransaction := &transactions.Transaction{
		ID:      "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		Inputs:  []*transactions.TransactionInput{},
		Outputs: []*transactions.TransactionOutput{},
	}

	// get transaction
	actualTransaction, _ := GenerateTransaction(receiverAddress, ownerAddress, amount, unspentTransactionOutputs)

	// check result
	if actualTransaction == nil {
		t.Errorf("generated transaction should not be nil")
		return
	}

	// check Inputs length
	if len(actualTransaction.Inputs) != len(expectedTransaction.Inputs) {
		t.Errorf("generated transaction Inputs length not correct. Expected: %d Actual: %d", len(expectedTransaction.Inputs), len(actualTransaction.Inputs))
	}
	// check Outputs length
	if len(actualTransaction.Outputs) != len(expectedTransaction.Outputs) {
		t.Errorf("generated transaction Outputs length not correct. Expected: %d Actual: %d", len(expectedTransaction.Inputs), len(actualTransaction.Inputs))
	}
}

// Test_GenerateTransaction_NotEnoughAmount : Test if nil is returned when there is not enough amount in the unspent outputs.
func Test_GenerateTransaction_NotEnoughAmount(t *testing.T) {
	// inputs
	var receiverAddress = "receiver address"
	var ownerAddress = &Address{
		PublicKey:  testPublicKey,
		PrivateKey: testPrivateKey,
	}
	var amount = 10.5
	var unspentTransactionOutputs = []*transactions.UnspentTransactionOutput{
		{
			OutputID:    "1",
			OutputIndex: 1,
			Address:     testPublicKey,
			Amount:      2.5,
		},
	}

	// get transaction
	transaction, _ := GenerateTransaction(receiverAddress, ownerAddress, amount, unspentTransactionOutputs)

	// check result
	if transaction != nil {
		t.Errorf("generated transaction should be nil when there is not enough amount in unspent transaction outputs")
	}
}

// Test_GenerateTransaction_UnspentOutputsToInputs : Test that the unspent outputs are correctly converted to inputs.
func Test_GenerateTransaction_CorrectTransactionInputs(t *testing.T) {
	// inputs
	var receiverAddress = "receiver address"
	var ownerAddress = &Address{
		PublicKey:  testPublicKey,
		PrivateKey: testPrivateKey,
	}
	var amount = 5.0
	var unspentTransactionOutputs = []*transactions.UnspentTransactionOutput{
		{
			OutputID:    "1",
			OutputIndex: 1,
			Address:     testPublicKey,
			Amount:      2.5,
		},
		{
			OutputID:    "2",
			OutputIndex: 2,
			Address:     testPublicKey,
			Amount:      2.5,
		},
	}

	// expected result
	expectedTransaction := &transactions.Transaction{
		ID: "412626e3289ead11dd6b36fa6805a5437112b1b70649053319919841e27bd7b7",
		Inputs: []*transactions.TransactionInput{
			{
				OutputID:    "1",
				OutputIndex: 1,
				Signature:   "1b3545c8fc674d204d2858d1a1b077c8f64c75e2f81f704135fe344d039caff37a3d2790a42d851274eceb533d199b34c875703b6445dc4c1526ab2c1e22c6e06eb7fe09e157827915c46be0c71a354d42c82cbd09fe2d17f619902262582fdf90362a2c15b46343799164d2a66ca42e312cca0eae7b49c30284bf2a52661014630d7bc30bbf351fe9f70766a2c8a90ec8862afb58517526f71657cc268efe2fe23edb58a83e6fd05c01f875f1da1fa0821165a61bf761b78abe7960d2500f0c08a713f7606ad796cc423a7d5ce69dbb48a45406619e4da345913586fb273ebfe08e8b589e389e12b6b15f9582e962eaf4e60b7fb69047f6a12c381e5971725fb41d9b9dec503a37f78b210fdf7112368da88df553ac8ec8a23a837d9c57ecff4e140f0db473a64fa5cbafa4348b13746101b16c66c8095dcf0e134c46b8f193d9fcd2533706ec88d7be04cf1014ade3ab3438b244de7bd3e7689a6dc05351e2e0350763fc4c2399f9337fa6d32c63322ae971d7c0025b260d51ae50c954315e25dc9bfe8b552a9f55ddb114074655afb0dfe1e84c26f9bc68b0b55435f6ce64b51eb54b5a7aaf26425bfa0e7001ec668b0c9c83d463f5a92f38a8bf79c1ebe191afa5edf555dafd89f809f79daf4523783ab81402cc231a5bdf071ab78a9eb263d82ecd338864fa3dcab5f2204b0bf103be1fac1ccef802e2bcf88b1fa5963b",
			},
			{
				OutputID:    "2",
				OutputIndex: 2,
				Signature:   "1b3545c8fc674d204d2858d1a1b077c8f64c75e2f81f704135fe344d039caff37a3d2790a42d851274eceb533d199b34c875703b6445dc4c1526ab2c1e22c6e06eb7fe09e157827915c46be0c71a354d42c82cbd09fe2d17f619902262582fdf90362a2c15b46343799164d2a66ca42e312cca0eae7b49c30284bf2a52661014630d7bc30bbf351fe9f70766a2c8a90ec8862afb58517526f71657cc268efe2fe23edb58a83e6fd05c01f875f1da1fa0821165a61bf761b78abe7960d2500f0c08a713f7606ad796cc423a7d5ce69dbb48a45406619e4da345913586fb273ebfe08e8b589e389e12b6b15f9582e962eaf4e60b7fb69047f6a12c381e5971725fb41d9b9dec503a37f78b210fdf7112368da88df553ac8ec8a23a837d9c57ecff4e140f0db473a64fa5cbafa4348b13746101b16c66c8095dcf0e134c46b8f193d9fcd2533706ec88d7be04cf1014ade3ab3438b244de7bd3e7689a6dc05351e2e0350763fc4c2399f9337fa6d32c63322ae971d7c0025b260d51ae50c954315e25dc9bfe8b552a9f55ddb114074655afb0dfe1e84c26f9bc68b0b55435f6ce64b51eb54b5a7aaf26425bfa0e7001ec668b0c9c83d463f5a92f38a8bf79c1ebe191afa5edf555dafd89f809f79daf4523783ab81402cc231a5bdf071ab78a9eb263d82ecd338864fa3dcab5f2204b0bf103be1fac1ccef802e2bcf88b1fa5963b",
			},
		},
		Outputs: []*transactions.TransactionOutput{
			{
				Address: receiverAddress,
				Amount:  5,
			},
		},
	}

	// get transaction
	actualTransaction, _ := GenerateTransaction(receiverAddress, ownerAddress, amount, unspentTransactionOutputs)

	// check result
	if actualTransaction == nil {
		t.Errorf("generated transaction should not be nil")
		return
	}

	// check Inputs length
	if len(actualTransaction.Inputs) != len(expectedTransaction.Inputs) {
		t.Errorf("generated transaction Inputs length not correct. Expected: %d Actual: %d", len(expectedTransaction.Inputs), len(actualTransaction.Inputs))
	}
	// validate each
	for index, actualInput := range actualTransaction.Inputs {
		// validate expected
		if !cmp.Equal(actualInput, expectedTransaction.Inputs[index]) {
			t.Errorf("transaction input does not match")
		}
	}
}

// Test_GenerateTransaction_UnspentOutputsToInputs : Test that the unspent outputs are correctly converted to outputs.
func Test_GenerateTransaction_CorrectTransactionOutputs(t *testing.T) {
	// inputs
	var receiverAddress = "receiver address"
	var ownerAddress = &Address{
		PublicKey:  testPublicKey,
		PrivateKey: testPrivateKey,
	}
	var amount = 4.5
	var unspentTransactionOutputs = []*transactions.UnspentTransactionOutput{
		{
			OutputID:    "1",
			OutputIndex: 1,
			Address:     testPublicKey,
			Amount:      2.5,
		},
		{
			OutputID:    "2",
			OutputIndex: 2,
			Address:     testPublicKey,
			Amount:      2.5,
		},
	}

	// expected result
	expectedTransaction := &transactions.Transaction{
		ID: "cf7ec82d853afbfe017e84156ea3f48a038ab11843cef8c5676b36cce7f49e6d",
		Inputs: []*transactions.TransactionInput{
			{
				OutputID:    "1",
				OutputIndex: 1,
				Signature:   "4da8e07a93309c43d4c2de1b5cf441526bb1b9d83b73b2bf8325a1173d3d405ae0f65a8ce3be17986edc2db2b0ffa6cd8a726bf6e730ea7aa8a7b0ece3c6ac43ebdd9c8939a9cd015f92c372581fcc7d768ccca64388bee3491444668690a955962d42cfb0cf9854a0f61528288e547d48ff9f7d2531a92c3e22d870b51f534c96aecc5c8f6a56ba15ebac640c9c42bad1ca25ba3ec065b603517c89a6c9f26847daa68acc6037f1af3649c013f618c4df1b54c125231f28d64ce20283e5ed392a6d80a8ad2ec65f4f2f61ba27472922b5dfadfc8a060d4cd0cb98d4db1ece2cca9950a907edfc34127bd78877df8bdaf3053c6d46e66f3b66d03d66426bf106f7aac975997dfc550e459d9dfadcce924ae3a34880635ec6b3640766285dce496c41ec3f3ced16e85650e2a543018519b4f467f15bdb3466aa04ea1b1231d4be9d98ae68ff16d227a703d88ffedac68fe07ffba63afb89023a425011a43776d572dc398b68720fb6e199ba3865e38809ab1b9160f5f217e735248c03d93e9051aa40ee73c7044676b1964b5bc31e67f4beaafe213b3bf79cc9593a868e6c73ce74fb7cbfa8c4b648b7434f3097d2db7239ad55092bfe0da97391e0a2f7b7ca518fa89e8ce8eb52e6de5c4208654992522b57a1fee47132930c1b5f9e96c006baebb7e83656989a22bd3d9e654d332b4a9b268c274c787b65a88a421fa2ab0f20",
			},
			{
				OutputID:    "2",
				OutputIndex: 2,
				Signature:   "4da8e07a93309c43d4c2de1b5cf441526bb1b9d83b73b2bf8325a1173d3d405ae0f65a8ce3be17986edc2db2b0ffa6cd8a726bf6e730ea7aa8a7b0ece3c6ac43ebdd9c8939a9cd015f92c372581fcc7d768ccca64388bee3491444668690a955962d42cfb0cf9854a0f61528288e547d48ff9f7d2531a92c3e22d870b51f534c96aecc5c8f6a56ba15ebac640c9c42bad1ca25ba3ec065b603517c89a6c9f26847daa68acc6037f1af3649c013f618c4df1b54c125231f28d64ce20283e5ed392a6d80a8ad2ec65f4f2f61ba27472922b5dfadfc8a060d4cd0cb98d4db1ece2cca9950a907edfc34127bd78877df8bdaf3053c6d46e66f3b66d03d66426bf106f7aac975997dfc550e459d9dfadcce924ae3a34880635ec6b3640766285dce496c41ec3f3ced16e85650e2a543018519b4f467f15bdb3466aa04ea1b1231d4be9d98ae68ff16d227a703d88ffedac68fe07ffba63afb89023a425011a43776d572dc398b68720fb6e199ba3865e38809ab1b9160f5f217e735248c03d93e9051aa40ee73c7044676b1964b5bc31e67f4beaafe213b3bf79cc9593a868e6c73ce74fb7cbfa8c4b648b7434f3097d2db7239ad55092bfe0da97391e0a2f7b7ca518fa89e8ce8eb52e6de5c4208654992522b57a1fee47132930c1b5f9e96c006baebb7e83656989a22bd3d9e654d332b4a9b268c274c787b65a88a421fa2ab0f20",
			},
		},
		Outputs: []*transactions.TransactionOutput{
			{
				Address: receiverAddress,
				Amount:  4.5,
			},
			{
				Address: ownerAddress.PublicKey,
				Amount:  0.5,
			},
		},
	}

	// get transaction
	actualTransaction, _ := GenerateTransaction(receiverAddress, ownerAddress, amount, unspentTransactionOutputs)

	// check result
	if actualTransaction == nil {
		t.Errorf("generated transaction should not be nil")
		return
	}

	// check Outputs length
	if len(actualTransaction.Outputs) != len(expectedTransaction.Outputs) {
		t.Errorf("generated transaction Outputs length not correct. Expected: %d Actual: %d", len(expectedTransaction.Outputs), len(actualTransaction.Outputs))
	}
	// validate each
	for index, actualOutput := range actualTransaction.Outputs {
		// validate expected
		if !cmp.Equal(actualOutput, expectedTransaction.Outputs[index]) {
			t.Errorf("transaction output does not match")
		}
	}
}
