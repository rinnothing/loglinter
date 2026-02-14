package p

import (
	"log/slog"

	"go.uber.org/zap"
)

func nonEnglish() {
	slog.Info("по-русски") // want "log messages should be english only, but 1 char in \"по-русски\" isn't"

	slogger := slog.New(slog.Default().Handler())
	slogger.Info("по-русски") // want "log messages should be english only, but 1 char in \"по-русски\" isn't"

	zap.L().Error("по-русски") // want "log messages should be english only, but 1 char in \"по-русски\" isn't"

	zap.S().Error("по-русски")                     // want "log messages should be english only, but 1 char in \"по-русски\" isn't"
	zap.S().Errorf("по-русски %s", "zap template") // want "log messages should be english only, but 1 char in \"по-русски %s\" isn't"
	zap.S().Errorln("по-русски")                   // want "log messages should be english only, but 1 char in \"по-русски\" isn't"
	zap.S().Errorw("по-русски", "key", "value")    // want "log messages should be english only, but 1 char in \"по-русски\" isn't"
}
