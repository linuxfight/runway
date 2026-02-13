package app

import (
	"fmt"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func New() *fx.App {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("[CONFIG] .env file not found, loading with system enviroment...")
	} else {
		fmt.Println("[CONFIG] .env file found, loading...")

	}

	return fx.New(
		Module,
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{
				Logger: log,
			}
		}),
	)
}
