package blockchain

import (
	"log"

	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

type Blockchain struct {
	// ensure capitalization of Blocks to make it public
	// Visibility in go is determined by capitalization of identifiers
	tip []byte
	Db  *bolt.DB
}

// Appends a block to the blockchain
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := bc.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})
	if err != nil {
		// Handle the error, e.g., log it, return an empty byte slice, or panic
		log.Panic(err) // or handle as appropriate
	}

	newBlock := NewBlock(data, lastHash)

	err = bc.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			// Handle the error, e.g., log it, return an empty byte slice, or panic
			log.Panic(err) // or handle as appropriate
		}
		err = b.Put([]byte("l"), newBlock.Hash)
		bc.tip = newBlock.Hash

		return nil
	})

}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)

	if err != nil {
		// Handle the error, e.g., log it, return an empty byte slice, or panic
		log.Panic(err) // or handle as appropriate
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				// Handle the error, e.g., log it, return an empty byte slice, or panic
				log.Panic(err) // or handle as appropriate
			}
			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				// Handle the error, e.g., log it, return an empty byte slice, or panic
				log.Panic(err) // or handle as appropriate
			}
			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				// Handle the error, e.g., log it, return an empty byte slice, or panic
				log.Panic(err) // or handle as appropriate
			}
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	bc := Blockchain{tip, db}

	return &bc
}
