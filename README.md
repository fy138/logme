# logme
a simple golang log base on golang  log module

#Example
```go
//some code 
Logger = logme.NewLogMe("", logme.DEBUG)
or 
Logger = logme.NewLogMe("./mylog.log", logme.INFO)

//use
Logger.Info("abc")
//[INFO]2020/10/23 01:44:35 main.go:1: abc
Logger.Errorf("you have error %s","abc")
//[ERROR]2020/10/23 01:44:35 main.go:1: you have error abc

//all mothod
Logger.Debug
Logger.Debugf
Logger.Info
Logger.Infof
Logger.Warn
Logger.Warnf
Logger.Error
Logger.Errorf
Logger.Fatal
Logger.Fatalf
```
