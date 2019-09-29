package main


import(
	"fmt"
	"github.com/NikitaMasand/golang-blockchain/blockchain"
	
)

func main(){
	chain := blockchain.InitBlockChain();
	
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")


	// _ is just for index of block
	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
    fmt.Printf("Hash: %x\n", block.Timestamp)
	}

}
