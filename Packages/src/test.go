package main

import(
	"fmt"
	"consensus"
	"strconv"
)

func main(){
	chain := consensus.InitBlockChain();
	consensus.AddBlock(chain, "First Block after Genesis")
	consensus.AddBlock(chain, "Second Block after Genesis")
	consensus.AddBlock(chain, "Third Block after Genesis")
	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %s\n", block.Header.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Success with nonce: %d\n", block.Header.Nonce)
		pow := consensus.NewProof(block)
		fmt.Printf("Target: %s\n", pow.Target)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
