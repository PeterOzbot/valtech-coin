package wallet

import (
	"naivecoin/transactions"
)

//GenerateTransaction :  Creates transaction for specific address with inputs from owner address and outputs to receiver address or owner if there is some remaining amount.
func GenerateTransaction(receiverAddress string, ownerAddress *Address, amount float64, unspentTransactionOutputs []*transactions.UnspentTransactionOutput) (*transactions.Transaction, error) {
	// if address is nil return
	if ownerAddress == nil {
		return nil, nil
	}

	// filter unspent transaction outputs
	filteredUnspentOutputs := FilterUnspentTransactionOutput(unspentTransactionOutputs, ownerAddress.PublicKey)

	// get the outputs that match the desired amount
	selectedUnspentOutputs, leftoverAmount := GetUnspentOutputCombination(filteredUnspentOutputs, amount)

	//check the amounts
	if !validateAmount(selectedUnspentOutputs, leftoverAmount, amount) {
		return nil, nil
	}

	// create new inputs from selected unspent outputs
	newInputs := make([]*transactions.TransactionInput, len(selectedUnspentOutputs))
	for index, selectedOutput := range selectedUnspentOutputs {
		newInputs[index] = &transactions.TransactionInput{
			OutputID:    selectedOutput.OutputID,
			OutputIndex: selectedOutput.OutputIndex,
		}
	}

	// create new outputs
	var newOutputs []*transactions.TransactionOutput

	// create receiver output
	receiverOutput := &transactions.TransactionOutput{
		Address: receiverAddress,
		Amount:  amount,
	}
	// if there is some leftover amount then we need to create another output to transfer amount back to owner
	if leftoverAmount != 0 {
		ownerOutput := &transactions.TransactionOutput{
			Address: ownerAddress.PublicKey,
			Amount:  leftoverAmount,
		}
		newOutputs = []*transactions.TransactionOutput{receiverOutput, ownerOutput}
	} else {
		newOutputs = []*transactions.TransactionOutput{receiverOutput}
	}

	// create transaction
	newTransaction := &transactions.Transaction{
		Inputs:  newInputs,
		Outputs: newOutputs,
	}

	// calculate the ID for transaction
	newTransaction.ID = newTransaction.CalculateID()

	// check if inputs can be signed and sign them
	for _, inputTransaction := range newTransaction.Inputs {

		// check if input can be signed
		if transactions.CanSign(unspentTransactionOutputs, inputTransaction, ownerAddress.PublicKey) {

			// sign input
			var err error
			inputTransaction.Signature, err = newTransaction.Sign(ownerAddress.PrivateKey)

			// check of error
			if err != nil {
				return nil, err
			}
		} else {
			// if transaction cant be signed then return
			return nil, nil
		}
	}

	// return new transation
	return newTransaction, nil
}

// validate the amount from selected outputs
// the amount + leftoverAmount must equal total unspent outputs amount
func validateAmount(selectedUnspentOutputs []*transactions.UnspentTransactionOutput, leftoverAmount float64, amount float64) bool {
	// initialize amount
	var selectedAmount = 0.0

	// add amounts from selected outputs
	for _, selectedOutput := range selectedUnspentOutputs {
		selectedAmount += selectedOutput.Amount
	}

	// return if the amounts are equal
	return selectedAmount == leftoverAmount+amount
}
