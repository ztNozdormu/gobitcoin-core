package boltutil

import (
	"github.com/boltdb/bolt"
	"log"
)

// 优化 通过初始化方法结合配置新建数据库链接对象
/**
 * 写入数据到bolt中
 */
func WriteBolt(key []byte,value []byte,dbName string,bucketName string){
	//1. 打开数据库
	db, err := bolt.Open(dbName, 0600, nil)
	defer db.Close()

	if err != nil {
		log.Panic("打开数据库失败！")
	}

	//将要操作数据库（改写）
	db.Update(func(tx *bolt.Tx) error {
		//2. 找到抽屉bucket(如果没有，就创建）
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			//没有抽屉，我们需要创建
			bucket, err = tx.CreateBucket([]byte(bucketName))
			if err != nil {
				log.Panic("创建bucket失败")
			}
		}
		//3. 写数据
		bucket.Put(key, value)
		return nil
	})
}

func ReadBolt(key []byte,dbName string,bucketName string)[]byte{
	//1. 打开数据库
	db, err := bolt.Open(dbName, 0600, nil)
	//defer db.Close()
    var result []byte
	if err != nil {
		log.Panic("打开数据库失败！")
	}
	//4. 读数据
	 db.View(func(tx *bolt.Tx) error {
		//1. 找到抽屉，没有的化直接报错退出
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			log.Panic("bucket 不应该为空，请检查!!!!")
		}
		//2. 直接读取数据
		result = bucket.Get(key)
		return nil
	})
	return result
}