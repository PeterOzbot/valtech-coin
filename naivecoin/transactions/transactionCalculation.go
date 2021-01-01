package transactions

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
)

//CalculateID : Calculates transaction ID from transaction outputs
func (transaction *Transaction) CalculateID() string {
	if transaction == nil {
		return ""
	}

	// init builder
	var transactionContentBuilder strings.Builder

	// combine transaction inputs
	for _, transactionInput := range transaction.Inputs {
		transactionContentBuilder.WriteString(transactionInput.OutputID)
		transactionContentBuilder.WriteString(strconv.Itoa(transactionInput.OutputIndex))
	}

	// combine transaction outputs
	for _, transactionOutput := range transaction.Outputs {
		transactionContentBuilder.WriteString(transactionOutput.Address)
		transactionContentBuilder.WriteString(strconv.FormatFloat(transactionOutput.Amount, 'f', 8, 64))
	}

	// get whole string representation
	stringRepresentation := transactionContentBuilder.String()

	// hash
	// create hasher and write the bytes to it
	hasher := sha256.New()
	hasher.Write([]byte(stringRepresentation))

	// return string
	return hex.EncodeToString(hasher.Sum(nil))
}
