package client

type Logger interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
}

type noopLogger struct {
}

func (n noopLogger) Warnf(format string, v ...interface{}) {

}

func (n noopLogger) Infof(format string, args ...interface{}) {

}

func (n noopLogger) Errorf(format string, args ...interface{}) {

}

func (n noopLogger) Debugf(format string, args ...interface{}) {

}

// NewNoop returns a logger that does nothing.
func NewNoop() Logger {
	return &noopLogger{}
}
