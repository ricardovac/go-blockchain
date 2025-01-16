package blocks

import (
	"github.com/ricardovac/go-blockchain/internal/config"
	"github.com/ricardovac/go-blockchain/internal/httpserver"
	"github.com/ricardovac/go-blockchain/internal/services/blocks"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Opts struct {
	fx.In

	Config       config.Config
	HttpServer   *httpserver.Server
	BlockService *blocks.Service
	Logger       *zap.Logger
}

func New(opts Opts) (*Block, error) {
	b := &Block{
		config:       opts.Config,
		logger:       opts.Logger,
		blockService: opts.BlockService,
	}

	group := opts.HttpServer.Group("/blocks")

	group.POST("", b.blockService.HandleWriteBlock)
	group.GET("", b.blockService.HandlegetBlocks)

	return b, nil
}

type Block struct {
	config       config.Config
	logger       *zap.Logger
	blockService *blocks.Service
}
