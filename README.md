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
- Node.js 14 or higher
- pnpm

## Start

- Install dependencies

```bash
cd apps/web
pnpm install
```

```bash
cd apps/api
go mod download
```

- `cp .env.example .env` and fill required envs

- Run api

```bash
go run apps/api/cmd/main.go
```

- Run frontend

```bash
cd apps/web
pnpm dev
```
