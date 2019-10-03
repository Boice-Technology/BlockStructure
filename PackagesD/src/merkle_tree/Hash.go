// Generates SHA256(SHA256()) hash of the input string.

package merkle_tree

import elliptic_curve_key_generation "github.com/Boice-Technology/DistributedLedgerCryptographyPieces/Packages/src/elliptic_curve_key_generation"

func Hash(inputString string) string {
	intermediateHash := elliptic_curve_key_generation.SHA256(inputString)
	finalHash := elliptic_curve_key_generation.SHA256(intermediateHash)
	return finalHash
}