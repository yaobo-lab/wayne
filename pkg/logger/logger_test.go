package logger

import (
	"testing"
	"time"

	"go.uber.org/zap"
)

type Job struct {
}

func (s *Job) Do() error {
	Errorf("this is a error message: %s", "hello world")
	return nil
}

func TestLog(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", "https://www.badu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func TestDebug(t *testing.T) {
	option := &LogOption{}
	CreateLogger(option)
	Debugf("this is a debug message: %s", "hello world")
}

func TestDebug2(t *testing.T) {
	option := &LogOption{}
	CreateLogger(option)
}

func TestInfo(t *testing.T) {
	option := &LogOption{}
	CreateLogger(option)
	Infof("this is a info message: %s", "hello world")
}

func TestWarn(t *testing.T) {
	option := &LogOption{}
	CreateLogger(option)
	Warnf("this is a warn message: %s", "hello world")
}

func TestPanic(t *testing.T) {
	option := &LogOption{}
	CreateLogger(option)
	Panicf("this is a panic message: %s", "hello world")
}

func TestFatalf(t *testing.T) {
	option := &LogOption{}
	CreateLogger(option)
	Fatalf("this is a fatal error message: %s", "hello world")
}

func TestFilter(t *testing.T) {

	option := &LogOption{
		OutputFile:      "logs",
		Filter:          []string{"logger1"},
		Level:           DebugLevel,
		Output:          "multi",
		DisableLineHook: false,
	}
	CreateLogger(option)
	Debugf("this is a info message: %s", "hello world")
}

func Test_日志文件按数量(t *testing.T) {
	option := &LogOption{
		OutputFile:           "logs",
		Level:                DebugLevel,
		Output:               "multi",
		DisableSoftLink:      true,
		LogFileRotationCount: 3,
		LogFileMaxSize:       1,
		LogFilePathFormat:    "mqtt.msg",
		CallerSkip:           -1,
	}

	log := CreateLogger(option).Sugar()
	for {
		log.Info("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	}
}
