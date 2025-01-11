package blocks

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Message struct {
	BPM int
}

func (s *Service) HandleWriteBlock(c *gin.Context) {
	difficulty, err := strconv.Atoi(c.DefaultQuery("difficulty", fmt.Sprint(s.config.DefaultDifficulty)))
	if err != nil || difficulty < 1 || difficulty > 6 {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "invalid difficulty",
			Code:    http.StatusBadRequest,
			Details: "difficulty must be an integer between 1 and 6",
		})
		return
	}

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
		newBlock, genErr = s.generateBlock(previousBlock, m.BPM, difficulty)
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
		difficulty = s.adjustDifficulty(difficulty)
		c.JSON(http.StatusCreated, BlockResponse{
			Block:       newBlock,
			ChainLength: len(s.blockchain),
			Difficulty:  difficulty,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrorResponse{
			Error:   "invalid block",
			Code:    http.StatusBadRequest,
			Details: "block is not valid",
		}})
	}
}

func (s *Service) HandleGetBlockchain(c *gin.Context) {
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
