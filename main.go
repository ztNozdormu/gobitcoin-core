package main

func main(){
	/**
	 *  bolt读写数据测试
	 */
	//boltutil.WriteBolt([]byte("lastBlockHashKey1"),[]byte("我是最后一个区块的哈希值55555"),"boltDemo1.db","boltBucketDemo")
	//result:= boltutil.ReadBolt([]byte("lastBlockHashKey1"),"boltDemo1.db","boltBucketDemo")
	//fmt.Printf("读取bolt数据库,bucketName为【boltBucketDemo】,key为【lastBlockHashKey】的值为: %s\n:",result)
	////block:=NewBlock("我在博学谷挖到一个BTC",[]byte{})
	//fmt.Printf("前一个区块的数据: %x\n",block.PreBlockHash)
	//fmt.Printf("当前区块的数据: %x\n",block.CurBlockHash)
	//fmt.Printf("区块交易的数据:%s\n",block.Data)
	//fmt.Println("区块交易的数据:",string(block.Data))
	//fmt.Println("hello BTC")
	//blockChain:=NewBlockChain()
	//blockChain.AddBlock("老王转了50BTC给小红")
	//blockChain.AddBlock("小黑挖到一个新的区块L:wq")
	//for i,block:=range blockChain.blocks{
	//	fmt.Printf("当前区块的高度:%d======\n",i)
	//	fmt.Printf("前一个区块的数据: %x\n",block.PreBlockHash)
	//	fmt.Printf("当前区块的数据: %x\n",block.CurBlockHash)
	//	//fmt.Printf("区块交易的数据:%s\n",block.Data)
	//	fmt.Println("区块交易的数据:",string(block.Data))
	//}
	// 迭代器输出
	//it:=blockChain.NewIterator()
	//
	//for{
	//	// 迭代 游标左移动
	//	block:=it.Next()
	//	fmt.Printf("前一个区块的数据: %x\n",block.PreBlockHash)
	//	fmt.Printf("当前区块的数据: %x\n",block.CurBlockHash)
	//	//fmt.Printf("区块交易的数据:%s\n",block.Data)
	//	fmt.Println("区块交易的数据:",string(block.Data))
	//	if len(block.PreBlockHash)==0{
	//		fmt.Println("区块链迭代完成!!!")
	//		break
	//	}
	//}

	bc := NewBlockChain("区块链交易")
	cli := CLI{bc}
	cli.GetBalance("区块链交易")
	cli.Run()
}