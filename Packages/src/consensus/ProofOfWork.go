package consensus

import ("block_structure";
		)

type ProofOfWork struct{
	Block *block_structure.Block
	Target string
}