package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

const reward=12.5
// 1. 定义交易结构
type Transaction struct{
	TXID []byte // 交易ID
	TXInputs []TXInput // 交易输入数组
	TXOutputs []TXOutput // 交易输出数组
}
// 定义交易输入
type TXInput struct{
     // 引用的交易ID
     TXid []byte
     // 引用的output的索引值
     Index int64
     // 解锁脚本，我们用地址来模拟
     Sig string
}
// 定义交易输出
type TXOutput struct{
	// 转账金额
	Value float64
	//// 引用的output的索引值
	//Index int64
	// 解锁脚本，我们用地址来模拟
	PubkeyHash string
}
// 设置交易ID
func (tx *Transaction) SetHash(){
	var buffer bytes.Buffer
	encoder:=gob.NewEncoder(&buffer)
	err:=encoder.Encode(tx)
	if err!=nil{
		log.Panic(err)
	}
	data:=buffer.Bytes()
	hash:=sha256.Sum256(data)
	tx.TXID=hash[:]
}
// 判断当前交易是否为挖矿交易
func (tx *Transaction) IsCoinbaseTx() bool{
	    // 1..交易中只有一个Input
		// 2. 交易ID为空
		// 3. 交易索引为-1
		if len(tx.TXInputs)==1&&bytes.Equal(tx.TXInputs[0].TXid,[]byte{})&&tx.TXInputs[0].Index==-1{
          return true
		}
	return false
}
// 2. 提供创建交易的方法(挖矿交易)
func NewCoinbaseTX(address string,data string)*Transaction{
	// 挖矿交易的特点：
	  // 1. 只有一个Input
	  // 2.无需引用交易id
	  // 3.无需引用index
	  // 矿工由于挖矿时无需指定签名，所以这个sig可以由矿工自己自定义传，一般传矿池的名称
	input := TXInput{
		[]byte{},
		-1,
		data,
	}
	output:=TXOutput{
		reward,
		address,
	}
	// 对于挖矿来说只有一个input 和一个output
	tx:=Transaction{[]byte{},[]TXInput{input},[]TXOutput{output}}
	tx.SetHash()
	return &tx
}
// 3. 创建普通交易方法
// 1.找到最合理UTXO集合 map[string]int64
// 2. 将这些UTXO做一转成inputs
// 3. 创建outputs
// 4  如果有零钱，要找零
// from 转给 to
func NewTransaction(from ,to string,amount float64,blockChan *BlockChain)*Transaction{
    // 1.找到合理的utxos集合
	utxos,resValue:=blockChan.finNeedUTXOs(from,amount)
	if resValue<amount{
		fmt.Print("余额不足，无法交易")
		return nil
	}
	var inputs []TXInput
	var outputs []TXOutput
    // 2. 创建交易输入 将这些UTXO转成inputs
    for id,indexArray:=range utxos{
    	for _,i:= range indexArray {
			input:=TXInput{[]byte(id),int64(i),from}
			inputs = append(inputs, input)
		}
	}
    // 收钱方账户收钱
    output:=TXOutput{amount,to}
    outputs = append(outputs, output)
    // 找零
    if resValue>amount{
    	// 转出方账户减钱
		outputs = append(outputs,TXOutput{resValue-amount,from})
	}

    tx:=Transaction{[]byte{},inputs,outputs}
    tx.SetHash()
    return &tx
}
// 4.根据交易调整程序
