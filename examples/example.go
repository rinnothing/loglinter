package main

import (
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	slog.Info("hello")

	slogger := slog.New(slog.Default().Handler())
	slogger.Debug("first")

	zap.L().Error("second")

	zlogger := zap.NewExample()
	defer zlogger.Sync()

	zlogger.Info("bye")
}
