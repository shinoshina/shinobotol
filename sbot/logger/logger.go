package logger

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type logLevel int

const (
	DEBUG logLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)
const (
	DEBUGC = 32 //RED
	INFOC  = 32
	WARNC  = 36
	ERRORC = 33
	FATALC = 31
)

var Color = [5]int{DEBUGC, INFOC, WARNC, ERRORC, FATALC}
var Prefix = [5]string{"[DEBUG] ", "[INFO] ", "[WARN] ", "[ERROR] ", "[FATAL] "}

var defLogger = &Logger{
	output: os.Stderr,
	format: LshortFile | Ldate,
	fileno: false,
}

const (
	Ldate = 1 << iota
	LshortFile
	LlongFile
)

type Logger struct {
	mu       sync.Mutex
	output   io.Writer
	format   int
	timeout  time.Time
	fileno   bool
	inciseno bool

	buffer           []byte
	currentFileBytes int
}
type Option func(*Logger)

func New(option ...Option) *Logger {
	l := new(Logger)

	for _, o := range option {
		o(l)
	}

	return l
}
func NewDefault() Logger {
	return Logger{
		output: os.Stdout,
		format: LshortFile | Ldate,
		fileno: false,
	}
}
func WithFormat(flag int) Option {
	return func(l *Logger) {
		l.format |= flag
	}
}
func WithOutPut(output io.Writer) Option {
	return func(l *Logger) {
		l.output = output
		if output != os.Stderr || output != os.Stdout {
			l.fileno = true
			l.currentFileBytes = 0
			l.buffer = make([]byte, 0)
		}
	}
}
func WithIncise(timeout time.Time) Option {
	return func(l *Logger) {
		l.timeout = timeout
		l.inciseno = true
	}
}
func (l *Logger) Output(level logLevel, s string) {

	s += "\n"
	var ctime string
	var location string
	if l.format&Ldate != 0 {
		timeUnix := time.Now().Unix()
		ctime = time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05") + " "
	}
	if l.format&(LshortFile|LlongFile) != 0 {
		_, path, line, _ := runtime.Caller(1)
		location = fmt.Sprintf("[%s:%d] ", path, line)
	}
	l.mu.Lock()
	defer l.mu.Unlock()

	if !l.fileno {
		l.output.Write([]byte(colorConvert(Color[level], Prefix[level]+ctime+location+s)))
		return
	} else if level == DEBUG {
		os.Stderr.Write([]byte(colorConvert(Color[level], Prefix[level]+ctime+location+s)))
		return
	}

	l.buffer = append(l.buffer, Prefix[level]...)
	l.buffer = append(l.buffer, ctime...)
	l.buffer = append(l.buffer, location...)
	l.buffer = append(l.buffer, s...)
	fmt.Println("length: ", len(l.buffer))
	if len(l.buffer) >= 2000 {
		l.output.Write(l.buffer)
		l.currentFileBytes += len(l.buffer)
		if l.currentFileBytes >= 250 {
			l.output = newFile(l.output)
			l.currentFileBytes = 0
		}
		l.buffer = l.buffer[:0]
	}
}

func colorConvert(color int, raw string) (s string) {

	return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color, raw)
}

func CurrentTime() (ctime string) {

	timeUnix := time.Now().Unix()
	ctime = time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")
	return
}

func newFile(old io.Writer) io.Writer {
	name := old.(*os.File).Name()

	spoint := strings.Split(name, ".")
	sdash := strings.Split(spoint[0], "-")
	if len(sdash) == 1 {
		name = sdash[0] + "-1." + spoint[1]
	} else {
		seri, err := strconv.ParseInt(sdash[1], 10, 64)
		if err != nil {
			fmt.Println(nil)
		}
		seri++
		seris := strconv.Itoa(int(seri))
		name = sdash[0] + "-" + seris + "." + spoint[1]
	}
	newfile, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	return newfile
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Output(DEBUG, fmt.Sprintf(format, v...))
}
func (l *Logger) Debug(v ...interface{}) {
	l.Output(DEBUG, fmt.Sprint(v...))
}
func (l *Logger) Infof(format string, v ...interface{}) {
	l.Output(INFO, fmt.Sprintf(format, v...))
}
func (l *Logger) Info(v ...interface{}) {
	l.Output(INFO, fmt.Sprint(v...))
}
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Output(WARN, fmt.Sprintf(format, v...))
}
func (l *Logger) Warn(v ...interface{}) {
	l.Output(WARN, fmt.Sprint(v...))
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Output(ERROR, fmt.Sprintf(format, v...))
}
func (l *Logger) Error(v ...interface{}) {
	l.Output(ERROR, fmt.Sprint(v...))
}
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Output(FATAL, fmt.Sprintf(format, v...))
	os.Exit(-1)
}
func (l *Logger) Fatal(v ...interface{}) {
	l.Output(FATAL, fmt.Sprint(v...))
}
