package blocks

import (
	"errors"
)

var (
	ErrInvalidBlock  = errors.New("invalid block")
	ErrInvalidChain  = errors.New("invalid chain")
	ErrMiningTimeout = errors.New("mining timeout")
	ErrInvalidBPM    = errors.New("invalid bpm value provided")
	ErrSerialize		 = errors.New("failed to serialize block")	
)
