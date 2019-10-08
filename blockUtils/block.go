package blockchain

import(
	"bytes"
	"crypto/sha256"
	"time"
	"strconv"
)

type Block struct{
	BlockSize int
	BlockHeader
	Hash string
	//data can be coinbase transactions
	//images, passwords, certificates or anything in future
	Data string
	Height int
	TransactionCounter int
	Transactions []*Transaction

	//for difficulty refer this in future: 
	//https://en.bitcoin.it/wiki/Difficulty
}

type BlockHeader struct{
	Version float32
	PrevHash string
	MerkleRoot string
	Timestamp string
	Nonce int
	DifficultyTarget int
}

type Transaction struct{

}
//blockchain is an array of pointers to block. as we move further, we will make 
//the blockchain structure more perfect
type BlockChain struct{
	//making Blocks public as we need a chain in the main function
	Blocks []*Block 
}
// deriving the hash of block using it's data and prev hash
func (b *Block)DeriveHash() {
	timestamp := b.BlockHeader.Timestamp
	DataAndHash := [][]byte {[]byte(b.Data),[]byte(timestamp),[]byte(b.BlockHeader.PrevHash)}
	info := bytes.Join(DataAndHash,[]byte{})
	//will use a better algorithm as we move ahead in future.
	hash := sha256.Sum256(info)
	//sha can not be applied to strings but byte arrays..check the reason
	b.Hash = string(hash[:])
}

func CreateBlock(data string, prevHash string) *Block {
	// block := &Block{strconv.FormatInt(time.Now().Unix(),10),"", prevHash, data}
	block := &Block{
		BlockSize: 4,
		BlockHeader: BlockHeader{
				Version: 1.0,
				PrevHash: prevHash,
				MerkleRoot: "merkle root",
				Timestamp: strconv.FormatInt(time.Now().Unix(),10),
				Nonce: 1,
				DifficultyTarget: 1,
			},
		Data: data, 
		Height: 0,
		Hash: "",
		TransactionCounter: 1,
		Transactions: nil,
	}
	block.DeriveHash();
	return block
}

//to add a block to blockchain
func (chain *BlockChain) AddBlock(data string){
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new);
} 


//for the very first block which does not point to any previous block 
//sending data of genesis block as boice :)
func Genesis() *Block{
	return CreateBlock("Boice", "")
}

func InitBlockChain() *BlockChain{
	//creating an array of blocks by calling genesis function and returning 
	//the reference to this blockchain
	return &BlockChain{[]*Block{Genesis()}}
}
