// loggerinterface.go - содержит описание интерфейса Logger, данному интерфейсу должно
// соответствовать любое средство логирования применяемое в приложении
package logger

type Logger interface {
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
	Errorf(format string, args ...interface{})
	Error(args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
}
