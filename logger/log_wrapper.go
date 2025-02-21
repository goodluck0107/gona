package logger

type logw struct {
}

func (l *logw) StartUp(v ...interface{}) {
	StartUp(v...)
}
func (l *logw) Info(v ...interface{}) {
	Info(v...)
}
func (l *logw) Debug(v ...interface{}) {
	Debug(v...)
}
func (l *logw) Warn(v ...interface{}) {
	Warn(v...)
}
func (l *logw) Error(v ...interface{}) {
	Error(v...)
}
