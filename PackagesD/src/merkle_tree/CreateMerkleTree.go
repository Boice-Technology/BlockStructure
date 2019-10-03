// Creates merkle tree from the transactions as leaf nodes.

package merkle_tree

func CreateMerkleTree(transactions []string) [][]string {
	var listOfList [][]string = make([][]string, 0)
	listOfList = append(listOfList, transactions)
	i := 0
	for len(listOfList[i]) != 1{
		if len(listOfList[i]) % 2 == 1 {
			listOfList[i] = append(listOfList[i], listOfList[i][len(listOfList[i])-1])
		}
		var list []string = make([]string, 0)
		j := 0
		for j < len(listOfList[i]) {
			hash := Hash(listOfList[i][j] + listOfList[i][j+1])
			list = append(list, hash)
			j += 2
		}
		listOfList = append(listOfList, list)
		i += 1
	}
	length := len(listOfList)
	for i = 0 ; i < length / 2 ; i++ {
		listOfList[i], listOfList[length - i - 1] = listOfList[length - i -1], listOfList[i]
	}
	return listOfList
}