package consensus

import("crypto/sha256";
		"fmt";
		"strings";
	)

func (pow *ProofOfWork) Validate() bool {
	data := pow.InitData(pow.Block.Header.Nonce)
	hash := sha256.Sum256(data)
	hexHash := fmt.Sprintf("%x", hash)
	return strings.Compare(hexHash, pow.Target) == -1
}