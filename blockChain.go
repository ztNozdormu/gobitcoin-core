package main

// 4.引入区块链
type BlockChain struct {
	blocks []*Block
}
// 5.创建区块链
func NewBlockChain() *BlockChain{
	// 创建创世块,并作为第一个区块放入到区块链中
	genesisBlock:=GenesisBlock()
	return &BlockChain{
		blocks:[]*Block{genesisBlock},
	}
}
// 6.创建一个创世块
func GenesisBlock()*Block{
	return NewBlock("我是GOBTC的第一个区块!",[]byte{})
}
// 7.区块链中创建区块
func (blockChain *BlockChain)AddBlock(data string){
	// 1.获取区块链中最后一个区块
	lastBlcok:=blockChain.blocks[len(blockChain.blocks)-1]
	// 2.取它的HASH作为最新区块的前区块HASH
	preBlockhash:=lastBlcok.CurBlockHash
	// 3.创建新的区块
	block:=NewBlock(data,preBlockhash)
	// 4.将创建的最新区块追加进区块链中
	blockChain.blocks=append(blockChain.blocks,block)
}