package logging

import (
	"github.com/aminesmkhani/go-clean/config"
	"github.com/rs/zerolog"
	"go.uber.org/zap/zapcore"
)

type zeroLogger struct {
	cfg    *config.Config
	logger *zerolog.Logger
}

var zeroLogLevelMap = map[string]zerolog.Level{
	"debug": zerolog.DebugLevel,
	"info":  zerolog.InfoLevel,
	"warn":  zerolog.WarnLevel,
	"error": zerolog.ErrorLevel,
	"fatal": zerolog.FatalLevel,
}

func newZeroLogger(cfg *config.Config) *zeroLogger {

}

func (l *zeroLogger) getLogLevel() zerolog.Level {
	level, exists := zeroLogLevelMap[l.cfg.Logger.Level]
	if !exists {
		return zerolog.DebugLevel
	}
	return level
}
