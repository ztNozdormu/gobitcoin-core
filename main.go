package main

import (
	"fmt"
)
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