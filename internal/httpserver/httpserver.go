package httpserver

import (
	"context"
	"fmt"

	"github.com/ricardovac/go-blockchain/internal/config"
	"go.uber.org/zap"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Opts struct {
	fx.In
	LC     fx.Lifecycle
	Config config.Config
	Logger *zap.Logger
}

func New(opts Opts) *Server {
	gin.SetMode(gin.ReleaseMode)
	s := &Server{
		Engine: gin.New(),
	}

	s.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/health"),
		cors.New(
			cors.Config{
				AllowAllOrigins:  true,
				AllowMethods:     []string{"*"},
				AllowHeaders:     []string{"*"},
				ExposeHeaders:    []string{"*"},
				AllowCredentials: true,
			},
		),
		gin.Recovery(),
	)

	opts.LC.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					s.StartServer(opts.Config.Port)
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				s.StopServer()
				return nil
			},
		},
	)

	return s
}

type Server struct {
	*gin.Engine
	Api *gin.RouterGroup
}

func (c *Server) StartServer(port int) {
	c.Run(fmt.Sprintf(":%d", port))
}

func (c *Server) StopServer() {

}
