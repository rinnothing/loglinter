package p

import (
	"log/slog"

	"go.uber.org/zap"
)

func correct() {
	slog.Info("slog")

	slogger := slog.New(slog.Default().Handler())
	slogger.Info("slog struct")

	zap.L().Error("zap hello")

	zap.S().Error("zap sugar hello")
	zap.S().Errorf("zap sugar template %s", "zap template")
	zap.S().Errorln("zap sugar with newline")
	zap.S().Errorw("zap sugar with key-value", "key", "value")
}
