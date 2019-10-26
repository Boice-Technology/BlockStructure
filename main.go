package main


import(
	"fmt"
	"./blockUtils"
	"strconv"
)

func main(){
	chain := blockchain.InitBlockChain();
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")
	for _, block := range chain.Blocks {

		fmt.Printf("Previous Hash: %s\n", block.BlockHeader.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Success with nonce: %d\n", block.BlockHeader.Nonce)
		pow := blockchain.NewProof(block)
		fmt.Printf("Target: %s\n", pow.Target)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
