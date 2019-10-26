package block_structure

type Block struct{
	BlockSize int64
	Header BlockHeader
	Hash string
	Data string
	TransactionCounter int64
	Transactions []*Transaction
}