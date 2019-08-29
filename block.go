package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"
	"log"
	"time"
)

// 0 定义区块结构体
type Block struct {
	// 1.版本号
	Version uint64
    // 2.梅克尔根，这就是一个哈希值，我们先不管，我们后面v4再介绍
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
	// b 交易数据 交易数据数组
	Transactions  []*Transaction


	//Data []byte
}
// 1 创建区块
func NewBlock(Txs []*Transaction,preBlockHash []byte) *Block{
	block:=Block{
		Version:00,
		PreBlockHash:preBlockHash,
		MerkelRoot:[]byte{},
		TimeStamp:uint64(time.Now().Unix()),
		Difficulty:0,
		Nonce: 0,
		CurBlockHash:[]byte{},
		Transactions:Txs,
		//Data:[]byte(data),
	}
	// 通过具体的算法逻辑进行计算当前区块HASH--pow工作量证明算法
//	block.SetCurHash()
    pow:=NewProofOfWork(&block)
    // 不停的通过工作量证明算法得出data 和随机数
	CurBlockHash,Nonce:=pow.PowRun()
    // 根据挖矿结果对区块进行更新
	block.CurBlockHash=CurBlockHash
	block.Nonce=Nonce

	return &block
}

//序列化
func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer

	//- 使用gob进行序列化（编码）得到字节流
	//1. 定义一个编码器
	//2. 使用编码器进行编码
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&block)
	if err != nil {
		log.Panic("编码出错!")
	}

	//fmt.Printf("编码后的小明：%v\n", buffer.Bytes())

	return buffer.Bytes()
}

//反序列化
func Deserialize(data []byte) Block {

	decoder := gob.NewDecoder(bytes.NewReader(data))

	var block Block
	//2. 使用解码器进行解码
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic("解码出错!")
	}

	return block
}

// 3.设置当前区块HASH
func (block *Block)SetCurHash(){
	//var blockInfo []byte
	// 1.拼装数据
	//blockInfo=append(blockInfo,Uint64ToBytes(block.Version)...)
	//blockInfo=append(blockInfo,block.PreBlockHash...)
	//blockInfo=append(blockInfo, block.MerkelRoot...)
	//blockInfo=append(blockInfo, Uint64ToBytes(block.TimeStamp)...)
	//blockInfo=append(blockInfo,Uint64ToBytes(block.Difficulty)...)
	//blockInfo=append(blockInfo,Uint64ToBytes(block.Nonce)...)
	//blockInfo=append(blockInfo, block.Data...)
	// 代码优化 通过二维切片拼接转换为一维切片
	tmp:=[][]byte{
		Uint64ToBytes(block.Version),
		block.PreBlockHash,
		block.MerkelRoot,
		Uint64ToBytes(block.TimeStamp),
		Uint64ToBytes(block.Difficulty),
		Uint64ToBytes(block.Nonce),
		//block.Data,
	}
	blockInfo:=bytes.Join(tmp,[]byte{})
	// 2.shua256
	curBlockHash:= sha256.Sum256(blockInfo)
	block.CurBlockHash=curBlockHash[:]
}
// 模拟梅克尔根的生成，这里只做简单的拼接不用二叉树
func (block *Block) MakeMerkelRoot() []byte{
	return []byte{}
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