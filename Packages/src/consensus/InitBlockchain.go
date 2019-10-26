package consensus

import ("block_structure";
		)

func InitBlockChain() * block_structure.Blockchain{
	var blockArray [] * block_structure.Block = make([] * block_structure.Block, 0)
	blockArray = append(blockArray, Genesis())
	var blockchain *block_structure.Blockchain = new(block_structure.Blockchain)
	blockchain.Blocks = blockArray
	return blockchain
}