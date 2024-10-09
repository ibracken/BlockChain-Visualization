This project is a simple implementation of a blockchain in Go. It demonstrates the core concepts of blockchain, including block creation, proof of work, and chain management, with an easy-to-use command line interface (CLI).

Features

Create and add new blocks to the blockchain

Proof of Work (PoW) mechanism for mining blocks

Persistent storage of blockchain data using BoltDB

Command Line Interface (CLI) for interacting with the blockchain

File Structure

simple_blockchain/
├── blockchain/
│   ├── Block.go                  # Defines the Block struct and related functions
│   ├── Blockchain.go             # Manages the blockchain and storage
│   ├── BlockchainIterator.go     # Allows traversal through the blockchain
│   ├── CLI.go                    # Command Line Interface to interact with the blockchain
│   └── ProofOfWork.go            # Implements the Proof of Work mechanism
└── main.go                       # Entry point for the application

Installation and Setup

Prerequisites

Go (version 1.17 or later)

Steps

Clone the repository:

git clone <repository-url>
cd simple_blockchain

Initialize the Go module:

go mod init simple_blockchain

Install dependencies (if any):

go mod tidy

Running the Application

Usage

The application is run from the command line. It accepts the following commands:

Add a Block:

go run main.go addblock -data "Your Block Data"

This command will add a new block to the blockchain with the specified data.

Print the Blockchain:

go run main.go printchain

This command will print all the blocks in the blockchain, showing their details.

Example

Adding a Block

go run main.go addblock -data "Hello Blockchain!"

This will create a new block containing the message "Hello Blockchain!".

Printing the Blockchain

go run main.go printchain

This will display all blocks in the chain, including their data, hash, and previous hash.

Concepts Covered

Block: A unit in the blockchain that contains data, timestamp, and the hash of the previous block.

Proof of Work (PoW): A consensus mechanism that secures the blockchain by requiring computational work to add a new block.

Persistence: Blockchain data is stored using BoltDB, ensuring it is saved and retrievable.

Project Structure Details

Block.go: Defines the Block struct, which represents each block in the blockchain, and includes methods for creating a new block and computing its hash.

Blockchain.go: Defines the Blockchain struct, which manages the chain of blocks and the database storage using BoltDB.

CLI.go: Provides a command line interface to interact with the blockchain by allowing commands like addblock and printchain.

ProofOfWork.go: Implements proof of work, securing the blockchain by making the creation of new blocks computationally expensive.

Dependencies

BoltDB: A fast key/value store for Go used to persist blockchain data.

Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

License

This project is licensed under the MIT License. See the LICENSE file for more details.

Contact

For questions or support, please reach out to the project maintainer.
