// Verify whether a transaction is included or not from the given transaction branch

package merkle_tree

func VerifyTransactionBranch(transactionBranch [][]string, transaction string) bool {
	var verdict bool = true
	if (transaction == transactionBranch[0][0]) || (transaction == transactionBranch[0][1]) {
		length := len(transactionBranch)
		for i := 0 ; i < length - 2 ; i++ {
			checkHash := Hash(transactionBranch[i][0] + transactionBranch[i][1])
			if (checkHash == transactionBranch[i+1][0]) || (checkHash == transactionBranch[i+1][1]) {
				continue
			} else{
				verdict = false
				break
			}
		}
		checkHash := Hash(transactionBranch[length-2][0] + transactionBranch[length-2][1])
		if checkHash != transactionBranch[length-1][0]{
			verdict = false
		}
	} else{
		verdict = false
	}
	return verdict
}

