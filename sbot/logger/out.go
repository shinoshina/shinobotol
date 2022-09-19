package logger


func Debugf(format string, v ...interface{}) {
	defLogger.Debugf(format,v...)
}
func Debug(v... interface{}){
	defLogger.Debug(v...)
}

func Infof(format string, v ...interface{}) {
	defLogger.Infof(format,v...)
}
func Info(v... interface{}){
	defLogger.Info(v...)
}

func Warnf(format string, v ...interface{}) {
	defLogger.Warnf(format,v...)
}
func Warn(v... interface{}){
	defLogger.Warn(v...)
}

func Errorf(format string, v ...interface{}) {
	defLogger.Errorf(format,v...)
}
func Error(v... interface{}){
	defLogger.Error(v...)
}

func Fatalf(format string, v ...interface{}) {
	defLogger.Fatalf(format,v...)
}
func Fatal(v... interface{}){
	defLogger.Fatal(v...)
}
