package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newZapLogger() *zap.SugaredLogger {
	var cfg *zap.Config
	if _debugMode {
		cfg = &zap.Config{
			Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
			Development: true,
			Encoding:    "console",
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey:       "M",
				LevelKey:         "L",
				TimeKey:          "T",
				NameKey:          "N",
				CallerKey:        "C",
				FunctionKey:      zapcore.OmitKey,
				StacktraceKey:    "S",
				LineEnding:       zapcore.DefaultLineEnding,
				EncodeLevel:      zapcore.CapitalLevelEncoder,
				EncodeTime:       zapcore.ISO8601TimeEncoder,
				EncodeDuration:   zapcore.StringDurationEncoder,
				EncodeCaller:     zapcore.ShortCallerEncoder,
				EncodeName:       zapcore.FullNameEncoder,
				ConsoleSeparator: "    ",
			},
			OutputPaths:      []string{"stderr"},
			ErrorOutputPaths: []string{"stderr"},
		}
	} else {
		cfg = &zap.Config{
			Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
			Development: false,
			Sampling: &zap.SamplingConfig{
				Initial:    100,
				Thereafter: 100,
			},
			Encoding: "json",
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey:     "msg",
				LevelKey:       "level",
				TimeKey:        "ts",
				NameKey:        "logger",
				CallerKey:      "caller",
				FunctionKey:    "func",
				StacktraceKey:  "stacktrace",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.LowercaseLevelEncoder,
				EncodeTime:     zapcore.EpochTimeEncoder,
				EncodeDuration: zapcore.SecondsDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
				EncodeName:     zapcore.FullNameEncoder,
			},
			OutputPaths:      []string{"stderr"},
			ErrorOutputPaths: []string{"stderr"},
		}
	}
	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return l.Sugar()
}
