package transactions

import "testing"

// Test_Validate_ID : Tests validation checks transaction ID.
func Test_Validate_ID(t *testing.T) {

	// create inputs
	var transaction = &Transaction{
		ID: "invalid-ID",
		Inputs: []*TransactionInput{
			{
				OutputID:    "OutputID",
				OutputIndex: 0,
				Signature:   "",
			},
		},
		Outputs: []*TransactionOutput{
			{
				Address: "address",
				Amount:  10,
			},
		},
	}
	var unspentTransactionOutputs []*UnspentTransactionOutput

	// validate
	result, _ := transaction.IsValid(unspentTransactionOutputs)

	// result should false
	if result {
		t.Errorf("transaction is invalid so result should be false")
	}
}

// Test_Validate_UnspentOutputsAmountEqualsTransactionOutputs : Tests that validation checks that the combined unspent outputs amount equals transaction outputs amount. Unspent outputs are found with transaction inputs.
func Test_Validate_UnspentOutputsAmountEqualsTransactionOutputs(t *testing.T) {

	// create inputs
	var transaction = &Transaction{
		ID: "260bb98233c75fad443812c04174f5aeb5e43183ff27877f7f873f9978f8e017",
		Inputs: []*TransactionInput{
			{
				OutputID:    "OutputID",
				OutputIndex: 5,
				Signature:   "16c9db3eb0a5dd2c444ff1c961d5eed4da6fec1f77d1131a223c20424f7582892064345eb71273f9fede291121837ed3901ccc1c68609cc718203ea680833dd9b6a9c4653de7df0cfa3ad114d4c8adace3e6374d688bdb755ca2e7c92a0368b618ec1f546c8c32e9d3b8d339bfb674388bc38d943fe7a0c0cba18f5d76b94382492c03a9bc1c555b741d3162a227ab14d6cecd064c9d04e0fc586f548421d0dfd56559f4333e6d1d29c0a3d0d3f67088deaa6a5f4dee8462856fd88836243d10aba9ec643435f426c969b509eb60bfa3960b8b4ebd0d1cd1d12116c0e691b84bd5c620a4bb577ed1d4dfe10a31787c01f92cb99943308d31dcda230884f9768a2f6e0fea94b87d81ea554d984f105ae8602b4bfd749a9373fd662bd32612f3cb1af8f616fa366d2ce8d850d465ccde49d052401e8f8e58cab0ca6ee8b7846f8a67005451dce4459ad875f744a313e517bab6dd3d9afc0164640e5f0adcd78822572321b1293a7f6af1b83c5de8d165d80b12bf665d1f9915689d08fbff5a83e9ff7acdf39e1ea771003d0a47d312b3c5f03f21be25ecfb435a4e8b47a702dac0760509b93823cc6c80906af8db26625d6b6a8991577a307555c760db6819a7b12202a90d448769e6d8b7c34f06f3de09acd2d6b21833ac43346e28487ca76b678bbc0e9e8a3d1c54a7d54493902edb45ceccd46574ecbd97502c5e052f56e86d",
			},
		},
		Outputs: []*TransactionOutput{
			{
				Address: "address",
				Amount:  100,
			},
		},
	}
	var unspentTransactionOutputs = []*UnspentTransactionOutput{
		{
			OutputID:    "OutputID",
			OutputIndex: 5,
			Address:     publicKey,
			Amount:      15,
		},
	}

	// validate
	result, _ := transaction.IsValid(unspentTransactionOutputs)

	// result should false
	if result {
		t.Errorf("transaction is invalid so result should be false")
	}
}

// Test_Validate_UnspentTransactionOutputDoesNotExists : Tests that validation checks that unspent transaction output exists. Unspent outputs are found with transaction inputs.
func Test_Validate_UnspentTransactionOutputDoesNotExists(t *testing.T) {

	// create inputs
	var transaction = &Transaction{
		ID: "8ba63eb9886292ca2b3c5a9a360b42b8459096c36f90dd6a1bbcb03a78ebb0b2",
		Inputs: []*TransactionInput{
			{
				OutputID:    "OutputID",
				OutputIndex: 5,
				Signature:   "",
			},
			// this input does not have unspent output so the validation should fail
			{
				OutputID:    "OutputID-1",
				OutputIndex: 1,
				Signature:   "",
			},
		},
		Outputs: []*TransactionOutput{
			{
				Address: "address",
				Amount:  100,
			},
		},
	}
	var unspentTransactionOutputs = []*UnspentTransactionOutput{
		{
			OutputID:    "OutputID",
			OutputIndex: 5,
			Address:     "address",
			Amount:      100,
		},
	}

	// validate
	result, _ := transaction.IsValid(unspentTransactionOutputs)

	// result should false
	if result {
		t.Errorf("transaction is invalid so result should be false")
	}
}

