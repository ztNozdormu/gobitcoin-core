package main

import (
	"github.com/boltdb/bolt"
	"log"
)

/**
 * 区块链迭代器
 */

// 定义迭代器结构体
type BlockChainIterator struct {
	db *bolt.DB
	// 游标用于不断索引
	currentHashPointer []byte
}
// 创建迭代器
func (block *BlockChain) NewIterator() *BlockChainIterator{
	return &BlockChainIterator{
		block.db,
		// 最初指向区块链的最后一个区块，随着Next的调用，不断变化
		block.tail,
    }
}
// 迭代器是属于区块链的 next方法是属于迭代器的
/**
 * 1.返回当前的区块
 * 2.指针前移
 */
func (it BlockChainIterator) Next() *Block{
	var block Block
	// 读取数据库中的区块链
	it.db.View(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte("blockBucket"))
		if bucket==nil{
			log.Panic("迭代器遍历时区块链不应该为空")
		}
		blockTmp:=bucket.Get(it.currentHashPointer)
		// 解码操作
		block = Deserialize(blockTmp)
		// 游标左移动
		it.currentHashPointer=block.PreBlockHash
		return nil
	})
	return &block
}