package providers

import (
	"github.com/objforce/objforce/cmd/meta-api/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLoggerProvider(config *config.AppConfig) *zap.SugaredLogger {
	var log *zap.Logger
	if config.Env != "production" {
		logConfig := zap.NewDevelopmentConfig()
		logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		log, _ = logConfig.Build()
	} else {
		log, _ = zap.NewProduction()
	}
	return log.Sugar()
}