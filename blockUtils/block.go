package blockchain

import(
	"time"
	"strconv"
)

type Block struct{
	BlockSize int
	BlockHeader
	Hash string
	Data string
	Height int
	TransactionCounter int
	Transactions []*Transaction
}

type BlockHeader struct{
	Version float32
	PrevHash string
	MerkleRoot string
	Timestamp string
	Nonce int
<<<<<<< HEAD
	TargetDifficulty string
=======
	TargetDifficulty string //hexadecimal
>>>>>>> 0651e95163650c272bd1059909e9ba7bfd0ac28f
}

type Transaction struct{

}

type BlockChain struct{
	Blocks []*Block 
}

func CreateBlock(data string, prevHash string) *Block {
	block := &Block{
		BlockSize: 4,
		BlockHeader: BlockHeader{
				Version: 1.0,
				PrevHash: prevHash,
				MerkleRoot: "merkle root",
				Timestamp: strconv.FormatInt(time.Now().Unix(),10),
				Nonce: 0,
<<<<<<< HEAD
				TargetDifficulty: "0x2003a30c",
=======
				TargetDifficulty: "0x1903a30c",
>>>>>>> 0651e95163650c272bd1059909e9ba7bfd0ac28f
			},
		Data: data, 
		Height: 0,
		Hash: "",
		TransactionCounter: 1,
		Transactions: nil,
	}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.BlockHeader.Nonce = nonce
	block.Hash = hash
	return block
}

func (chain *BlockChain) AddBlock(data string){
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new);
} 

func Genesis() *Block{
	return CreateBlock("Boice", "")
}

func InitBlockChain() *BlockChain{
	return &BlockChain{[]*Block{Genesis()}}
}
