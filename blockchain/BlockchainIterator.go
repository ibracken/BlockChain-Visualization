package blockchain

import (
	"log"

	"github.com/boltdb/bolt"
)

type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.Db}

	return bci
}

func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})
	if err != nil {
		// Handle the error, e.g., log it, return an empty byte slice, or panic
		log.Panic(err) // or handle as appropriate
	}

	i.currentHash = block.PrevBlockHash

	return block
}
