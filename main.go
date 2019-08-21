package main

import (
	"crypto/sha256"
	"fmt"
)

// 0 定义区块结构体
type Block struct {
	// 前一区块哈希
	PreBlockHash []byte
	// 当前区块哈希
	CurBlockHash []byte
	// 交易数据
	Data []byte
}
// 1 创建区块
func NewBlock(data string,preBlockHash []byte) *Block{
	block:=Block{
		PreBlockHash:preBlockHash,
		CurBlockHash:[]byte{}, // 后面通过具体的算法逻辑进行计算 TODO
		Data:[]byte(data),
	}
	// 通过具体的算法逻辑进行计算当前区块HASH
	block.SetHash()
	return &block
}
// 3.生成HASH
func (block *Block)SetHash(){
  // 1.拼装数据
  blockInfo:=append(block.PreBlockHash,block.Data...)
  // 2.shua256
  curBlockHash:= sha256.Sum256(blockInfo)
  block.CurBlockHash=curBlockHash[:]
}
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
func main(){
	//block:=NewBlock("我在博学谷挖到一个BTC",[]byte{})
	//fmt.Printf("前一个区块的数据: %x\n",block.PreBlockHash)
	//fmt.Printf("当前区块的数据: %x\n",block.CurBlockHash)
	//fmt.Printf("区块交易的数据:%s\n",block.Data)
	//fmt.Println("区块交易的数据:",string(block.Data))
	//fmt.Println("hello BTC")
	blockChain:=NewBlockChain()
	blockChain.AddBlock("老王转了50BTC给小红")
	blockChain.AddBlock("小黑挖到一个新的区块L:wq")
	for i,block:=range blockChain.blocks{
		fmt.Printf("当前区块的高度:%d======\n",i)
		fmt.Printf("前一个区块的数据: %x\n",block.PreBlockHash)
		fmt.Printf("当前区块的数据: %x\n",block.CurBlockHash)
		//fmt.Printf("区块交易的数据:%s\n",block.Data)
		fmt.Println("区块交易的数据:",string(block.Data))
	}

}