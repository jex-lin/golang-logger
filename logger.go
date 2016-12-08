package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

type Level int

const (
	DEBUG Level = iota // 0
	INFO
	NOTICE
	WARN
	ERROR
	CRITICAL
)

type Log struct {
	Logger  *log.Logger
	Level   Level    // Filter level
	Trigger struct { // Trigger func when level reached.
		Level Level
		Do    func()
	}
}

type CallInfo struct {
	PkgName  string
	FileName string
	FuncName string
	Line     int
}

func New(out io.Writer) *Log {
	logger := log.New(out, "", 0)
	return &Log{Logger: logger}
}

// New and create log file
func NewLogFile(filename string) *Log {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Failed to create log file, error:", err)
		os.Exit(1)
	}
	logger := log.New(f, "", 0)
	return &Log{Logger: logger}
}

func (l *Log) SetLevel(str string) *Log {
	l.Level = StrToLevel(str)
	return l
}

func (l *Log) SetTrigger(str string, do func()) *Log {
	l.Trigger.Level = StrToLevel(str)
	l.Trigger.Do = do
	return l
}

func (l *Log) Debug(v ...interface{}) {
	info := GetCallInfo()
	l.FormatPrint(DEBUG, info, fmt.Sprint(v...))
}

func (l *Log) Debugf(format string, v ...interface{}) {
	info := GetCallInfo()
	l.FormatPrint(DEBUG, info, fmt.Sprintf(format, v...))
}

func (l *Log) Info(v ...interface{}) {
	info := GetCallInfo()
	l.FormatPrint(INFO, info, fmt.Sprint(v...))
}

func (l *Log) Infof(format string, v ...interface{}) {
	info := GetCallInfo()
	l.FormatPrint(INFO, info, fmt.Sprintf(format, v...))
}

func (l *Log) Notice(v ...interface{}) {
	info := GetCallInfo()
	l.FormatPrint(NOTICE, info, fmt.Sprint(v...))
}

func (l *Log) Noticef(format string, v ...interface{}) {
	info := GetCallInfo()
	l.FormatPrint(NOTICE, info, fmt.Sprintf(format, v...))
}

func (l *Log) Warn(v ...interface{}) {
	info := GetCallInfo()
	l.FormatPrint(WARN, info, fmt.Sprint(v...))
}

func (l *Log) Warnf(format string, v ...interface{}) {
	info := GetCallInfo()
	l.FormatPrint(WARN, info, fmt.Sprintf(format, v...))
}

func (l *Log) Error(v ...interface{}) {
	info := GetCallInfo()
	l.FormatPrint(ERROR, info, fmt.Sprint(v...))
}

func (l *Log) Errorf(format string, v ...interface{}) {
	info := GetCallInfo()
	l.FormatPrint(ERROR, info, fmt.Sprintf(format, v...))
}

func (l *Log) Critical(v ...interface{}) {
	info := GetCallInfo()
	l.FormatPrint(CRITICAL, info, fmt.Sprint(v...))
}

func (l *Log) Criticalf(format string, v ...interface{}) {
	info := GetCallInfo()
	l.FormatPrint(CRITICAL, info, fmt.Sprintf(format, v...))
}

func (l *Log) Format(level Level, info *CallInfo) (str string) {
	level_str := LevelToStr(level)
	str = fmt.Sprintf("%s [%s] %s %s(:%d) >", time.Now().Format("2006/01/02 15:04:05"), level_str, info.PkgName, info.FuncName, info.Line)
	return
}

func (l *Log) FormatPrint(level Level, info *CallInfo, v ...interface{}) {
	str := l.Format(level, info)

	if level >= l.Level {
		l.Logger.Println(str, v[0])
	}

	// Trigger Do (if existed) by specific level
	if level >= l.Trigger.Level && l.Trigger.Do != nil {
		l.Trigger.Do()
	}
}

func LevelToStr(level Level) string {
	switch level {
	case DEBUG:
		return "Debug"
	case INFO:
		return "Info"
	case NOTICE:
		return "Notice"
	case WARN:
		return "Warn"
	case ERROR:
		return "Error"
	case CRITICAL:
		return "Critical"
	default:
		return "Unknown"
	}
}

func StrToLevel(str string) Level {
	switch strings.ToLower(str) {
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "notice":
		return NOTICE
	case "warn":
		return WARN
	case "error":
		return ERROR
	case "critical":
		return CRITICAL
	default:
		return 100 // Return a big number to make level not triggered.
	}
}

func GetCallInfo() *CallInfo {
	pc, file, line, _ := runtime.Caller(2)
	_, fileName := path.Split(file)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	packageName := ""
	funcName := parts[pl-1]

	if parts[pl-2][0] == '(' {
		// If you want to get more detail about caller. e.g. (t *MyStruct).Do
		// Just enable the code below; otherwise, you will only get function name. e.g. Do
		// funcName = parts[pl-2] + "." + funcName
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}

	// Get last element of path.   e.g. myapp/utility/config  => common
	tmp := strings.Split(packageName, "/")
	packageName = tmp[len(tmp)-1]

	return &CallInfo{
		PkgName:  packageName,
		FileName: fileName,
		FuncName: funcName,
		Line:     line,
	}
}
