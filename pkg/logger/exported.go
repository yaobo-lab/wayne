package logger

import (
	"fmt"

	"go.uber.org/zap"
)

// WithField creates an entry from the standard logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithField(key string, value interface{}) *zap.SugaredLogger {
	fs := []zap.Field{zap.Any(key, value)}
	return zap.L().WithOptions(zap.AddCallerSkip(-1)).With(fs...).Sugar()
}

// WithFields creates an entry from the standard logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithFields(fields map[string]interface{}) *zap.SugaredLogger {
	if len(fields) > 0 {
		fs := make([]zap.Field, 0)
		for key, value := range fields {
			fs = append(fs, zap.Any(key, value))
		}
		return zap.L().WithOptions(zap.AddCallerSkip(-1)).With(fs...).Sugar()
	}
	return zap.S()
}

// Debug logs a message at level Debug on the standard zap.S().
func Debug(args ...interface{}) {
	zap.S().Debug(args...)
}

// Print logs a message at level Info on the standard zap.S().
func Print(args ...interface{}) {
	fmt.Print(args...)
}

// Info logs a message at level Info on the standard zap.S().
func Info(args ...interface{}) {
	zap.S().Info(args...)
}

// Warn logs a message at level Warn on the standard zap.S().
func Warn(args ...interface{}) {
	zap.S().Warn(args...)
}

// Error logs a message at level Error on the standard zap.S().
func Error(args ...interface{}) {
	zap.S().Error(args...)
}

// Panic logs a message at level Panic on the standard zap.S().
func Panic(args ...interface{}) {
	zap.S().Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	zap.S().Fatal(args...)
}

// Debugf logs a message at level Debug on the standard zap.S().
func Debugf(format string, args ...interface{}) {
	zap.S().Debugf(format, args...)
}

// Printf logs a message at level Info on the standard zap.S().
func Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// Infof logs a message at level Info on the standard zap.S().
func Infof(format string, args ...interface{}) {
	zap.S().Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard zap.S().
func Warnf(format string, args ...interface{}) {
	zap.S().Warnf(format, args...)
}

// Errorf logs a message at level Error on the standard zap.S().
func Errorf(format string, args ...interface{}) {
	zap.S().Errorf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Panicf(format string, args ...interface{}) {
	zap.S().Panicf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	zap.S().Fatalf(format, args...)
}

// WarnNTrace logs a message at level Warn on the standard zap.S().
func WarnNTrace(args ...interface{}) {
	l := zap.S().WithOptions(zap.AddStacktrace(zap.FatalLevel + 1))
	l.Warn(args...)
}

// ErrorNTrace logs a message at level Error on the standard zap.S().
func ErrorNTrace(args ...interface{}) {
	l := zap.S().WithOptions(zap.AddStacktrace(zap.FatalLevel + 1))
	l.Error(args...)
}

// WarnNTracef logs a message at level Warn on the standard zap.S().
func WarnNTracef(format string, args ...interface{}) {
	l := zap.S().WithOptions(zap.AddStacktrace(zap.FatalLevel + 1))
	l.Warnf(format, args...)
}

// ErrorNTracef logs a message at level Error on the standard zap.S().
func ErrorNTracef(format string, args ...interface{}) {
	l := zap.S().WithOptions(zap.AddStacktrace(zap.FatalLevel + 1))
	l.Errorf(format, args...)
}
