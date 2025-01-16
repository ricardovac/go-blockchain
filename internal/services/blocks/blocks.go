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
	defaultDifficulty := opts.Config.DefaultDifficulty

	s := &Service{
		config:     opts.Config,
		logger:     opts.Logger,
		blockchain: make([]Block, 0),
	}

	// genesis block
	genesisBlock := Block{
		Index:      0,
		Timestamp:  time.Now().Format(time.RFC3339Nano),
		Data:       "Genesis Block",
		Difficulty: defaultDifficulty,
		PrevHash:   "",
	}
	mineBlock(&genesisBlock, 2)
	s.blockchain = append(s.blockchain, genesisBlock)
	s.logger.Info("Genesis block created", zap.Any("block", genesisBlock))

	return s
}

type Service struct {
	config     config.Config
	logger     *zap.Logger
	blockchain []Block
}

func (s *Service) generateBlock(oldBlock Block, data string, difficulty int) (Block, error) {
	if data == "" {
		return Block{}, fmt.Errorf("data must be greater than 0")
	}

	newBlock := Block{
		Index:      oldBlock.Index + 1,
		Timestamp:  time.Now().Format(time.RFC3339Nano),
		Data:       data,
		PrevHash:   oldBlock.Hash,
		Difficulty: difficulty,
	}

	mineBlock(&newBlock, difficulty)

	return newBlock, nil
}

func (s *Service) calculateAverageMineTime() float64 {
	if len(s.blockchain) < 2 {
		return 0
	}

	var total float64
	for i := 1; i < len(s.blockchain); i++ {
		total += s.blockchain[i].Stats.MiningTime
	}

	return total / float64(len(s.blockchain)-1)
}

func (s *Service) adjustDifficulty(difficulty int) int {
	if len(s.blockchain) > 0 {
		lastBlock := s.blockchain[len(s.blockchain)-1]
		if lastBlock.Stats.MiningTime < 1.0 {
			return min(difficulty+1, 6)
		} else if lastBlock.Stats.MiningTime > 10.0 {
			return max(1, difficulty-1)
		}
	}
	return difficulty
}

func mineBlock(block *Block, difficulty int) {
	start := time.Now()
	attempts := 0
	prefix := strings.Repeat("0", difficulty)

	if block.Index == 0 {
		block.Hash = calculateHash(*block)
		block.Stats = BlockStats{
			MinedAt:    time.Now(),
			MiningTime: 0,
			Attempts:   1,
		}
		return
	}

	for {
		block.Nonce++
		attempts++
		block.Hash = calculateHash(*block)
		if strings.HasPrefix(block.Hash, prefix) {
			break
		}
		fmt.Printf("\r%x", block.Hash)
		if attempts > 1000000 {
			block.Difficulty--
			prefix = strings.Repeat("0", block.Difficulty)
		}
	}

	block.Stats = BlockStats{
		MinedAt:    time.Now(),
		MiningTime: time.Since(start).Seconds(),
		Attempts:   attempts,
	}

	block.Stats.Process = struct {
		CurrentHash  string  `json:"currentHash"`
		HashesPerSec float64 `json:"hashesPerSec"`
		Target       string  `json:"target"`
	}{
		CurrentHash:  block.Hash,
		HashesPerSec: float64(attempts) / time.Since(start).Seconds(),
		Target:       strings.Repeat("0", difficulty),
	}
}

func calculateHash(block Block) string {
	record := fmt.Sprintf("%d%s%s%s%d%d",
		block.Index,
		block.Timestamp,
		block.Data,
		block.PrevHash,
		block.Nonce,
		block.Difficulty,
	)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
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

	prefix := strings.Repeat("0", newBlock.Difficulty)
	if !strings.HasPrefix(newBlock.Hash, prefix) {
		return false
	}

	if newBlock.Timestamp <= oldBlock.Timestamp {
		return false
	}

	if newBlock.Difficulty < 1 || newBlock.Difficulty > 4 {
		return false
	}

	return true
}
