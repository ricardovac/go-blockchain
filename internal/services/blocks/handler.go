package blocks

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Message struct {
	BPM int
}

func (s *Service) HandleWriteBlock(c *gin.Context) {
	var m Message
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add timeout context
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	done := make(chan struct{})
	var newBlock Block
	var err error

	go func() {
		previousBlock := s.blockchain[len(s.blockchain)-1]
		newBlock, err = s.generateBlock(previousBlock, m.BPM)
		close(done)
	}()

	select {
	case <-ctx.Done():
		c.JSON(http.StatusRequestTimeout, ErrMiningTimeout)
		return
	case <-done:
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrInvalidBlock)
			return
		}
	}

	if isBlockValid(newBlock, s.blockchain[len(s.blockchain)-1]) {
		s.blockchain = append(s.blockchain, newBlock)
		s.adjustDifficulty()
		c.JSON(http.StatusCreated, gin.H{
			"block":       newBlock,
			"chainLength": len(s.blockchain),
			"difficulty":  s.difficulty,
			"miningStats": newBlock.Stats,
		})
	} else {
		c.JSON(http.StatusBadRequest, ErrInvalidBlock)
	}
}

func (s *Service) HandleGetBlockchain(c *gin.Context) {
	bytes, err := json.MarshalIndent(s.blockchain, "", "  ")
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrSerialize)
		return
	}

	c.JSON(http.StatusOK, string(bytes))
}
