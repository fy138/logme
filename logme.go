package logme

import (
	"fmt"
	"log"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

type Logme struct {
	Logfile  string
	Logger   *log.Logger
	Loglevel int
}

const (
	colorBlack = uint8(iota + 90)
	colorRed
	colorGreen
	colorYellow
	colorBlue
	colorBagenta
	colorCyan
	colorWhite
)
const (
	DEBUG int = 1
	INFO  int = 2
	WARN  int = 3
	ERROR int = 4
	FATAL int = 5
)

func NewLogMe(logfile string, loglevel int) *Logme {
	lm := &Logme{}
	lm.Logfile = logfile
	lm.Loglevel = loglevel

	if len(logfile) == 0 {
		lm.Logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile|log.LUTC)
	} else {
		writer, _ := rotatelogs.New(
			lm.Logfile+".%Y%m%d%H%M",
			rotatelogs.WithLinkName(lm.Logfile),
			rotatelogs.WithMaxAge(time.Duration(3600*24)*time.Second),
			rotatelogs.WithRotationTime(time.Duration(3600*24)*time.Second),
		)
		lm.Logger = log.New(writer, "", log.LstdFlags|log.Lshortfile|log.LUTC)
	}
	return lm
}
func (lm *Logme) Debug(v ...interface{}) {
	if lm.Loglevel > 1 {
		return //不显示
	}
	lm.Logger.SetPrefix(setColor(DEBUG))
	//lm.Logger.Output(2, fmt.Sprint(v...))
	lm.Logger.Output(2, fmt.Sprint(v...))
}
func (lm *Logme) Debugf(format string, v ...interface{}) {
	if lm.Loglevel > 1 {
		return //不显示
	}
	lm.Logger.SetPrefix(setColor(DEBUG))
	//lm.Logger.Output(2, fmt.Sprint(v...))
	lm.Logger.Output(2, fmt.Sprintf(format, v...))
}
func (lm *Logme) Info(v ...interface{}) {
	if lm.Loglevel > 2 {
		return //不显示
	}
	lm.Logger.SetPrefix(setColor(INFO))
	//lm.Logger.Output(2, fmt.Sprint(v...))
	lm.Logger.Output(2, fmt.Sprint(v...))
}
func (lm *Logme) Infof(format string, v ...interface{}) {
	if lm.Loglevel > 2 {
		return //不显示
	}
	lm.Logger.SetPrefix(setColor(INFO))
	lm.Logger.Output(2, fmt.Sprintf(format, v...))
}
func (lm *Logme) Warn(v ...interface{}) {
	if lm.Loglevel > 3 {
		return //不显示
	}
	lm.Logger.SetPrefix(setColor(WARN))
	lm.Logger.Output(2, fmt.Sprint(v...))
}
func (lm *Logme) Warnf(format string, v ...interface{}) {
	if lm.Loglevel > 3 {
		return //不显示
	}
	lm.Logger.SetPrefix(setColor(WARN))
	lm.Logger.Output(2, fmt.Sprintf(format, v...))
}
func (lm *Logme) Error(v ...interface{}) {
	if lm.Loglevel > 4 {
		return //不显示
	}
	lm.Logger.SetPrefix(setColor(ERROR))
	lm.Logger.Output(2, fmt.Sprint(v...))
}
func (lm *Logme) Errorf(format string, v ...interface{}) {
	if lm.Loglevel > 4 {
		return //不显示
	}
	lm.Logger.SetPrefix(setColor(ERROR))
	lm.Logger.Output(2, fmt.Sprintf(format, v...))
}
func (lm *Logme) Fatal(v ...interface{}) {
	if lm.Loglevel > 4 {
		return //不显示
	}
	lm.Logger.SetPrefix(setColor(FATAL))
	lm.Logger.Fatal(v...)
}
func (lm *Logme) Fatalf(format string, v ...interface{}) {
	if lm.Loglevel > 4 {
		return //不显示
	}
	lm.Logger.SetPrefix(setColor(FATAL))
	lm.Logger.Fatalf(format, v...)
}
func setColor(level int) string {
	switch level {
	case ERROR:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorRed, "[ERROR]")
	case WARN:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorYellow, "[WARN]")
	case INFO:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorGreen, "[INFO]")
	case DEBUG:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorBlue, "[DEBUG]")
	default:
		return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorBlack, "[FATAL]")
	}

}
