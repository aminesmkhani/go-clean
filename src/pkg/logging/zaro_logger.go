package logging

import (
	"os"

	"github.com/aminesmkhani/go-clean/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
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
	logger := &zeroLogger{cfg: cfg}
	logger.Init()
	return logger
}

func (l *zeroLogger) getLogLevel() zerolog.Level {
	level, exists := zeroLogLevelMap[l.cfg.Logger.Level]
	if !exists {
		return zerolog.DebugLevel
	}
	return level
}


func (l *zeroLogger) Init() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	file, err := os.OpenFile(l.cfg.Logger.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Could not open log file")
	}
	var logger = zerolog.New(file).
		With().
		Timestamp().
		Str("AppName","MyApp").
		Str("LoggerName","ZeroLog").
		Logger()
	zerolog.SetGlobalLevel(l.getLogLevel())
	l.logger = &logger
}



func (l *zeroLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, cat, sub)
	l.logger.
	Debug()
	Str("category", string(cat)).
	Str("subcategory", string(sub)).
	Fields(params).
	Msg(msg).
}

func (l *zeroLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args)
}