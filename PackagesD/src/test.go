package main

import ("fmt";
		"merkle_tree";
		)

func main(){
	var transactions []string = make([]string, 0)
	transactions = append(transactions,"ABDC","ABCD","DCBE")
	fmt.Println(merkle_tree.CreateTransactionBranch(transactions, "ABDC"))
}