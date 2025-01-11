# Go Blockchain Simulator

A simple blockchain implementation in Go that demonstrates the basic concepts of blockchain technology including mining, proof of work, and difficulty adjustment.

## Features

- Block mining with adjustable difficulty
- Proof of work implementation
- Dynamic difficulty adjustment
- REST API endpoints
- CORS support for frontend integration

## Prerequisites

- Go 1.20 or higher
- Make (optional, for using Makefile commands)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/ricardovac/go-blockchain.git
cd go-blockchain
```

2. Create and configure environment variables:

```bash
cp .env.example .env
```

Configure the following variables in
PORT=8080 # API port
DEFAULT_DIFFICULTY=2 # Initial mining difficulty (1-6)

3. Install dependencies:

```bash
go mod download
```

Running the Application

```bash
go run cmd/main.go
```

## API Endpoints

### Create a new block

```http
POST /blocks?difficulty=2
Content-Type: application/json

{
    "BPM": 100
}
```

### Get Blockchain

```http
GET /blocks
```
