// Generates a transaction branch of merkle tree for particular transaction.

package merkle_tree

func CreateTransactionBranch(transactions []string, transaction string) [][]string {
	merkleTree := CreateMerkleTree(transactions)
	length := len(merkleTree)
	for i := 0 ; i < length/2 ; i ++ {
		merkleTree[i], merkleTree[length - i - 1] = merkleTree[length - i -1], merkleTree[i]
	}
	var index int
	for i := 0 ;i < len(merkleTree[0]) ; i++ {
		if(merkleTree[0][i]==transaction){
			index = i
			break
		}
	}
	list := make([]string, 0)
	list = append(list,merkleTree[0][index])
	transactionBranch := make([][]string, 0)
	if index % 2 == 0{
		list = append(list, merkleTree[0][index+1])
	} else {
		list = append(list, merkleTree[0][index-1])
	}
	transactionBranch = append(transactionBranch, list)
	for i := 0 ; i < len(merkleTree) - 2 ; i++ {
		index = index / 2
		list := make([]string, 0)
		list = append(list,merkleTree[i+1][index])
		if index % 2 == 0{
			list = append(list, merkleTree[i+1][index+1])
		} else {
			list = append(list, merkleTree[i+1][index-1])
		}
		transactionBranch = append(transactionBranch, list)
	}
	transactionBranch = append(transactionBranch, merkleTree[length-1])
	return transactionBranch
}
