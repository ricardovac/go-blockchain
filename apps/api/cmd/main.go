package main

import (
	"github.com/ricardovac/go-blockchain/internal/app"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		app.App,
	).Run()
}
