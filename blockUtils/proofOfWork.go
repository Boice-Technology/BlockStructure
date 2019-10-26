package blockchain

import(
	"math/big"
	"math"
	"bytes"
	"crypto/sha256"
	"fmt"
	"strings"
	"strconv"
	"../PackagesD/src/string_padding"
)

type ProofOfWork struct{
	Block *Block
	Target string
}

func NewProof(b *Block) *ProofOfWork {
	target := evaluateFunc(b)

	pow := &ProofOfWork{b, target}

	return pow
}

func evaluateFunc(b *Block) string {

	targetBits := b.BlockHeader.TargetDifficulty
	exponentHex := targetBits[2:4] 
	coefficientHex := targetBits[4:] 
	eight,_ := new(big.Int).SetString("8",10)
	three,_ := new(big.Int).SetString("3",10)
	exp_big,_ := new(big.Int).SetString(exponentHex,16)
	subtract := new(big.Int).Sub(exp_big,three)
	power := new(big.Int).Mul(subtract,eight) 
	coefficientInt,_ := new(big.Int).SetString(coefficientHex,16)
	right_exp,_ := new(big.Int).SetString("1",10)
	x,_ := new(big.Int).SetString("2",10)
	right_exp.Exp(x,power,nil)
	targetInt := new(big.Int).Mul(coefficientInt,right_exp)
	target := fmt.Sprintf("%x",targetInt)
	target = string_padding.LeftPadHexadecimalString(target,32)
	return target
}

func (pow *ProofOfWork) InitData(nonce int) []byte {
	timestamp := pow.Block.BlockHeader.Timestamp
	DataAndPrevHash := [][]byte {
		[]byte(pow.Block.Data),
		[]byte(timestamp),
		[]byte(pow.Block.BlockHeader.PrevHash),
		[]byte(pow.Block.BlockHeader.TargetDifficulty),
		[]byte(strconv.Itoa(nonce)),
	}
	info := bytes.Join(DataAndPrevHash,[]byte{})

	return info

}

func (pow *ProofOfWork) Run() (int, string){
	var hash [32]byte
	nonce := 0
	hexHash := ""
	for nonce < math.MaxInt64{
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)
		hexHash = fmt.Sprintf("%x", hash)
		if strings.Compare(hexHash,pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()
	return nonce, hexHash
}

func (pow *ProofOfWork) Validate() bool {
	data := pow.InitData(pow.Block.BlockHeader.Nonce)
	hash := sha256.Sum256(data)
	hexHash := fmt.Sprintf("%x", hash)
	return strings.Compare(hexHash, pow.Target) == -1
}

