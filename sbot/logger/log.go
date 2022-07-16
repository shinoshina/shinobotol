package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"shinobot/sbot/tick"
)

type level int8

const (
	DEBUG level = iota
	INFO
	WARNNING
	ERROR
	FATAL
)
const (
	logFlag = log.Ldate | log.Lshortfile
)
const (
	DEBUGC = 32 //RED
	INFOC  = 32
	WARNC  = 36
	ERRORC = 33
	FATALC = 31
)

var (
	logFile     io.Writer
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errLogger   *log.Logger
	fatalLogger *log.Logger
)

func init() {

	debugLogger = log.New(os.Stderr, colorConvert(DEBUGC, "[DEBUG] "), 0)
	infoLogger = log.New(os.Stderr, colorConvert(INFOC, "[INFO] "), 0)
	warnLogger = log.New(os.Stderr, colorConvert(WARNC, "[WARN] "), 0)
	errLogger = log.New(os.Stderr, colorConvert(ERRORC, "[ERROR] "), 0)
	fatalLogger = log.New(os.Stderr, colorConvert(FATALC, "[FATAL] "), 0)

}

func colorConvert(color int, raw string) (s string) {

	return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color, raw)
}

func handleRaw(raw string, color int) (s string) {
	_, path, line, _ := runtime.Caller(2)

	l := fmt.Sprintf("[Location] %s [Line] %d\n", path, line)
	//f := runtime.FuncForPC(pc).Name()
	//i := l + "[Caller] " + f + "\n"

	k := raw + "\n" + "[Time] " + tick.CurrentTime() + "\n" + l
	s = fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color, k)
	return

}
func Debugf(format string, v ...interface{}) {
	debugLogger.Printf(handleRaw(format, DEBUGC), v...)
}
func Infof(format string, v ...interface{}) {
	infoLogger.Printf(handleRaw(format, INFOC), v...)
}
func Warnf(format string, v ...interface{}) {
	warnLogger.Printf(handleRaw(format, WARNC), v...)
}
func Errorf(format string, v ...interface{}) {
	errLogger.Printf(handleRaw(format, ERRORC), v...)
}
func Fatalf(format string, v ...interface{}) {
	fatalLogger.Fatalf(handleRaw(format, FATALC), v...)
}
