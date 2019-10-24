package blockchain

import(
	"math/big"
	"math"
	"bytes"
	"crypto/sha256"
	"fmt"
)

const Difficulty = 12

type ProofOfWork struct{
	Block *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) InitData(nonce int) []byte {
	timestamp := pow.Block.BlockHeader.Timestamp
	DataAndPrevHash := [][]byte {
		[]byte(pow.Block.Data),
		[]byte(timestamp),
		[]byte(pow.Block.BlockHeader.PrevHash),
		[]byte(fmt.Sprintf("%x", int64(Difficulty))),
		[]byte(fmt.Sprintf("%x", int64(nonce))),
	}
	info := bytes.Join(DataAndPrevHash,[]byte{})

	return info

}

func (pow *ProofOfWork) Run() (int, string){
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64{
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		//converting hash to bigint to compare this with target
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
		
	}
	fmt.Println()
	return nonce, string(hash[:])
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int
	data := pow.InitData(pow.Block.BlockHeader.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1

}