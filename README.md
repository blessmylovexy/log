### 使用默认的日志记录器：

```go
package main

import (
	"github.com/blessmylovexy/log"
)

func main() {
	log.Info("Info log")
	log.Infof("%s", "InfoF log")

	//动态设置日记级别
	log.Debug("我不应该被输出")
	log.Config.SetLevel("debug")
	log.Debug("Debug log")

	//设置ProjectName
	log.Config.SetProjectName("mlog")
	log.ApplyConfig()
	log.Infof("SetProjectName: %s", "mlog")

	//禁用JSON格式
	log.Config.DisableJSONFormat()
	log.ApplyConfig()
	log.Info("DisableJSONFormat")
}
```

```json
{"L":"INFO","T":"2020-04-27T10:15:47.230+0800","C":"_test_/main.go:8","M":"Info log"}
{"L":"INFO","T":"2020-04-27T10:15:47.230+0800","C":"_test_/main.go:9","M":"InfoF log"}
{"L":"DEBUG","T":"2020-04-27T10:15:47.230+0800","C":"_test_/main.go:14","M":"Debug log"}
{"L":"INFO","T":"2020-04-27T10:15:47.230+0800","N":"mlog","C":"_test_/main.go:19","M":"SetProjectName: mlog"}
2020-04-27T10:15:47.230+0800	INFO	mlog	_test_/main.go:24	DisableJSONFormat
```



### 新建日志记录器：

```go
package main

import (
	"github.com/blessmylovexy/log/logger"
)

func main() {
	log1 := logger.New()
	log1.Config.SetProjectName("log1")
	log1.ApplyConfig()
	log1.Error("msg log1")

	log2 := logger.New()
	log2.Config.SetProjectName("log2")
	log2.ApplyConfig()
	log2.Error("msg log2")
}
```

```json
{"L":"ERROR","T":"2020-04-27T10:17:01.163+0800","N":"log1","C":"_test_/main.go:11","M":"msg log1"}
{"L":"ERROR","T":"2020-04-27T10:17:01.163+0800","N":"log2","C":"_test_/main.go:16","M":"msg log2"}
```



### 格式化日志记录：

```go
package main

import (
	"github.com/blessmylovexy/log"
)

func main() {
	log.Infow("with fild",
		"url", "www.url.com",
		"duration", 3,
	)
}
```

```json
{"L":"INFO","T":"2020-04-27T10:17:28.377+0800","C":"_test_/main.go:8","M":"with fild","url":"www.url.com","duration":3}
```