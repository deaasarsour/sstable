# sstable: Simple LSM Tree Database

sstable is a simple implementation of a Log-Structured Merge-Tree (LSM Tree) database in GoLang. It provides an efficient way to store and retrieve key-value pairs with high write throughput and fast read operations.

## Features

- **LSM Tree Structure**: Utilizes the LSM Tree architecture for efficient storage and retrieval of key-value pairs.
- **SSTable Storage**: Data is stored in Sorted String Tables (SSTables) which provide efficient read operations.
- **Basic API**: Provides basic CRUD operations for interacting with the database.
- **Simple Implementation**: Written in GoLang, sstable aims to be easy to understand and extend.

## Usage

### Manual Installation and Running

To manually install and run sstable in your GoLang project, follow these steps:

1. **Install GoLang**: Ensure you have GoLang installed on your system.

2. **Clone the Repository**: Clone the sstable repository to your local machine:

3. **Run the Code**: Navigate to the cloned repository and run the GoLang code: 

### Docker Image (Coming Soon)

In the future, we plan to provide a Docker image for easy deployment. Stay tuned for updates!

## Using the SSTable Client

The sstable package comes with a command-line client for interacting with the database. The client supports two commands:

- **read**: Retrieve the value associated with a specified key.

- **write**: Store a key-value pair in the database.
