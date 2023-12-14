// Package common provides utility functions shared across the project.
package common

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitialiseLogger initalises the Zap logger with the specificed log level
func InitialiseLogger(cfg Config) (*zap.SugaredLogger, error) {
	var logLevel zapcore.Level
	err := logLevel.UnmarshalText([]byte(cfg.LogLevel))
	if err != nil {
		return nil, err
	}

	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(logLevel),
		Encoding:         "console",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	// Using the sugared logger
	sugar := logger.Sugar()
	sugar.Debug("Logger construction successful")

	zap.ReplaceGlobals(logger) // Replace the global logger
	return sugar, nil
}
