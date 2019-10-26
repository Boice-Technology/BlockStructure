package consensus

import ("strconv";
		"bytes";
		)

func (pow *ProofOfWork) InitData(nonce int64) []byte {
	timestamp := pow.Block.Header.Timestamp
	DataAndPrevHash := [][]byte {
		[]byte(pow.Block.Data),
		[]byte(timestamp),
		[]byte(pow.Block.Header.PrevHash),
		[]byte(pow.Block.Header.TargetDifficulty),
		[]byte(strconv.Itoa(int(nonce))),
	}
	info := bytes.Join(DataAndPrevHash,[]byte{})
	return info
}
