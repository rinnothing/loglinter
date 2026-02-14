package p

import (
	"log/slog"

	"go.uber.org/zap"
)

func specialSymbols() {
	slog.Info("emoji ⭐") // want "log messages shouldn't contain special symbols or emoji, but 7 char in \"emoji ⭐\" seem to break the rule"

	slogger := slog.New(slog.Default().Handler())
	slogger.Info("exclamation mark !") // want "log messages shouldn't contain special symbols or emoji, but 18 char in \"exclamation mark !\" seem to break the rule"

	zap.L().Error("two dots :") // want "log messages shouldn't contain special symbols or emoji, but 10 char in \"two dots :\" seem to break the rule"

	zap.S().Error("result ,")                       // want "log messages shouldn't contain special symbols or emoji, but 8 char in \"result ,\" seem to break the rule"
	zap.S().Errorf("points ... %s", "zap template") // want "log messages shouldn't contain special symbols or emoji, but 8 char in \"points ... %s\" seem to break the rule"
	zap.S().Errorln("emoji ⭐")                      // want "log messages shouldn't contain special symbols or emoji, but 7 char in \"emoji ⭐\" seem to break the rule"
	zap.S().Errorw("emoji ⭐", "key", "value")       // want "log messages shouldn't contain special symbols or emoji, but 7 char in \"emoji ⭐\" seem to break the rule"
}
