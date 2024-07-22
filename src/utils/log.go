package utils

type LogInfoFunc func(format string, v ...any)
type LogErrorFunc func(format string, v ...any)

type LogWriter struct {
	Info  LogInfoFunc
	Error LogErrorFunc
}

var instance LogWriter

func Log() *LogWriter {
	return &instance
}

func (m *LogWriter) Init(info LogInfoFunc, e LogErrorFunc) {
	m.Info = info
	m.Error = e
}
