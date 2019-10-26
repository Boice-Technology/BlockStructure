package block_structure

type Block struct{
	BlockSize int
	Header BlockHeader
	Hash string
	Data string
	Height int
	TransactionCounter int
	Transactions []*Transaction
}