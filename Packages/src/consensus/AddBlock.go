package consensus

import ("block_structure";
		)

func  AddBlock(chain * block_structure.Blockchain, data string){
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new);
}