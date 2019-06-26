package main

import (
	"fmt"
	"time"
)

//打印区块链
func main() {
	bc := NewBlockChain()

	//变量区块数据
	time.Sleep(1 * time.Second)

	bc.AddBlock("20190626号BTC暴涨10.5%")
	time.Sleep(1 * time.Second)

	bc.AddBlock("20190627号BTC暴涨12.8%")
	time.Sleep(1 * time.Second)

	for i, block := range bc.Blocks {
		fmt.Printf("\n+++++++++ 当前区块高度: %d ++++++++++\n", i)
		fmt.Printf("Version : %d\n", block.Version)
		fmt.Printf("PrevHash : %x\n", block.PrevHash)
		fmt.Printf("MerkleRoot : %x\n", block.MerkleRoot)
		fmt.Printf("TimeStamp : %d\n", block.TimeStamp)
		fmt.Printf("Bits : %d\n", block.Bits)
		fmt.Printf("Nonce : %d\n", block.Nonce)
		fmt.Printf("Hash : %x\n", block.Hash)
		fmt.Printf("Data : %s\n", block.Data)
	}
}
