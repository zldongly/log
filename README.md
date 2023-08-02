# log
日志接口  
适配 zap.SugaredLogger、testing.T

## Example
```go
package main

import (
    "github.com/zldongly/log/logger"
    "go.uber.org/zap"
)

func main() {
    logger, _ := zap.NewDevelopment()
    sugar := logger.Sugar()
	
    log := logger.NewZapSuger(sugar)
    log = log.With("version", "v0.1.0")
    log.Info("Info")
}
```

## bug
* testing.T stack信息