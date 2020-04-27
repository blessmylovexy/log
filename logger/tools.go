package logger

import (
	"io"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func getEncoder() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		LevelKey:       "L",
		TimeKey:        "T",
		MessageKey:     "M",
		NameKey:        "N",
		CallerKey:      "C",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func getLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

func getFileWriter(path, name string, rotationTime, rotationCount uint) io.Writer {
	writer, err := rotatelogs.New(
		filepath.Join(path, name+".%Y%m%d%H.log"),
		rotatelogs.WithRotationTime(time.Duration(rotationTime)*time.Hour), // 日志切割时间间隔
		rotatelogs.WithRotationCount(rotationCount),                        // 文件最大保存份数
	)
	if err != nil {
		panic(err)
	}
	return writer
}
