package logger

type Logger interface {
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
	Errorf(format string, args ...interface{})
	Error(args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
}
