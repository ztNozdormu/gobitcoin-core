package main

import "crypto/sha256"

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
		CurBlockHash:[]byte{},
		Data:[]byte(data),
	}
	// 通过具体的算法逻辑进行计算当前区块HASH
	block.SetCurHash()
	return &block
}
// 3.设置当前区块HASH
func (block *Block)SetCurHash(){
	// 1.拼装数据
	blockInfo:=append(block.PreBlockHash,block.Data...)
	// 2.shua256
	curBlockHash:= sha256.Sum256(blockInfo)
	block.CurBlockHash=curBlockHash[:]
}
