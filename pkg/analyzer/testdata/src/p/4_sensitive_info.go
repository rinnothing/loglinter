package p

import (
	"log/slog"

	"go.uber.org/zap"
)

func sensitiveInfo() {
	slog.Info("user password=") // want "log messages shouldn't contain sensitive information but there's \"password=\" in string \"user password=\""

	slogger := slog.New(slog.Default().Handler())
	slogger.Info("api key=") // want "log messages shouldn't contain sensitive information but there's \"key=\" in string \"api key=\""

	zap.L().Error("security token=") // want "log messages shouldn't contain sensitive information but there's \"token=\" in string \"security token=\""

	zap.S().Error("the password=")                     // want "log messages shouldn't contain sensitive information but there's \"password=\" in string \"the password=\""
	zap.S().Errorf("user password=%s", "zap template") // want "log messages shouldn't contain sensitive information but there's \"password=\" in string \"user password=%s\""
	zap.S().Errorln("api key=")                        // want "log messages shouldn't contain sensitive information but there's \"key=\" in string \"api key=\""
	zap.S().Errorw("security token=", "key", "value")  // want "log messages shouldn't contain sensitive information but there's \"token=\" in string \"security token=\""
}
