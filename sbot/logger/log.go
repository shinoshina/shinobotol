package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"time"
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

func handleRawf(raw string, color int) (s string) {
	_, path, line, _ := runtime.Caller(2)

	l := fmt.Sprintf("[Location] %s [Line] %d\n", path, line)
	//f := runtime.FuncForPC(pc).Name()
	//i := l + "[Caller] " + f + "\n"

	k := raw + "\n" + "[Time] " + CurrentTime() + "\n" + l
	s = fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color, k)
	return
}
func handleRawr(color int,v ...interface{}) (s string){
	_, path, line, _ := runtime.Caller(2)
	l := fmt.Sprintf("[Location] %s [Line] %d\n", path, line)
	k := fmt.Sprint(v...) + "\n" + "[Time] " + CurrentTime() + "\n" + l
	s = fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color, k)
	return 
}
func CurrentTime() (ctime string) {

	timeUnix := time.Now().Unix()
	ctime = time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")
	return

}


func Debug(v ...interface{}){
	debugLogger.Print(handleRawr(DEBUGC,v...))
}
func Debugf(format string, v ...interface{}) {
	debugLogger.Printf(handleRawf(format, DEBUGC), v...)
}
func Info(v ...interface{}){
	infoLogger.Print(handleRawr(DEBUGC,v...))
}
func Infof(format string, v ...interface{}) {
	infoLogger.Printf(handleRawf(format, INFOC), v...)
}
func Warn(v ...interface{}){
	warnLogger.Print(handleRawr(DEBUGC,v...))
}
func Warnf(format string, v ...interface{}) {
	warnLogger.Printf(handleRawf(format, WARNC), v...)
}
func Error(v ...interface{}){
	errLogger.Print(handleRawr(DEBUGC,v...))
}
func Errorf(format string, v ...interface{}) {
	errLogger.Printf(handleRawf(format, ERRORC), v...)
}
func Fatal(v ...interface{}){
	fatalLogger.Fatal(handleRawr(DEBUGC,v...))
}
func Fatalf(format string, v ...interface{}) {
	fatalLogger.Fatalf(handleRawf(format, FATALC), v...)
}

