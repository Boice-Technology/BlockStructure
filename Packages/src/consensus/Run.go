package consensus

import ("math";
		"crypto/sha256";
		"fmt";
		"strings"
		)

func (pow *ProofOfWork) Run() (int64, string){
	var hash [32]byte
	nonce := int64(0)
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
	return nonce, hexHash
}