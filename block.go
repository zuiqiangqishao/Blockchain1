package main

import (
	"bytes"
	"encoding/gob"
	"time"
)

//定义一个区块
type Block struct {
	Version       int64  //版本号
	PrevBlockHash []byte //父区块头哈希值
	Hash          []byte //本区块hash值
	MerKelRoot    []byte //Merkel根
	TimeStamp     int64  //时间戳
	Bits          int64  //难度值
	Nonce         int64  //随机值
	//Data          []byte //交易信息
	Transactions []*Transaction //交易信息
}

/**
创建一个新的区块
 */
func NewBlock(txs []*Transaction, prevBlockHash []byte) *Block {
	var block Block;
	block = Block{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		//Hash TODO
		MerKelRoot: []byte{},
		TimeStamp:  time.Now().Unix(),
		Bits:       targetBits,
		Nonce:      0,
		Transactions:       txs}
	//block.SetHash()
	pow := NewProofOfWork(&block)
	nonce, hash := pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	return &block
}

//创建一个初始块
func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{}) //TODO
}

//序列化block
func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(block)
	CheckErr(err, "编码block失败")
	return buffer.Bytes()
}

func Unserialize(data []byte) *Block {
	if len(data) == 0 {
		return nil;
	}
	var block Block;
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	CheckErr(err, "反序列化失败")
	return &block
}

/*
//设置一个区块的hash值
func (block *Block) SetHash() {
	tmp := [][]byte{
		IntToByte(block.Version),
		block.PrevBlockHash,
		block.MerKelRoot,
		IntToByte(block.TimeStamp),
		IntToByte(block.Bits),
		IntToByte(block.Nonce),
		block.Data}
	data := bytes.Join(tmp, []byte{})
	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}
*/
