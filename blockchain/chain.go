package blockchain

import (
	"fmt"
	"sync"

	"github.com/Woo-Yong0405/nomadcoin/db"
	"github.com/Woo-Yong0405/nomadcoin/utils"
)

type blockchain struct {
	NewestHash  string `json:"newestHash"`
	Height      int    `json:"height"`
	CurrentDiff int    `json:"currentDifficulty"`
}

const (
	defDiff       int = 2
	diffInterval  int = 5
	blockInterval int = 2
	allowedRange  int = 2
)

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) persist() {
	db.SaveChain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

func (b *blockchain) Blocks() []*Block {
	var blocks []*Block
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

func (b *blockchain) recalcDiff() int {
	allBlocks := b.Blocks()
	newestBlock := allBlocks[0]
	recalcBlock := allBlocks[diffInterval-1]
	actualTime := (newestBlock.Timestamp - recalcBlock.Timestamp) / 60
	expectedTime := diffInterval * blockInterval
	if actualTime <= (expectedTime - allowedRange) {
		return b.CurrentDiff + 1
	} else if actualTime >= (expectedTime + allowedRange) {
		return b.CurrentDiff - 1
	}
	return b.CurrentDiff
}

func (b *blockchain) difficulty() int {
	if b.Height == 0 {
		return defDiff
	} else if b.Height%diffInterval == 0 {
		return b.recalcDiff()
	} else {
		return b.CurrentDiff
	}
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{
				Height: 0,
			}
			persistedBlockchain := db.Checkpoint()
			if persistedBlockchain == nil {
				b.AddBlock("Genesis")
			} else {
				b.restore(persistedBlockchain)
			}
		})
	}
	fmt.Println(b.NewestHash)
	return b
}
