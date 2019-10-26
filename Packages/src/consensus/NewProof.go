package consensus

import ("block_structure";
		)	

func NewProof(b *block_structure.Block) *ProofOfWork {
	target := EvaluateFunction(b)
	pow := &ProofOfWork{b, target}
	return pow
}