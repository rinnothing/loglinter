package main

import (
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	slog.Info("Это первый лог")

	slogger := slog.New(slog.Default().Handler())
	slogger.Debug("This is the second log")

	zap.L().Error("and this is the 3d")

	zlogger := zap.NewExample()
	defer zlogger.Sync()

	zlogger.Info("bye!!!!⭐")

	zlogger.Info("password=")
}
