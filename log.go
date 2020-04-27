package log

import (
	"github.com/blessmylovexy/log/logger"
)

var (
	wlog *logger.Logger
	/*Config log配置文件*/
	Config *logger.Config
)

func init() {
	wlog = logger.New()
	Config = wlog.Config
	wlog.Config.SetCallerSkip(2)
	wlog.ApplyConfig()
}

/*ApplyConfig 应用当前Config配置*/
func ApplyConfig() {
	wlog.ApplyConfig()
}

/*Debug Debug log*/
func Debug(args ...interface{}) {
	wlog.Debug(args...)
}

/*Debugf Debug format log*/
func Debugf(template string, args ...interface{}) {
	wlog.Debugf(template, args...)
}

/*Debugw Debugw log*/
func Debugw(msg string, keysAndValues ...interface{}) {
	wlog.Debugw(msg, keysAndValues...)
}

/*Info Info log*/
func Info(args ...interface{}) {
	wlog.Info(args...)
}

/*Infof Info format log*/
func Infof(template string, args ...interface{}) {
	wlog.Infof(template, args...)
}

/*Infow Infow log*/
func Infow(msg string, keysAndValues ...interface{}) {
	wlog.Infow(msg, keysAndValues...)
}

/*Warn Warn log*/
func Warn(args ...interface{}) {
	wlog.Warn(args...)
}

/*Warnf Warn format log*/
func Warnf(template string, args ...interface{}) {
	wlog.Warnf(template, args...)
}

/*Warnw Warnw log*/
func Warnw(msg string, keysAndValues ...interface{}) {
	wlog.Warnw(msg, keysAndValues...)
}

/*Error Error log*/
func Error(args ...interface{}) {
	wlog.Error(args...)
}

/*Errorf Error format log*/
func Errorf(template string, args ...interface{}) {
	wlog.Errorf(template, args...)
}

/*Errorw Errorw log*/
func Errorw(msg string, keysAndValues ...interface{}) {
	wlog.Errorw(msg, keysAndValues...)
}

/*Panic Panic log*/
func Panic(args ...interface{}) {
	wlog.Panic(args...)
}

/*Panicf Panic format log*/
func Panicf(template string, args ...interface{}) {
	wlog.Panicf(template, args...)
}

/*Panicw Panicw log*/
func Panicw(msg string, keysAndValues ...interface{}) {
	wlog.Panicw(msg, keysAndValues...)
}

/*Fatal Fatal log*/
func Fatal(args ...interface{}) {
	wlog.Fatal(args...)
}

/*Fatalf Fatal format log*/
func Fatalf(template string, args ...interface{}) {
	wlog.Fatalf(template, args...)
}

/*Fatalw Fatalw log*/
func Fatalw(msg string, keysAndValues ...interface{}) {
	wlog.Fatalw(msg, keysAndValues...)
}
