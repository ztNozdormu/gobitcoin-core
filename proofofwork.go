package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// 1.定义一个工作量证明结构体
type ProofOfWork struct {
	// 区块
	block *Block
	// 目标值 一个非常大的数，有很丰富的方法API:比较，赋值...
	target *big.Int
}
// 2.创建POW的函数
func NewProofOfWork(block *Block) *ProofOfWork{
	pow:=ProofOfWork{
		block:block,
	}
	// 指定的难度值 ，现在是一个string类型需要用Big.int进行转换
	targetStr:="0000100000000000000000000000000000000000000000000000000000000000"
	// 辅助变量，将难度值转成big.int类型的数据
	tmpInt:=big.Int{}
	// 将难度值赋值给big.int
	tmpInt.SetString(targetStr,16)
	pow.target=&tmpInt

	return &pow
}

// 3.提供计算不断计算HASH的函数 返回区块data 和随机数
func (pow *ProofOfWork) PowRun()([]byte,uint64){
	//1. 拼装数据（区块的数据，还有不断变化的随机数）
	//2. 做哈希运算
	//3. 与pow中的target进行比较
	//a. 找到了，退出返回
	//b. 没找到，继续找，随机数加1
	var nonce uint64
	block:=pow.block
	var Hash [32]byte
	fmt.Println("开始挖矿...")
	for{
		//  拼装数据 有不断变化的随机数
		tmp:=[][]byte{
			Uint64ToBytes(block.Version),
			block.PreBlockHash,
			block.MerkelRoot, // 通过交易数据【二叉树算法】生成梅克尔根
			Uint64ToBytes(block.TimeStamp),
			Uint64ToBytes(block.Difficulty),
			Uint64ToBytes(nonce),
			//block.Data,
		}
		// 将二维切片数组拼接起来返回一个一维的切片数组
		blockInfo:=bytes.Join(tmp,[]byte{})
		// 对拼装的数据进行sha256运算 Sum256(data []byte) [Size]byte
		Hash=sha256.Sum256(blockInfo)
		// 将计算的hash转换为Big.int
		tmpBgInt:=big.Int{}
		// 将当前hash转换为的big.int
		tmpBgInt.SetBytes(Hash[:])
		// 与pow中的目标值进行比较 如果比目标值小那么说明挖矿成功，如果没有挖矿成功则继续挖矿
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		//
		//func (x *Int) Cmp(y *Int) (r int)
		if tmpBgInt.Cmp(pow.target) == -1{
			// 挖到则退出 并进行广播通知 TODO
			fmt.Printf("挖矿成功!!! 挖到的区块是hash:%x, 挖矿次数Nonce: %d\n",Hash,nonce)
			return Hash[:],nonce
		}else{
			// 没有挖到则继续挖，Nonce随机数+1
			nonce++
		}
	}
}
// 4.提供一个校验函数 isValid