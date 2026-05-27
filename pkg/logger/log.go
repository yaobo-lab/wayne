package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger(name string, level string) *Logger {
	l := zap.S().Named(name)
	if level != "" {
		levelEnum, err := zapcore.ParseLevel(level)
		if err == nil {
			l = l.WithOptions(zap.IncreaseLevel(levelEnum))
		}
	}
	log := &Logger{SugaredLogger: l}
	return log
}

func (log *Logger) WithNamed(name string) *Logger {
	return &Logger{SugaredLogger: log.SugaredLogger.Named(name)}
}

// WarnNTrace logs a message at level Warn on the standard zap.S().
func (log *Logger) WarnNTrace(args ...interface{}) {
	l := log.WithOptions(zap.AddStacktrace(zap.FatalLevel + 1))
	l.Warn(args...)
}

// ErrorNTrace logs a message at level Error on the standard zap.S().
func (log *Logger) ErrorNTrace(args ...interface{}) {
	l := log.WithOptions(zap.AddStacktrace(zap.FatalLevel + 1))
	l.Error(args...)
}

// WarnNTracef logs a message at level Warn on the standard zap.S().
func (log *Logger) WarnNTracef(format string, args ...interface{}) {
	l := log.WithOptions(zap.AddStacktrace(zap.FatalLevel + 1))
	l.Warnf(format, args...)
}

// ErrorNTracef logs a message at level Error on the standard zap.S().
func (log *Logger) ErrorNTracef(format string, args ...interface{}) {
	l := log.WithOptions(zap.AddStacktrace(zap.FatalLevel + 1))
	l.Errorf(format, args...)
}
