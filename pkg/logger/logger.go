package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/lestrrat-go/strftime"

	rotatelogs "wayne/pkg/file-rotatelogs"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 定义键名
const (
	TraceIDKey        = "tid"
	SpanTitleKey      = "span_title"
	SpanFunctionKey   = "span_function"
	VersionKey        = "version"
	AppNameKey        = "name"
	AppNoKey          = "appno"
	TimestampKey      = "timestamp"
	FromDeviceIDKey   = "frmDeviceID"
	ProductTypeKey    = "productType"
	ProductSubTypeKey = "productSubType"

	SourceKey = "source"
)

func checkDefault(option *LogOption) *LogOption {
	if option.Level <= 0 {
		option.Level = WarnLevel
	}
	if option.Format == "" {
		option.Format = "json"
	}
	if option.Output == "" {
		option.Output = "stdout"
	}
	if option.OutputFile == "" {
		option.OutputFile = "logs"
	}
	if option.TimeFormat == "" {
		option.TimeFormat = "time"
	}
	if option.LogFileMaxAge <= 0 {
		option.LogFileMaxAge = 7
	}
	if option.LogFileRotationTime <= 0 {
		option.LogFileRotationTime = 24 * 60 * 60
	}
	if option.LogFilePathFormat == "" {
		option.LogFilePathFormat = "log.%Y-%m-%d.log"
	}
	if option.LogFileMaxNum <= 0 {
		option.LogFileMaxNum = 5
	}
	if option.LogFileMaxSize <= 0 {
		option.LogFileMaxSize = 50
	}
	if option.CallerSkip == 0 {
		option.CallerSkip = 1
	} else if option.CallerSkip < 0 {
		option.CallerSkip = 0
	}
	if option.LogFileRotationCount > 0 {
		option.DisableMaxAge = false
		option.LogFileMaxAge = -1
		option.LogFileMaxNum = option.LogFileRotationCount
	}
	return option
}

// 获取写入io
func getWriter(opt *LogOption) io.Writer {
	switch opt.Output {
	case "file", "multi":
		fileName := filepath.Join(opt.OutputFile, opt.LogFilePathFormat)
		options := []rotatelogs.Option{
			rotatelogs.WithRotationTime(time.Duration(opt.LogFileRotationTime)), // 日志切割时间间隔
			rotatelogs.WithRotationSize((1 << 20) * opt.LogFileMaxSize),         //默认50M
			rotatelogs.WithRotationCount(uint(opt.LogFileRotationCount)),        //日志切割数量
			rotatelogs.ForceNewFile(),                                           //必须
			rotatelogs.WithHandler(rotatelogs.HandlerFunc(func(e rotatelogs.Event) {
				if e.Type() != rotatelogs.FileRotatedEventType {
					return
				}

				baseFile := baseFilename(fileName, time.Duration(opt.LogFileRotationTime))
				if baseFile == "" {
					return
				}
				pattern := fmt.Sprintf("%s*", baseFile)
				matches, err := filepath.Glob(pattern)
				if err != nil {
					return
				}

				files := make(map[string]int64)
				for _, p := range matches {
					//忽略锁定的文件
					if strings.HasSuffix(p, "_lock") || strings.HasSuffix(p, "_symlink") {
						continue
					}

					fi, err := os.Stat(p)
					if err != nil {
						continue
					}

					//日志文件修改时间排序
					files[p] = fi.ModTime().Unix()
				}

				//超过最大日志文件数，则删除最早创建的
				if len(matches) > opt.LogFileMaxNum {
					delNum := len(matches) - opt.LogFileMaxNum
					rotationFile(files, delNum)
				}
			})),
		}

		if !opt.DisableMaxAge {
			options = append(options, rotatelogs.WithMaxAge(time.Duration(opt.LogFileMaxAge)*24*time.Hour)) // 文件最大保存时间
		} else {
			// 如果禁用文件清理，则保留30天的日志
			options = append(options, rotatelogs.WithMaxAge(30*24*time.Hour)) // 文件最大保存时间
		}

		// 添加日志软链接
		if !opt.DisableSoftLink {
			options = append(options, rotatelogs.WithLinkName("log")) // 生成软链，指向最新日志文件
		}
		logWriter, err := rotatelogs.New(
			fileName,
			options...,
		)

		if err != nil {
			panic(err)
		}

		if opt.Output == "multi" {
			out := io.MultiWriter(os.Stdout, logWriter)
			return out
		}
		return logWriter
	default:
		sink, _, err := zap.Open(opt.Output)
		if err != nil {
			panic(err)
		}
		return sink
	}
}

// 日志级别转换到对应zap日志级别
func convertLevel(lvl LogLevel) zapcore.Level {
	switch lvl {
	case PanicLevel:
		return zapcore.PanicLevel
	case FatalLevel:
		return zapcore.FatalLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case InfoLevel:
		return zapcore.InfoLevel
	case DebugLevel:
		return zapcore.DebugLevel
	}
	return -1
}

func baseFilename(path string, rotationTime time.Duration) string {
	pattern, err := strftime.New(path)
	if err != nil {
		return ""
	}

	now := time.Now()

	var base time.Time
	if now.Location() != time.UTC {
		base = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.UTC)
		base = base.Truncate(rotationTime)
		base = time.Date(base.Year(), base.Month(), base.Day(), base.Hour(), base.Minute(), base.Second(), base.Nanosecond(), base.Location())
	} else {
		base = now.Truncate(rotationTime)
	}
	return pattern.FormatString(base)
}

// 滚动覆盖日志文件
func rotationFile(files map[string]int64, delNum int) {
	type KVSort struct {
		Key string
		Val int64
	}

	sortList := make([]KVSort, 0)

	for k, v := range files {
		sortList = append(sortList, KVSort{Key: k, Val: v})
	}

	sort.Slice(sortList, func(i, j int) bool {
		return sortList[i].Val < sortList[j].Val // 升序
	})

	//需要删除的日志文件数
	num := delNum
	for _, value := range sortList {
		if num < 1 {
			break
		}

		_ = os.Remove(value.Key)
		num--
	}
}
