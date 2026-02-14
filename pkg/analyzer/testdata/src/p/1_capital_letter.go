package p

import (
	"log/slog"

	"go.uber.org/zap"
)

func capitalLetter() {
	slog.Info("Slog") // want "log messages should start from lowercase letter, but \"Slog\" doesn't"

	slogger := slog.New(slog.Default().Handler())
	slogger.Info("Slog struct") // want "log messages should start from lowercase letter, but \"Slog struct\" doesn't"

	zap.L().Error("Zap hello") // want "log messages should start from lowercase letter, but \"Zap hello\" doesn't"

	zap.S().Error("Zap sugar hello")                           // want "log messages should start from lowercase letter, but \"Zap sugar hello\" doesn't"
	zap.S().Errorf("Zap sugar template %s", "zap template")    // want "log messages should start from lowercase letter, but \"Zap sugar template %s\" doesn't"
	zap.S().Errorln("Zap sugar with newline")                  // want "log messages should start from lowercase letter, but \"Zap sugar with newline\" doesn't"
	zap.S().Errorw("Zap sugar with key-value", "key", "value") // want "log messages should start from lowercase letter, but \"Zap sugar with key-value\" doesn't"
}
