package blocks

import "time"

type BlockStats struct {
	MinedAt    time.Time `json:"minedAt"`
	MiningTime float64   `json:"miningTime"` // seconds
	Attempts   int       `json:"attempts"`   // number of hash attempts
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
