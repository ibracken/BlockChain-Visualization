package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// 'b' pointer to a Block struct.
func (b *Block) SetHash() {
	// Converts timestamp into cytes
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	// Concatenates all the fields of the block
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	// Computes thes SHA-256 hash of the concatenated headers
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]

}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)

	if err != nil {
		// Handle the error, e.g., log it, return an empty byte slice, or panic
		log.Panic(err) // or handle as appropriate
	}

	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)

	if err != nil {
		// Handle the error, e.g., log it, return an empty byte slice, or panic
		log.Panic(err) // or handle as appropriate
	}

	return &block
}
