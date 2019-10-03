package merkletree

// Node Structure for Node in merkle tree
type Node struct {
	data          int64
	hash          string
	parentPointer *Node
	childPointer  *Node
}

// This will take list of pointers of node and insert nodes at appropriate position
// For zero index system children of nth parent are always at 2n+1, 2n+2
// Call check_consistency()
func addNode() {

}

// This will take list of pointers of node and delete nodes at appropriate position
// For zero index system children of nth parent are always at 2n+1, 2n+2
// Call check_consistency()
func deleteNode() {

}

// Inconsistencies can be removed by scanning the list for NULL values.
// While adding any of the two conditions can arise
//	1) The node is odd node. In that case duplicate the node.
//  2) The node is to be added is in place of duplicated odd node.
// While deleting this condition can arise
//	1) Along its path to root somewhere odd node are created. Then duplicate its sibling
func checkConsistency() {

}

// This function in an attempt to calculate root hash will calculate all the relevant node hashes too.
// One can also assume root to be parent of a sub-tree
func computerootHash() {

}

// This will check for if the tree has been modified by checking root hash with the known, stored root hash.
func checkRootHash() {

}
