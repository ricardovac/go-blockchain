package blocks

import "time"

type BlockStats struct {
	MinedAt    time.Time `json:"minedAt"`
	MiningTime float64   `json:"miningTime"` // seconds
	Attempts   int       `json:"attempts"`   // number of hash attempts
	Process    struct {
		CurrentHash  string  `json:"currentHash"`
		HashesPerSec float64 `json:"hashesPerSec"`
		Target       string  `json:"target"`
	} `json:"progress"`
}

type Block struct {
	Index      int        `json:"index"`
	Timestamp  string     `json:"timestamp"`
	BPM        int        `json:"bpm"`
	Hash       string     `json:"hash"`
	PrevHash   string     `json:"prevHash"`
	Nonce      int        `json:"nonce"`
	Difficulty int        `json:"difficulty"`
	Stats      BlockStats `json:"stats"`
}

type BlockResponse struct {
	Block       Block `json:"block"`
	ChainLength int   `json:"chainLength"`
	Difficulty  int   `json:"difficulty"`
}

type ChainResponse struct {
	Blocks     []Block `json:"blocks"`
	Difficulty int     `json:"difficulty"`
	Stats      struct {
		TotalBlocks int     `json:"totalBlocks"`
		AvgMineTime float64 `json:"avgMineTime"`
	} `json:"stats"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Code    int    `json:"code"`
	Details string `json:"details,omitempty"`
}
