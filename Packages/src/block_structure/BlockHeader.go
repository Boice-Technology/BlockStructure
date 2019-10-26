package block_structure

type BlockHeader struct{
	Version float32
	PrevHash string
	MerkleRoot string
	Timestamp string
	Nonce int64
	TargetDifficulty string
}