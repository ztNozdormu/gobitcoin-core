package main

import "fmt"

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
	return &block
}
func main(){
	block:=NewBlock("我在博学谷挖到一个BTC",[]byte{})
	fmt.Printf("前一个区块的数据: %x\n",block.PreBlockHash)
	fmt.Printf("当前区块的数据: %x\n",block.CurBlockHash)
	fmt.Printf("区块交易的数据:%s\n",block.Data)
	fmt.Println("区块交易的数据:",string(block.Data))
	fmt.Println("hello BTC")
}