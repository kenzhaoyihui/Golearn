package main

import (
	"strconv"
	"bytes"
	"crypto/sha256"
	"time"
	"fmt"
	"os"
)

type Block struct {
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
}


func (block *Block) SetHash() {

	//int64 -> []byte
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))

	headers := bytes.Join([][]byte{block.PrevBlockHash, block.Data, timestamp}, []byte{})

	hash := sha256.Sum256(headers)

	block.Hash = hash[:]

}

func NewBlock(data string, prebBlockHash []byte) *Block {
	block := Block{
		Timestamp: time.Now().Unix(),
		Data: []byte(data),
		PrevBlockHash:prebBlockHash,
	}

	block.SetHash()
	return &block
}



type BlockChain struct {
	Blocks []*Block
}


func NewBlockChain() *BlockChain {
	blockChain := BlockChain{
		Blocks:[]*Block{NewGenesisBlock()},
	}
	return &blockChain
}

func (blockChain *BlockChain) AddBlock(data string) {
	prevBlock := blockChain.Blocks[len(blockChain.Blocks)-1]

	newBlock := NewBlock(data, prevBlock.Hash)

	blockChain.Blocks = append(blockChain.Blocks, newBlock)
}

// the first block
func NewGenesisBlock() *Block {
	genesisBlock := Block{
		Data:[]byte("Genesis block"),
		PrevBlockHash:[]byte{},
	}

	return &genesisBlock
}


func main() {
	//block := NewBlock("haha", []byte{})
	//fmt.Printf("block.hash = %x\n", block.Hash)

	bc := NewBlockChain()

	var cmd string
	for {
		fmt.Println("press '1' to add one block")
		fmt.Println("press '2' to show all blocks")
		fmt.Println("press other to exit")
		fmt.Println()

		fmt.Scanf("%s", &cmd)

		switch cmd {
		case "1":
			input := make([]byte, 1024)
			fmt.Println("Please input the data: ")
			os.Stdin.Read(input)
			bc.AddBlock(string(input))
		case "2":
			for i,block := range bc.Blocks {
				fmt.Println("========================")
				fmt.Printf("Num: %d; PreHash: %x; Data: %s; Hash: %x", i, block.PrevBlockHash, block.Data, block.Hash)
				fmt.Println()
			}

		default:
			fmt.Println("Exit!")
			return
		}
	}
}
