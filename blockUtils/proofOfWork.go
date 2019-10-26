package blockchain

import(
	"math/big"
	"math"
	"bytes"
	"crypto/sha256"
	"fmt"
	"strings"
	"strconv"
)

// const Difficulty = 12

type ProofOfWork struct{
	Block *Block
	Target string
}

func NewProof(b *Block) *ProofOfWork {
	// target := big.NewInt(1)
	// target.Lsh(target, uint(256-Difficulty))
	target := evaluateFunc(b)

	pow := &ProofOfWork{b, target}

	return pow
}

func evaluateFunc(b *Block) string {
	// difficulty := b.BlockHeader.Difficulty
	// targetbits := math.Log2(float64(difficulty))  //float64
	// //this is wrong!! :(

	// _,targetbitsHex := fmt.Printf("%x\n", math.Float32bits(float32(targetbits)))

	// exponent := targetbitsHex[:2]
	// coefficient := targetbitsHex[2:]
	// coeff_hex, err := hex.DecodeString(coefficient)
	// expo_hex, err := hex.DecodeString(exponent)


	// target := binary.BigEndian.Uint64(coeff_hex)* math.Pow(2, float64(8)*float64((binary.BigEndian.Uint64(expo_hex)-3)))
	// targetHex := fmt.Sprintf("%x", int64(target))
	// fmt.Println(targetHex)
	
	// return targetHex

	targetBits := b.BlockHeader.TargetDifficulty
	// 0x1903a30c
	exponentHex := targetBits[2:4] //19
	// fmt.Println("exponenthex", exponentHex)
	coefficientHex := targetBits[4:] //03130c
	// fmt.Println("coefficientHex", coefficientHex)
	power := 8*(hex2int(exponentHex)-3) //8 * (25-3) = 176`
	// fmt.Println("power", power)
	coefficientInt := hex2int(coefficientHex) //238,348 
	// fmt.Println("coefficientInt", coefficientInt)

	// targetInt := coefficientInt * math.Pow(2,float64(power))
	targetInt := coefficientInt * Pow(2,int64(power))
	//22,829,202,948,393,929,850,749,706,076,701,368,331,072,452,018,388,575,715,328
	fmt.Println("targetInt", targetInt)
	target := hexString(targetInt,false)
	fmt.Println("target ",target)
	return target
}

func Pow(x int, y int64) uint64 {
	// ans := 1
	// for a := 0; a < int(y); a++ {
	// 	ans = ans * x
	// 	fmt.Println("ans in loop ", ans)
	// }
	
	var one = big.NewInt(1)
    start := big.NewInt(0)
	end := big.NewInt(y)
	ans := big.NewInt(1)
    // i must be a new int so that it does not overwrite start
    for i := new(big.Int).Set(start); i.Cmp(end) < 0; i.Add(i, one) {
        ans = Mul(ans,x)
	}
	return ans
}

func hex2int(hexStr string) uint64 {
	// remove 0x suffix if found in the input string
	cleaned := strings.Replace(hexStr, "0x", "", -1)

	// base 16 for hexadecimal
	result, _ := strconv.ParseUint(cleaned, 16, 64)
	return uint64(result)
}

func hexString(i uint64, prefix bool) string {
	i64 := int64(i)

	if prefix {
			return "0x" + strconv.FormatInt(i64, 16) // base 16 for hexadecimal
	} else {
			return strconv.FormatInt(i64, 16) // base 16 for hexadecimal
	}
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
	//var intHash big.Int
	var hash [32]byte

	nonce := 0
	hashstr := ""
	for nonce < math.MaxInt64{
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)
		hexHash, _ := fmt.Printf("%x", hash)
		hashstr = string(hexHash)[:len(string(hexHash))-1]

		fmt.Printf("%s", hashstr)
		if strings.Compare(hashstr,pow.Target) == -1 {
			break
		} else {
			nonce++
		}
		
	}
	fmt.Println()
	return nonce, hashstr
}

func (pow *ProofOfWork) Validate() bool {
	//var intHash big.Int
	data := pow.InitData(pow.Block.BlockHeader.Nonce)

	hash := sha256.Sum256(data)
	hexHash, _ := fmt.Printf("%x", hash)
	hashstr := string(hexHash)[:len(string(hexHash))-1]

	return strings.Compare(hashstr,pow.Target) == -1
}

