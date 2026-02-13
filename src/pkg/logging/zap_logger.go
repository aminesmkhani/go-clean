package logging

import (
	"github.com/aminesmkhani/go-clean/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logLevelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zap.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

type zapLogger struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
}

func NewZapLogger(cfg *config.Config) *zapLogger {
	logger := &zapLogger{cfg: cfg}
	logger.Init()
	return logger
}

func (l *zapLogger) getLogLevel() zapcore.Level{
	level, exists := logLevelMap[l.cfg.Logger.Level]
	if !exists {
		return zapcore.DebugLevel
	}
	return level
}

func (l *zapLogger) Init() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename: l.cfg.Logger.FilePath,
		MaxSize: 1,
		MaxAge: 5,
		LocalTime: true,
		MaxBackups: 10,
		Compress: true,
	})

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		l.getLogLevel(),
	)

	logger := zap.New(core,zap.AddCaller(),zap.AddCallerSkip(1),zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()

	l.logger = logger
}

func (l *zapLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.Debugw(msg)
}

func (l *zapLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template,args)
}
