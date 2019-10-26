package consensus

import ("block_structure";
		"strconv";
		"time";
	)

func CreateBlock(data string, prevHash string) * block_structure.Block {
	block := &block_structure.Block{
		BlockSize: 4,
		Header: block_structure.BlockHeader{
				Version: 1.0,
				PrevHash: prevHash,
				MerkleRoot: "merkle root",
				Timestamp: strconv.FormatInt(time.Now().Unix(),10),
				Nonce: 0,
				TargetDifficulty: "0x2003a30c",
			},
		Data: data, 
		Height: 0,
		Hash: "",
		TransactionCounter: 1,
		Transactions: nil,
	}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Header.Nonce = nonce
	block.Hash = hash
	return block
}
