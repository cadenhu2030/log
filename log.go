package log

import (
	"flag"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

var (
	DebugMode = flag.Bool("debug", false, "sets log level to debug")

	DebugLevel = zerolog.DebugLevel
	InfoLevel  = zerolog.InfoLevel
	WarnLevel  = zerolog.WarnLevel
	ErrorLevel = zerolog.ErrorLevel
	FatalLevel = zerolog.FatalLevel
	PanicLevel = zerolog.PanicLevel
	NoLevel    = zerolog.NoLevel
	Disabled   = zerolog.Disabled
)

// Init 初始化日志包
func Init(logFile string) {
	// 自定义时间格式
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05.000"

	flag.Parse()

	// 设置日志级别
	level := zerolog.InfoLevel
	if *DebugMode {
		level = zerolog.DebugLevel
	}
	SetLevel(level)

	// 创建文件日志输出，支持滚动
	fileWriter := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    100, // MB
		MaxBackups: 7,   // 保留最近 7 个文件
		MaxAge:     30,  // 天
		Compress:   true,
	}

	// 控制台输出
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: zerolog.TimeFieldFormat,
		NoColor:    false,
	}

	// 双写
	multi := io.MultiWriter(consoleWriter, fileWriter)
	log.Logger = log.Output(multi)
}

// SetLevel 外部可直接设置全局日志级别
func SetLevel(level zerolog.Level) {
	zerolog.SetGlobalLevel(level)
}

// 日志方法封装，带时间戳
func Debug() *zerolog.Event {
	return log.Debug().Timestamp()
}
func Info() *zerolog.Event  { return log.Info().Timestamp() }
func Warn() *zerolog.Event  { return log.Warn().Timestamp() }
func Error() *zerolog.Event { return log.Error().Timestamp() }

// Printf 类似 fmt.Printf，使用 Info 级别
func Printf(format string, v ...interface{}) {
	log.Info().Timestamp().Msg(fmt.Sprintf(format, v...))
}

// Infof
func Infof(format string, v ...interface{}) {
	log.Info().Timestamp().Msg(fmt.Sprintf(format, v...))
}

// Debugf
func Debugf(format string, v ...interface{}) {
	log.Debug().Timestamp().Msg(fmt.Sprintf(format, v...))
}

// Warnf
func Warnf(format string, v ...interface{}) {
	log.Warn().Timestamp().Msg(fmt.Sprintf(format, v...))
}

// Errorf
func Errorf(format string, v ...interface{}) {
	log.Error().Timestamp().Msg(fmt.Sprintf(format, v...))
}

// Fatalf
func Fatalf(format string, v ...interface{}) {
	log.Fatal().Timestamp().Msg(fmt.Sprintf(format, v...))
}
