package app

import (
	"github.com/ricardovac/go-blockchain/internal/config"
	"github.com/ricardovac/go-blockchain/internal/httpserver"
	"github.com/ricardovac/go-blockchain/internal/httpserver/routes/blocks"
	blockservice "github.com/ricardovac/go-blockchain/internal/services/blocks"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var App = fx.Options(
	fx.Provide(
		config.New,
		zap.NewDevelopment,
		blockservice.New,
	),

	fx.Provide(
		httpserver.New,
	),

	fx.Invoke(
		blocks.New,
	),
)
