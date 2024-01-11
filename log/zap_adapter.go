package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Level string

type ConfigLog struct {
	Environment string
	Level       string
}

const (
	TRACE Level = "TRACE"
	DEBUG Level = "DEBUG"
	INFO  Level = "INFO"
	WARN  Level = "WARN"
	ERROR Level = "ERROR"
	PANIC Level = "PANIC"
)

func toZapCoreLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	default:
		return zapcore.InfoLevel
	}
}
func InitZapLogger(cfg ConfigLog) *zap.Logger {
	if cfg.Environment == "development" {
		stdout := zapcore.AddSync(os.Stdout)
		level := zap.NewAtomicLevelAt(toZapCoreLevel(cfg.Level))
		developmentCfg := zap.NewDevelopmentEncoderConfig()
		developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)
		core := zapcore.NewCore(consoleEncoder, stdout, level)
		return zap.New(core)
	} else { // for production
		stdout := zapcore.AddSync(os.Stdout)
		level := zap.NewAtomicLevelAt(toZapCoreLevel(cfg.Level))
		productionCfg := zap.NewProductionEncoderConfig()
		productionCfg.TimeKey = "timestamp"
		productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		jsonEncoder := zapcore.NewJSONEncoder(productionCfg)
		core := zapcore.NewCore(jsonEncoder, stdout, level)
		return zap.New(core)
	}
}
