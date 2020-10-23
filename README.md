# logme
a simple golang log base on golang  log module

# Example
```go
package main

import (
	"github.com/fy138/logme"
)

func main() {
	logger := logme.NewLogMe("", logme.DEBUG)

	logger.Debug("show debug")
	logger.Debugf("this is %s", "debug")

	logger.Info("my info")
	logger.Infof("this is %s", "info")

	logger.Warn("show warn")
	logger.Warnf("this is %s", "warn")

	logger.Error("show error")
	logger.Errorf("this is %s", "error")

	logger.Fatal("show fatal")
	logger.Fatalf("this is %s", "fatal")
}
```
![image](https://raw.githubusercontent.com/fy138/logme/main/1.png)
