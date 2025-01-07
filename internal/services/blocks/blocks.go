package blocks

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/ricardovac/go-blockchain/internal/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Opts struct {
	fx.In

	Config config.Config
	Logger *zap.Logger
}

func New(opts Opts) *Service {
	s := &Service{
		config:     opts.Config,
		logger:     opts.Logger,
		blockchain: make([]Block, 0),
		difficulty: 4,
	}

	// genesis block
	t := time.Now()
	genesisBlock := Block{
		Index:      0,
		Timestamp:  t.String(),
		BPM:        0,
		Difficulty: s.difficulty,
		PrevHash:   "", // Empty for genesis
	}
	mineBlock(&genesisBlock, s.difficulty)
	s.blockchain = append(s.blockchain, genesisBlock)
	s.logger.Info("Genesis block created", zap.Any("block", genesisBlock))

	return s
}

type Service struct {
	config     config.Config
	logger     *zap.Logger
	blockchain []Block
	difficulty int
}

func mineBlock(block *Block, difficulty int) {
	start := time.Now()
	attempts := 0
	prefix := strings.Repeat("0", difficulty)

	// initial values for genesis block
	if block.Index == 0 {
		block.Hash = calculateHash(*block)
		block.Stats = BlockStats{
			MinedAt:    time.Now(),
			MiningTime: 0,
			Attempts:   1,
		}
		return
	}

	// Proof of Work
	for {
		block.Nonce++
		attempts++
		block.Hash = calculateHash(*block)
		if strings.HasPrefix(block.Hash, prefix) {
			break
		}
		fmt.Printf("\r%x", block.Hash)
		if attempts > 1000000 {
			block.Difficulty-- // Reduce difficulty if mining takes too long
			prefix = strings.Repeat("0", block.Difficulty)
		}
	}

	block.Stats = BlockStats{
		MinedAt:    time.Now(),
		MiningTime: time.Since(start).Seconds(),
		Attempts:   attempts,
	}
}

func calculateHash(block Block) string {
	record := fmt.Sprintf("%d%s%d%s%d%d",
		block.Index,
		block.Timestamp,
		block.BPM,
		block.PrevHash,
		block.Nonce,
		block.Difficulty,
	)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func (s *Service) generateBlock(oldBlock Block, BPM int) (Block, error) {
	if BPM < 0 {
		return Block{}, ErrInvalidBPM
	}

	newBlock := Block{
		Index:      oldBlock.Index + 1,
		Timestamp:  time.Now().String(),
		BPM:        BPM,
		PrevHash:   oldBlock.Hash,
		Difficulty: s.difficulty,
	}

	mineBlock(&newBlock, s.difficulty)

	return newBlock, nil
}

func (s *Service) adjustDifficulty() {
	if len(s.blockchain) > 0 {
		lastBlock := s.blockchain[len(s.blockchain)-1]
		// Adjust difficulty based on mining time
		if lastBlock.Stats.MiningTime < 1.0 { // Too fast
			s.difficulty++
		} else if lastBlock.Stats.MiningTime > 10.0 { // Too slow
			s.difficulty = max(1, s.difficulty-1)
		}
	}
}

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}
