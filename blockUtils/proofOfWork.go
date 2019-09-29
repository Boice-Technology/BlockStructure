package blockchain

import(
	"math/big"
)
//PS: CUCKOO CYCLE AND HASHING IN FUTURE IMPLEMENTATIONS. THIS IS HASHCASE
//IMPLEMENTATION

/*
EXPLANATION:
consensus algorithm:
secure our blockchain by forcing the network to do the work to add a 
block in the chain
Blockchains use consensus algorithms to elect a leader
who will decide the contents of the next block.
That leader is also responsible for broadcasting the block to the network,
so that the other peers can verify the validity of its contents.
In POW, miners have to solve some diffiuclt mathematical problems

work must be hard to do
but reverse engineering (proving reverse) should be easy
*/

/*
taken from tensorflow
will read mastering bitcoin too and update accordingly in the next session

1. take the data from the block
2. create a counter (nonce) which starts at 0, and increments upwards
3. create hash of data plus counter
4. check the hash to see if it meets a set of requirements

Requirements:
1. first few bytes must be 0s (gets adjusted/goes up over time i.e more
no. of 0's)

Eg:
for given data of block as "hi" and given nonce as 7607
the hash is xyz402...

what should the value of nonce be so that the hash starts with exactly 4 0's
as this is the definition we have kept for signed documents
this is what miners find out and it's really difficult to. 

mining: try to find nonce value starting from 1, such that the 
first 4 digits of hash are 0's

best resource: https://anders.com/blockchain/blockchain
difficulty is the no. RELATED to the no. of 0's in initial hash
*/

// make sure to slowly increase the difficulty in future to account for 
//increasing no. of miners and increase in computation power in general. 
//to keep time to mine a block as constant and block rate = constant..

const Difficulty = 12

type ProofOfWork struct {
	Block *Block
	//no. representing requirements
	Target *big.Int
}

  func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
  //continuing this in remote

 }



