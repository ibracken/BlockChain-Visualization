package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// ProofOfWork represents the data needed to perform a proof-of-work on a block.
type ProofOfWork struct {
	block  *Block   // Pointer to the block that needs proof of work
	target *big.Int // Target value that the block's hash must be LESS THAN OR EQUAL TO
}

const targetBits = 20               // Usually unknown to the miner
const maxNonce = int(^uint(0) >> 1) // Maximum nonce value

// IntToHex converts an int64 to a byte slice representing the hexadecimal value.
func IntToHex(n int64) []byte {
	return []byte(fmt.Sprintf("%x", n))
}

// NewProofOfWork initializes and returns a new ProofOfWork instance
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	// Shift the bits of target to the left by (256 - targetBits) places.
	// targetBits determines the difficulty of the proof of work.
	target.Lsh(target, uint(256-targetBits))
	// Create a new ProofOfWork instance with the given block and the calculated target.
	pow := &ProofOfWork{b, target}
	return pow
}

// prepareData prepares the data to be hashed for the proof-of-work process.
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	// Join multiple pieces of data into a single byte slice.
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,       // Previous block's hash
			pow.block.Data,                // Current block's data
			IntToHex(pow.block.Timestamp), // Block's timestamp converted to hex
			IntToHex(int64(targetBits)),   // Target difficulty converted to hex
			IntToHex(int64(nonce)),        // Nonce converted to hex
		},
		[]byte{}, // Use an empty separator for joining the byte slices
	)
	return data // Return the joined byte slice
}

// Run performs the proof-of-work (mining) on the block.
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int // Variable to store the hash as an integer
	var hash [32]byte   // Variable to store the hash as a byte array
	nonce := 0          // Nonce value starts at 0
	// Print the data of the block being mined
	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	// Loop to find a valid nonce
	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		// Compute the SHA-256 hash of the prepared data
		hash = sha256.Sum256(data)
		// Print the current hash in hexadecimal format (for debugging purposes)
		fmt.Printf("\r%x", hash)

		// Convert the hash byte array to a big integer
		hashInt.SetBytes(hash[:])

		// Compare the hash integer with the target value; smaller than target = success
		if hashInt.Cmp(pow.target) == -1 {
			// If hash is less than the target, proof-of-work is successful
			break
		} else {
			// If hash is not less than the target, increment the nonce and try again
			nonce++
		}
	}

	// Print new lines after the mining process is completed
	fmt.Print("\n\n")

	// Return the nonce that resulted in a valid hash and the hash itself
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
