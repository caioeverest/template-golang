package logger

import (
	"os"
	"strings"

	"github.com/{{.Author}}/{{.RepositoryName}}/infra/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

var _global *Logger

func Get(conf config.Config) *Logger {
	if _global == nil {
		load(conf)
	}
	return _global
}

func load(conf config.Config) {
	var (
		logLevel   zapcore.Level
		encoderCfg zapcore.EncoderConfig
		logOutput  = os.Stdout
	)
	_global = &Logger{}

	switch strings.ToUpper(conf.Env) {
	case config.Development:
		logLevel = zap.DebugLevel
	default:
		logLevel = zap.InfoLevel
	}

	encoderCfg = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		NameKey:        "logger",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), logOutput, logLevel)
	l := zap.New(core).With(
		zap.String("api", conf.API),
		zap.String("version", conf.Version),
		zap.String("environment", conf.Env),
	)

	_global.SugaredLogger = l.Sugar()
}

func (l Logger) Printf(format string, args ...interface{}) {
	l.Infof(format, args...)
}
