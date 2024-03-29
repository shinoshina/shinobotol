package logger

import (
	"fmt"
	"io"
	"os"
	"path"
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
	DEBUGC = 34 //RED
	INFOC  = 32
	WARNC  = 36
	ERRORC = 33
	FATALC = 31
)

var DefaultColor = [5]int{DEBUGC, INFOC, WARNC, ERRORC, FATALC}
var DefaultPrefix = [5]string{"[DEBUG] ", "[INFO] ", "[WARN] ", "[ERROR] ", "[FATAL] "}

var defLogger = &Logger{
	output:   os.Stderr,
	format:   LshortFile | Ldate,
	fileno:   false,
	inciseno: false,
	colors:   DefaultColor,
	prefixs:  DefaultPrefix,
}

const (
	Ldate = 1 << iota
	LshortFile
	LlongFile
)

type Logger struct {
	mu     sync.Mutex
	output io.Writer
	format int

	size     int
	fileno   bool
	inciseno bool

	buffer           []byte
	currentFileBytes int

	colors  [5]int
	prefixs [5]string
}
type Option func(*Logger)

func New(option ...Option) *Logger {
	l := &Logger{
		fileno:   false,
		inciseno: false,

		format: LlongFile | Ldate,
		output: os.Stderr,

		colors:  [5]int{-1, -1, -1, -1, -1},
		prefixs: [5]string{"", "", "", "", ""},
	}

	for _, o := range option {
		o(l)
	}

	if l.inciseno && l.fileno {
		Fatal("config dismatch")
	}

	return l
}
func NewDefault() Logger {
	return Logger{
		output:   os.Stderr,
		format:   LshortFile | Ldate,
		fileno:   false,
		inciseno: false,
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
			l.buffer = make([]byte, 0)
		}
	}
}
func WithIncise(bits int) Option {
	return func(l *Logger) {
		if bits <= 10240 {
			defLogger.Fatalf("incise deadline to low")
		}
		l.size = bits
		l.inciseno = true
		l.currentFileBytes = 0
	}
}
func WithColorsDefault() Option {
	return func(l *Logger) {
		l.colors = DefaultColor
	}
}
func WithPrefixsDefault() Option {
	return func(l *Logger) {
		l.prefixs = DefaultPrefix
	}
}
func WithColors(colors [5]int) Option {
	return func(l *Logger) {
		l.colors = colors
	}
}
func WithPrefixs(prefixs [5]string) Option {
	return func(l *Logger) {
		l.prefixs = prefixs
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
		_, file, line, _ := runtime.Caller(3)
		if l.format&LshortFile != 0 {
			file = path.Base(file)
		}
		location = fmt.Sprintf("[%s:%d] ", file, line)
	}
	l.mu.Lock()
	defer l.mu.Unlock()

	if !l.fileno {
		l.output.Write([]byte(colorConvert(l.colors[level], l.prefixs[level]+ctime+location+s)))
		return
	} else if level == DEBUG {
		os.Stderr.Write([]byte(colorConvert(l.colors[level], l.prefixs[level]+ctime+location+s)))
		return
	}

	l.buffer = append(l.buffer, l.prefixs[level]...)
	l.buffer = append(l.buffer, ctime...)
	l.buffer = append(l.buffer, location...)
	l.buffer = append(l.buffer, s...)
	fmt.Println("length: ", len(l.buffer))
	if len(l.buffer) >= 2000 {
		l.output.Write(l.buffer)
		if l.inciseno {
			l.currentFileBytes += len(l.buffer)
			if l.currentFileBytes >= l.size {
				l.output = newFile(l.output)
				l.currentFileBytes = 0
			}
		}
		l.buffer = l.buffer[:0]
	}
}

func colorConvert(color int, raw string) (s string) {
	if color != -1 {
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color, raw)
	}else{
		return raw
	}
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
