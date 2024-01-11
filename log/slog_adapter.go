package log

import (
	"go.uber.org/zap/exp/zapslog"
	"log/slog"
)

func InitLogger(cfg ConfigLog) *slog.Logger {
	zapLogger := InitZapLogger(cfg)
	return slog.New(zapslog.NewHandler(zapLogger.Core(), nil))
}
