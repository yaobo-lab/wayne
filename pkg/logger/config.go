package logger

import (
	"os"
	"regexp"
)

// LogConfig 统一配置类
type LogConfig struct {
	Log *LogOption `toml:"log"`
}

type (
	// TraceIDFunc 定义获取跟踪ID的函数
	TraceIDFunc func() string
	// LogLevel type
	LogLevel uint32
)

// 日志配置
type LogOption struct {
	AppNo                  string      `toml:"app_no" json:"AppNo"`                                    // 应用编号
	AppName                string      `toml:"app_name" json:"AppName"`                                // 应用名称
	Level                  LogLevel    `toml:"log_level" json:"Level"`                                 // 日志级别(debug,info,warn,error,dpanic,panic,fatal)
	Format                 string      `toml:"format" json:"Format"`                                   // 日志格式（支持输出格式：text/json）
	Output                 string      `toml:"output" json:"Output"`                                   // 日志输出(支持：stdout/stderr/file/multi)
	OutputFile             string      `toml:"output_file" json:"OutputFile"`                          // 指定日志输出的文件路径 logs/app.log
	DisableCustomTimestamp bool        `toml:"disable_custom_timestamp" json:"DisableCustomTimestamp"` // 是否禁用自定义时间戳显示（无用）
	TimeFormat             string      `toml:"time_format" json:"TimeFormat"`                          // 日志时间格式（默认为时间戳timestamp）
	DisableLineHook        bool        `toml:"disable_line_hook" json:"DisableLineHook"`               // 是否禁用行号信息显示(WarnLevel以上才会显示)
	DisableSoftLink        bool        `toml:"disable_soft_link" json:"DisableSoftLink"`               // 是否禁用软链接
	DisableMaxAge          bool        `toml:"disable_max_age" json:"DisableMaxAge"`                   // 是否禁用文件清理
	LogFileMaxAge          int         `toml:"log_file_max_age" json:"LogFileMaxAge"`                  // 设置日志文件清理前的最长保存时间 天数
	LogFileRotationTime    int         `toml:"log_file_rotation_time" json:"LogFileRotationTime"`      // 设置日志分割的时间
	LogFileRotationCount   int         `toml:"log_file_rotation_count" json:"LogFileRotationCount"`    // 设置日志文件分割数量
	LogFilePathFormat      string      `toml:"log_file_path_format" json:"LogFilePathFormat"`          // 设置日志文件名规则
	LogFileMaxNum          int         `toml:"log_file_max_num" json:"LogFileMaxNum"`                  // 设置日志文件分割最大保留数量
	LogFileMaxSize         int64       `toml:"log_file_max_size" json:"LogFileMaxSize"`                // 设置单个日志文件最大值（单位：MB）
	FrmDeviceID            string      // 来源设备ID
	ProductType            string      // 来源设备所属产品类型
	ProductSubType         string      // 来源设备所属产品子类
	TIDFunc                TraceIDFunc // 获取跟踪ID的函数
	CallerSkip             int         //解决打印日志堆栈跳过层级
	Filter                 []string    //过滤的正则表达式
}

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel LogLevel = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

func (o *LogOption) GetExp() []*regexp.Regexp {
	if o.Filter == nil {
		return nil
	}
	arr := make([]*regexp.Regexp, 0)
	for _, v := range o.Filter {
		if exp, err := regexp.Compile(v); err == nil {
			arr = append(arr, exp)
		}
	}
	return arr
}
func (lvl LogLevel) Sting() string {
	switch lvl {
	case PanicLevel:
		return "panic"
	case FatalLevel:
		return "fatal"
	case ErrorLevel:
		return "error"
	case WarnLevel:
		return "warn"
	case InfoLevel:
		return "info"
	case DebugLevel:
		return "debug"
	case TraceLevel:
		return "trace"
	}
	return ""
}

// 判断文件是否存在
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
