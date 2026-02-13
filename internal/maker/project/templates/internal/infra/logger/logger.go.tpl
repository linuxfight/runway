package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"{{ .ModulePath }}/app/config"
)

func NewLogger(cfg *config.LogConfig) (*zap.Logger, error) {
	level := zapcore.InfoLevel
	if err := level.Set(cfg.Level); err != nil {
		return nil, fmt.Errorf("invalid log level: %w", err)
	}

	atomicLevel := zap.NewAtomicLevelAt(level)

	if cfg.JSON {
		// === Production / JSON ===
		zapCfg := zap.NewProductionConfig()
		zapCfg.Level = atomicLevel
		return zapCfg.Build()
	}

	// === Development / Pretty ===
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:       "ts",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",

		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		zapcore.AddSync(zapcore.Lock(os.Stdout)),
		atomicLevel,
	)

	return zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	), nil
}
