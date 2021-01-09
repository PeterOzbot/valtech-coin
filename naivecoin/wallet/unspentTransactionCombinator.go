package wallet

import (
	"math"
	"naivecoin/transactions"
)

//GetUnspentOutputCombination : Finds combination of unspent transactions that match the total amount the best. If amount is not exact match then leftover amount is also returned
func GetUnspentOutputCombination(unspentOutput []*transactions.UnspentTransactionOutput, totalAmount float64) ([]*transactions.UnspentTransactionOutput, float64) {

	// find the combination that the most closely matches the total amount
	selectedUnspentOutputs := getBestCombination(unspentOutput, totalAmount)

	// get the amount from unspent outputs in found combination of them
	amountFound := 0.0
	for _, selectedUnspentOutput := range selectedUnspentOutputs {
		amountFound += selectedUnspentOutput.Amount
	}

	// calculate the leftover amount, don't take more less than 0
	leftoverAmount := math.Max(amountFound-totalAmount, 0)

	// return
	return selectedUnspentOutputs, leftoverAmount
}

// goes through all combination of unspent outputs and tries to find the one that matches total amount the most
// Lecture 10 | Programming Abstractions (Stanford)
// https://www.youtube.com/watch?v=NdF1QDTRkck
func getBestCombination(unspentOutput []*transactions.UnspentTransactionOutput, totalAmount float64) []*transactions.UnspentTransactionOutput {

	// initialize
	var bestCombination []*transactions.UnspentTransactionOutput
	var minimalDiff float64 = math.MaxFloat64
	var subsetCombinationFunc func(unspentOutputs []*transactions.UnspentTransactionOutput, targetAmount float64, currentIteration []*transactions.UnspentTransactionOutput) bool

	// define recurson function
	subsetCombinationFunc = func(unspentOutputs []*transactions.UnspentTransactionOutput, targetAmount float64, currentIteration []*transactions.UnspentTransactionOutput) bool {

		// get current iteration sum
		partialSum := 0.0
		for _, unspentOutput := range currentIteration {
			partialSum += unspentOutput.Amount
		}

		// calculate new difference
		newDiff := partialSum - targetAmount

		// check if this iteration is not too small
		if newDiff >= 0 {

			// check if this iteration difference is smaller and if it is equal choose the one that has more elements
			if newDiff < minimalDiff || (newDiff == minimalDiff && len(bestCombination) < len(currentIteration)) {
				// copy current iteration into best combination
				// with appending to nil we force new allocation
				bestCombination = append([]*transactions.UnspentTransactionOutput(nil), currentIteration...)

				// update difference
				minimalDiff = newDiff
			}
		}

		// if difference is zero then we can use this solution so we stop and unwind recursion
		// this can be enabled if we prefer speed over the fact that the iteration with more elements will be selected
		// if minimalDiff == 0 {
		// 	return true
		// }

		// go thorugh all iterations
		for index, unspetOutput := range unspentOutputs {

			// get next
			remaining := unspentOutputs[index+1:]
			newIteration := append(currentIteration, unspetOutput)

			// call recursively again
			var found = subsetCombinationFunc(remaining, targetAmount, newIteration)

			// if true is returned we stop the recursion and finish
			if found {
				return true
			}
		}

		// return false so we don't stop the loop
		return false
	}

	// start the recursion
	subsetCombinationFunc(unspentOutput, totalAmount, make([]*transactions.UnspentTransactionOutput, 0, len(unspentOutput)))

	// when recursion stops return the combination found
	return bestCombination
}
