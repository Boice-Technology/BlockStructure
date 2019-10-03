// To check whether a transaction has not being modified given merkle tree root hash and the transactions.

package merkle_tree

func AreTransactionsIntact(transactions []string, merkleTreeRootHash string) bool {
	var verdict bool
	merkleTree := CreateMerkleTree(transactions)
	rootHash := merkleTree[0][0]
	if rootHash == merkleTreeRootHash {
		verdict = true
	} else {
		verdict = false
	}
	return verdict
}