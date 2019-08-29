package main

import (
	"bytes"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

// 4.引入区块链
//使用数据库代替数组
type BlockChain struct {
	db *bolt.DB
	tail []byte  //存储最后一个区块的哈希
	//blocks []*Block
}
// bolt数据库名称
const blockChainDb = "blockChain.db"
// bucket名称
const blockBucket = "blockBucket"
// 5.创建区块链
func NewBlockChain(address string) *BlockChain{
	// 创建创世块,并作为第一个区块放入到区块链中
	//genesisBlock:=GenesisBlock()
	//return &BlockChain{
	//	blocks:[]*Block{genesisBlock},
	//}
	//最后一个区块的哈希， 从数据库中读出来的
	var lastHash []byte

	//1. 打开数据库
	db, err := bolt.Open(blockChainDb, 0600, nil)
	//defer db.Close()

	if err != nil {
		log.Panic("打开数据库失败！")
	}

	//将要操作数据库（改写）
	db.Update(func(tx *bolt.Tx) error {
		//2. 找到抽屉bucket(如果没有，就创建）
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			//没有抽屉，我们需要创建
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic("创建bucket(b1)失败")
			}

			//创建一个创世块，并作为第一个区块添加到区块链中
			genesisBlock := GenesisBlock(address)
			fmt.Printf("诞生创世块genesisBlock :%s\n", genesisBlock)
			//3. 写数据
			//hash作为key， block的字节流作为value，尚未实现
			bucket.Put(genesisBlock.CurBlockHash, genesisBlock.Serialize())
			bucket.Put([]byte("LastHashKey"), genesisBlock.CurBlockHash)
			lastHash = genesisBlock.CurBlockHash

			////这是为了读数据测试，马上删掉,套路!
			//blockBytes := bucket.Get(genesisBlock.Hash)
			//block := Deserialize(blockBytes)
			//fmt.Printf("block info : %s\n", block)

		} else {
			lastHash = bucket.Get([]byte("LastHashKey"))
		}

		return nil
	})

	return &BlockChain{db, lastHash}
}
// 6.创建一个创世块 会产生挖矿交易
func GenesisBlock(address string)*Block{
	coinbase:=NewCoinbaseTX(address,"我是创世块") // 地址 创世块挖矿交易信息【一般传矿池名称】
	return NewBlock([]*Transaction{coinbase},[]byte{})
	//return NewBlock("我是GOBTC的第一个区块!",[]byte{})

}
// 7.区块链中创建区块
func (blockChain *BlockChain)AddBlock(data string){
	//// 1.获取区块链中最后一个区块
	//lastBlcok:=blockChain.blocks[len(blockChain.blocks)-1]
	//// 2.取它的HASH作为最新区块的前区块HASH
	//preBlockhash:=lastBlcok.CurBlockHash
	//// 3.创建新的区块
	//block:=NewBlock(data,preBlockhash)
	//// 4.将创建的最新区块追加进区块链中
	//blockChain.blocks=append(blockChain.blocks,block)
	//如何获取前区块的哈希呢？？
	db := blockChain.db //区块链数据库
	lastHash := blockChain.tail //最后一个区块的哈希

	db.Update(func(tx *bolt.Tx) error {

		//完成数据添加
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("bucket 不应该为空，请检查!")
		}

		//a. 创建新的区块 TODO
		block := NewBlock([]*Transaction{}, lastHash)

		//b. 添加到区块链db中
		//hash作为key， block的字节流作为value，尚未实现
		bucket.Put(block.CurBlockHash, block.Serialize())
		bucket.Put([]byte("LastHashKey"), block.CurBlockHash)

		//c. 更新一下内存中的区块链，指的是把最后的小尾巴tail更新一下
		blockChain.tail = block.CurBlockHash

		return nil
	})
}
func (bc *BlockChain) Printchain() {

	blockHeight := 0
	bc.db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("blockBucket"))

		//从第一个key-> value 进行遍历，到最后一个固定的key时直接返回
		b.ForEach(func(k, v []byte) error {
			if bytes.Equal(k, []byte("LastHashKey")) {
				return nil
			}

			block := Deserialize(v)
			//fmt.Printf("key=%x, value=%s\n", k, v)
			fmt.Printf("=============== 区块高度: %d ==============\n", blockHeight)
			blockHeight++
			fmt.Printf("版本号: %d\n", block.Version)
			fmt.Printf("前区块哈希值: %x\n", block.PreBlockHash)
			fmt.Printf("梅克尔根: %x\n", block.MerkelRoot)
			fmt.Printf("时间戳: %d\n", block.TimeStamp)
			fmt.Printf("难度值(随便写的）: %d\n", block.Difficulty)
			fmt.Printf("随机数 : %d\n", block.Nonce)
			fmt.Printf("当前区块哈希值: %x\n", block.CurBlockHash)
			fmt.Printf("区块数据 :%s\n", block.Transactions[0].TXInputs[0].Sig)
			return nil
		})
		return nil
	})
}
//  找到指定地址的所有的UTXO
func (blockChain *BlockChain) FindUTXOs(address string) []TXOutput{
	// 保存未消耗的outpututxo
     var UTXO []TXOutput
     // 保存已经消耗的outpututxo
     var spendedUtxo map[string][]int64
     // TODO
     // 1.遍历区块
     // 2.遍历交易
     // 3.遍历output，找到和自己相关的outxo(在添加oututxo之前判断是否消耗过)
     // 4.遍历input，找到自己话费过的utxo的集合（将自己花费过的标记出来）
	 // 1.遍历区块
     it:=blockChain.NewIterator()
     for{
     	block:=it.Next()
		 // 2.遍历交易
		for _,tx:=range block.Transactions{
			fmt.Printf("遍历区块中的交易ID: x%\n",tx.TXID)
			// 遍历交易中的oututxo 通过地址进行判断属于该地址的
			// 3.遍历output 找到和自己相关的utxo(在添加Output之前检查一下是否消耗过)
			OUTPUT: //循环跳出标记，加快循环
			for i,txoutput:=range tx.TXOutputs{
				fmt.Printf("current txoutput index :%d\n",i)
				// 在这里做一个过滤将所有消耗过的txOutput和当前即将添加的txoutput对比，如果相同，则跳过，否则添加
				// 判断标准：如果当前交易的id存在已经消耗的txoutputmap数据集合中，那么说明这个交易里面有消耗过的Output
				if spendedUtxo[string(tx.TXID)]!=nil{
                      for j:=range spendedUtxo[string(tx.TXID)]{
                       	if i==j{
                      		// 当前txoutput已经被消耗过了不添加
                          continue OUTPUT
						}
					  }
				}
				// 否则如果属于该指定地址那么添加进 UTXO数组中
				if txoutput.PubkeyHash==address{
					UTXO = append(UTXO, txoutput)
					fmt.Printf("没有消耗添加进该地址用户:%f\n",UTXO[0].Value)
				}else{
                    fmt.Printf("没有消耗但不属于该地址用户")
				}
			}
			// 判断该交易是否为挖矿交易 如果不是挖矿交易才进行遍历
			if !tx.IsCoinbaseTx(){
				// 遍历区块中与指定地址有关且已经被消耗的交易 inpututxo ；将已经消耗的utxo保存到map中
				for _,txinputut:=range tx.TXInputs{
					// 如果签名与地址相同 ----一般而言签名就是地址（这里需要确定）；如果相同说明该output已经被消耗过加入到spendedUtxo中 ?
					if txinputut.Sig==address{
						//indexArry:= spendedUtxo[string(inpututxo.TXid)]
						//indexArry = append(indexArry,inpututxo.Index)
						spendedUtxo[string(txinputut.TXid)]=append(spendedUtxo[string(txinputut.TXid)],txinputut.Index)
					}
				}
			}else{
				fmt.Printf("这是conbasiecoin挖矿交易不做input遍历!")
			}

		 }
	 }
     return UTXO
}
// 找到合理的utxos
func (blockChain *BlockChain) finNeedUTXOs(from string,amount float64)(map[string][]int64,float64){
   // 找到合理的utxos集合
	var utxos map[string][]int64
  // 找到的utxos里面包含钱的总数
  var calculation float64
  return utxos,calculation
}