// Test_Validate_UnspentTransactionOutputAddressMatchesSignature : Tests that validation checks that validation ID is signed with correct private key. Validation is done with public key that is the unspent transaction output address. Unspent outputs are found with transaction inputs.
func Test_Validate_UnspentTransactionOutputAddressMatchesSignature(t *testing.T) {

	// create inputs
	var transaction = &Transaction{
		ID: "260bb98233c75fad443812c04174f5aeb5e43183ff27877f7f873f9978f8e017",
		Inputs: []*TransactionInput{
			{
				OutputID:    "OutputID",
				OutputIndex: 5,
				Signature:   "696e76616c69642d7369676e6174757265",
			},
		},
		Outputs: []*TransactionOutput{
			{
				Address: "address",
				Amount:  100,
			},
		},
	}
	var unspentTransactionOutputs = []*UnspentTransactionOutput{
		{
			OutputID:    "OutputID",
			OutputIndex: 5,
			Address:     publicKey,
			Amount:      100,
		},
	}

	// validate
	result, _ := transaction.IsValid(unspentTransactionOutputs)

	// result should false
	if result {
		t.Errorf("transaction is invalid so result should be false")
	}
}

// Test_Validate_IsValid : Tests validation correctly returns that transaction is valid when it is.
func Test_Validate_IsValid(t *testing.T) {

	// create inputs
	var transaction = &Transaction{
		ID: "260bb98233c75fad443812c04174f5aeb5e43183ff27877f7f873f9978f8e017",
		Inputs: []*TransactionInput{
			{
				OutputID:    "OutputID",
				OutputIndex: 5,
				Signature:   "16c9db3eb0a5dd2c444ff1c961d5eed4da6fec1f77d1131a223c20424f7582892064345eb71273f9fede291121837ed3901ccc1c68609cc718203ea680833dd9b6a9c4653de7df0cfa3ad114d4c8adace3e6374d688bdb755ca2e7c92a0368b618ec1f546c8c32e9d3b8d339bfb674388bc38d943fe7a0c0cba18f5d76b94382492c03a9bc1c555b741d3162a227ab14d6cecd064c9d04e0fc586f548421d0dfd56559f4333e6d1d29c0a3d0d3f67088deaa6a5f4dee8462856fd88836243d10aba9ec643435f426c969b509eb60bfa3960b8b4ebd0d1cd1d12116c0e691b84bd5c620a4bb577ed1d4dfe10a31787c01f92cb99943308d31dcda230884f9768a2f6e0fea94b87d81ea554d984f105ae8602b4bfd749a9373fd662bd32612f3cb1af8f616fa366d2ce8d850d465ccde49d052401e8f8e58cab0ca6ee8b7846f8a67005451dce4459ad875f744a313e517bab6dd3d9afc0164640e5f0adcd78822572321b1293a7f6af1b83c5de8d165d80b12bf665d1f9915689d08fbff5a83e9ff7acdf39e1ea771003d0a47d312b3c5f03f21be25ecfb435a4e8b47a702dac0760509b93823cc6c80906af8db26625d6b6a8991577a307555c760db6819a7b12202a90d448769e6d8b7c34f06f3de09acd2d6b21833ac43346e28487ca76b678bbc0e9e8a3d1c54a7d54493902edb45ceccd46574ecbd97502c5e052f56e86d",
			},
		},
		Outputs: []*TransactionOutput{
			{
				Address: "address",
				Amount:  100,
			},
		},
	}
	var unspentTransactionOutputs = []*UnspentTransactionOutput{
		{
			OutputID:    "OutputID",
			OutputIndex: 5,
			Address:     publicKey,
			Amount:      100,
		},
	}

	// validate
	result, _ := transaction.IsValid(unspentTransactionOutputs)

	// check if validation correctly returned true
	if !result {
		t.Errorf("transaction is valid so result should be true")
	}
}
