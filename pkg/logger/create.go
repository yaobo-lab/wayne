package logger

import (
	"io"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var writer io.Writer

// 创建logger
func Create(option *LogOption) *zap.Logger {
	option = checkDefault(option)
	// 日志级别
	lv := convertLevel(option.Level)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = TimestampKey
	encoderConfig.CallerKey = SourceKey
	encoderConfig.LevelKey = "type"
	encoderConfig.MessageKey = "message"

	if option.TimeFormat == "time" {
		encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	} else {
		encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(t.UnixNano() / 1e6)
		}
	}

	// 日志打印格式
	var encoder zapcore.Encoder
	if option.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}
	writer = getWriter(option)
	arrCore := make([]zapcore.Core, 0)
	if option.Filter == nil || len(option.Filter) == 0 {
		arrCore = append(arrCore, zapcore.NewCore(encoder, zapcore.AddSync(writer), zap.NewAtomicLevelAt(lv)))
	} else {
		arrCore = append(arrCore, newCore(encoder, zapcore.AddSync(writer), zap.NewAtomicLevelAt(lv), option.GetExp()))
	}
	core := zapcore.NewTee(
		arrCore...,
	)
	l := zap.New(
		core,
		zap.AddStacktrace(zap.WarnLevel),
	)

	// 日志是否显示行号(warn级别以上才显示行号)
	if !option.DisableLineHook {
		// 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数
		l = l.WithOptions(
			zap.AddCaller(),
			zap.AddCallerSkip(option.CallerSkip), // 解决打印日志调用文件问题
		)
	}

	// 添加其他字段
	fields := make(map[string]interface{})
	if option.AppNo != "" {
		fields[AppNoKey] = option.AppNo
	}
	if len(option.AppName) > 0 {
		fields[AppNameKey] = option.AppName
	}
	if len(option.FrmDeviceID) > 0 {
		fields[FromDeviceIDKey] = option.FrmDeviceID
	}
	if len(option.ProductType) > 0 {
		fields[ProductTypeKey] = option.ProductType
	}
	if len(option.ProductSubType) > 0 {
		fields[ProductSubTypeKey] = option.ProductSubType
	}
	if option.TIDFunc != nil {
		fields[TraceIDKey] = option.TIDFunc()
	}
	if len(fields) > 0 {
		fs := make([]zap.Field, 0)
		for key, value := range fields {
			fs = append(fs, zap.Any(key, value))
		}
		l.With(fs...)

	}

	return l
}

// 创建logger 并设置为全局
func CreateLogger(option *LogOption) *zap.Logger {
	log := Create(option)
	zap.ReplaceGlobals(log)
	return log
}

// 使用默认参数创建日志并设置为全局日志
func CreateLoggerDefault() *zap.Logger {
	option := &LogOption{}
	return CreateLogger(option)
}

func GetLogWriter() io.Writer {
	return writer
}
