package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Logger() *zap.Logger {
	lc := zap.NewDevelopmentConfig()
	lc.DisableStacktrace = true
	lc.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := lc.Build()
	return logger
}
