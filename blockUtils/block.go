package blockchain

import(
	"bytes"
	"crypto/sha256"
	"time"
	"strconv"
)

type Block struct{
	Timestamp int64
	Hash []byte
	PrevHash []byte
	//data can be coinbase transactions
	//images, passwords, certificates or anything in future
	Data []byte 
	//Nonce int after implementing pow
	//add height when doing transactions
	

	//for difficulty refer this in future: 
	//https://en.bitcoin.it/wiki/Difficulty
}

//blockchain is an array of pointers to block. as we move further, we will make 
//the blockchain structure more perfect
type BlockChain struct{
	//making Blocks public as we need a chain in the main function
	Blocks []*Block 
}
// deriving the hash of block using it's data and prev hash
func (b *Block)DeriveHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp,10))
	DataAndHash := [][]byte {b.Data,timestamp,b.PrevHash}
	info := bytes.Join(DataAndHash,[]byte{})
	//will use a better algorithm as we move ahead in future.
	hash := sha256.Sum256(info)
	//sha can not be applied to strings but byte arrays..check the reason
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{time.Now().Unix(),[]byte{}, prevHash, []byte(data)}
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
	return CreateBlock("Boice", []byte{})
}

func InitBlockChain() *BlockChain{
	//creating an array of blocks by calling genesis function and returning 
	//the reference to this blockchain
	return &BlockChain{[]*Block{Genesis()}}
}
