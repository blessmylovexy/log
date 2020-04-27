package logger

import "go.uber.org/zap"

/*Config log配置文件*/
type Config struct {
	defaultLogLevel string          //默认日志记录级别
	stacktraceLevel string          //记录堆栈的级别
	atomicLevel     zap.AtomicLevel //用于动态更改日志记录级别
	projectName     string          //项目名称
	callerSkip      int             //CallerSkip次数
	jsonFormat      bool            //输出json格式
	consoleOut      bool            //是否输出到console
	fileOut         *fileOut
}

type fileOut struct {
	enable        bool   //是否将日志输出到文件
	path          string //日志保存路径
	name          string //日志保存的名称，不些随机生成
	rotationTime  uint   //日志切割时间间隔(小时)
	rotationCount uint   //文件最大保存份数
}

func newConfig() *Config {
	return &Config{
		defaultLogLevel: "info",
		stacktraceLevel: "panic",
		atomicLevel:     zap.NewAtomicLevel(),
		projectName:     "",
		callerSkip:      1,
		jsonFormat:      true,
		consoleOut:      true,
		fileOut: &fileOut{
			enable:        false,
			path:          "",
			name:          "",
			rotationTime:  24,
			rotationCount: 7,
		},
	}
}

/*SetLevel 设置日志记录级别*/
func (c *Config) SetLevel(level string) {
	c.atomicLevel.SetLevel(getLevel(level))
}

/*SetStacktraceLevel 设置堆栈跟踪的日志级别*/
func (c *Config) SetStacktraceLevel(level string) {
	c.stacktraceLevel = level
}

/*SetProjectName 设置ProjectName*/
func (c *Config) SetProjectName(projectName string) {
	c.projectName = projectName
}

/*SetCallerSkip 设置callerSkip次数*/
func (c *Config) SetCallerSkip(callerSkip int) {
	c.callerSkip = callerSkip
}

/*EnableJSONFormat 开启JSON格式化输出*/
func (c *Config) EnableJSONFormat() {
	c.jsonFormat = true
}

/*DisableJSONFormat 关闭JSON格式化输出*/
func (c *Config) DisableJSONFormat() {
	c.jsonFormat = false
}

/*EnableConsoleOut 开启Console输出*/
func (c *Config) EnableConsoleOut() {
	c.consoleOut = true
}

/*DisableConsoleOut 关闭Console输出*/
func (c *Config) DisableConsoleOut() {
	c.consoleOut = false
}

/*SetFileOut 设置日志输出文件*/
func (c *Config) SetFileOut(path, name string, rotationTime, rotationCount uint) {
	c.fileOut.enable = true
	c.fileOut.path = path
	c.fileOut.name = name
	c.fileOut.rotationTime = rotationTime
	c.fileOut.rotationCount = rotationCount
}
