package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"time"
)

// 0 定义区块结构体
type Block struct {
	// 1.版本号
	Version uint64
    // 2.梅克尔根
    MerkelRoot []byte
	// 3.当前时间戳
	TimeStamp uint64
	// 4.难度值
	Difficulty uint64
	// 5.随机数
	Nonce uint64
	// 6.前一区块哈希
	PreBlockHash []byte
	// a 当前区块哈希 ，正常区块中没有当前区块的哈希，这里为了后面验证方便做了简化
	CurBlockHash []byte
	// b 交易数据
	Data []byte
}
// 1 创建区块
func NewBlock(data string,preBlockHash []byte) *Block{
	block:=Block{
		Version:00,
		PreBlockHash:preBlockHash,
		MerkelRoot:[]byte{},
		TimeStamp:uint64(time.Now().Unix()),
		Difficulty:0,
		Nonce: 0,
		CurBlockHash:[]byte{},
		Data:[]byte(data),
	}
	// 通过具体的算法逻辑进行计算当前区块HASH
	block.SetCurHash()
	return &block
}
// 3.设置当前区块HASH
func (block *Block)SetCurHash(){
	var blockInfo []byte
	// 1.拼装数据
	blockInfo=append(blockInfo,Uint64ToBytes(block.Version)...)
	blockInfo=append(blockInfo,block.PreBlockHash...)
	blockInfo=append(blockInfo, block.MerkelRoot...)
	blockInfo=append(blockInfo, Uint64ToBytes(block.TimeStamp)...)
	blockInfo=append(blockInfo,Uint64ToBytes(block.Difficulty)...)
	blockInfo=append(blockInfo,Uint64ToBytes(block.Nonce)...)
	blockInfo=append(blockInfo, block.Data...)
	// 2.shua256
	curBlockHash:= sha256.Sum256(blockInfo)
	block.CurBlockHash=curBlockHash[:]
}
// uint64转byte数组
func Uint64ToBytes(num uint64)[]byte{
	var buffer bytes.Buffer
	err:=binary.Write(&buffer,binary.BigEndian,num)
	if err!=nil{
		log.Panic(err)
	}
	return buffer.Bytes()
}