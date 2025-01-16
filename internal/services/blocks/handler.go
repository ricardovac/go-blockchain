package blocks

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Data       string `json:"data"`
	Difficulty int    `json:"difficulty"`
}

func (s *Service) HandleWriteBlock(c *gin.Context) {
	var m Message
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "invalid request",
			Code:    http.StatusBadRequest,
			Details: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	done := make(chan struct{})
	var newBlock Block
	var genErr error

	go func() {
		previousBlock := s.blockchain[len(s.blockchain)-1]
		newBlock, genErr = s.generateBlock(previousBlock, m.Data, m.Difficulty)
		close(done)
	}()

	select {
	case <-ctx.Done():
		c.JSON(http.StatusRequestTimeout, ErrorResponse{
			Error:   "timeout",
			Code:    http.StatusRequestTimeout,
			Details: "request took too long",
		})
		return
	case <-done:
		if genErr != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Error:   "internal error",
				Code:    http.StatusInternalServerError,
				Details: genErr.Error(),
			})
			return
		}
	}

	if isBlockValid(newBlock, s.blockchain[len(s.blockchain)-1]) {
		s.blockchain = append(s.blockchain, newBlock)
		m.Difficulty = s.adjustDifficulty(m.Difficulty)
		c.JSON(http.StatusCreated, BlockResponse{
			Block:       newBlock,
			ChainLength: len(s.blockchain),
			Difficulty:  m.Difficulty,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrorResponse{
			Error:   "invalid block",
			Code:    http.StatusBadRequest,
			Details: "block is not valid",
		}})
	}
}

func (s *Service) HandleGetBlocks(c *gin.Context) {
	avgMineTime := s.calculateAverageMineTime()

	c.JSON(http.StatusOK, ChainResponse{
		Blocks:     s.blockchain,
		Difficulty: s.blockchain[len(s.blockchain)-1].Difficulty,
		Stats: struct {
			TotalBlocks int     `json:"totalBlocks"`
			AvgMineTime float64 `json:"avgMineTime"`
		}{
			TotalBlocks: len(s.blockchain),
			AvgMineTime: avgMineTime,
		},
	})
}

func (s *Service) HandleVerifyChain(c *gin.Context) {
	isValid, err := s.isChainValid()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "internal error",
			Code:    http.StatusInternalServerError,
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"valid": isValid})
}
