package ddns53

import (
	"log"
	"os"
)

// Logger interface for a logger to be used
// with the workers
type Logger interface {
	Debugf(string, ...interface{})
	Debug(...interface{})
	Infof(string, ...interface{})
	Info(...interface{})
	Printf(string, ...interface{})
	Print(...interface{})
	Warnf(string, ...interface{})
	Warn(...interface{})
	Errorf(string, ...interface{})
	Error(...interface{})
	Fatalf(string, ...interface{})
	Fatal(...interface{})
}

type defaultLogger struct {
	original *log.Logger
}

func newLogger() *defaultLogger {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	return &defaultLogger{
		original: logger,
	}
}

func (l *defaultLogger) Debugf(s string, i ...interface{}) {
	l.original.Printf("[Debug] "+s, i...)
}

func (l *defaultLogger) Debug(i ...interface{}) {
	is := make([]interface{}, len(i)+1)
	is[0] = "[Debug] "
	for idx, inter := range i {
		is[idx+1] = inter
	}
	l.original.Print(is...)
}

func (l *defaultLogger) Infof(s string, i ...interface{}) {
	l.original.Printf("[Info] "+s, i...)
}

func (l *defaultLogger) Info(i ...interface{}) {
	is := make([]interface{}, len(i)+1)
	is[0] = "[Info] "
	for idx, inter := range i {
		is[idx+1] = inter
	}
	l.original.Print(is...)
}

func (l *defaultLogger) Printf(s string, i ...interface{}) {
	l.original.Printf(s, i...)
}

func (l *defaultLogger) Print(i ...interface{}) {
	l.original.Print(i...)
}

func (l *defaultLogger) Warnf(s string, i ...interface{}) {
	l.original.Printf("[Warning] "+s, i...)
}

func (l *defaultLogger) Warn(i ...interface{}) {
	is := make([]interface{}, len(i)+1)
	is[0] = "[Warn] "
	for idx, inter := range i {
		is[idx+1] = inter
	}
	l.original.Print(is...)
}

func (l *defaultLogger) Errorf(s string, i ...interface{}) {
	l.original.Printf("[Error] "+s, i...)
}

func (l *defaultLogger) Error(i ...interface{}) {
	is := make([]interface{}, len(i)+1)
	is[0] = "[Error] "
	for idx, inter := range i {
		is[idx+1] = inter
	}
	l.original.Print(is...)
}

func (l *defaultLogger) Fatalf(s string, i ...interface{}) {
	l.original.Fatalf("[Fatal] "+s, i...)
}

func (l *defaultLogger) Fatal(i ...interface{}) {
	is := make([]interface{}, len(i)+1)
	is[0] = "[Fatal] "
	for idx, inter := range i {
		is[idx+1] = inter
	}
	l.original.Print(is...)
}

// NoopLogger is a logger that does not do anything
var NoopLogger = &defaultNoopLogger

// BaseLogger is a logger that prints all to standard out
var BaseLogger = newLogger()

type noopLogger struct{}

var defaultNoopLogger noopLogger

func (l *noopLogger) Debugf(string, ...interface{}) {}
func (l *noopLogger) Debug(...interface{})          {}
func (l *noopLogger) Infof(string, ...interface{})  {}
func (l *noopLogger) Info(...interface{})           {}
func (l *noopLogger) Printf(string, ...interface{}) {}
func (l *noopLogger) Print(...interface{})          {}
func (l *noopLogger) Warnf(string, ...interface{})  {}
func (l *noopLogger) Warn(...interface{})           {}
func (l *noopLogger) Errorf(string, ...interface{}) {}
func (l *noopLogger) Error(...interface{})          {}
func (l *noopLogger) Fatalf(string, ...interface{}) {}
func (l *noopLogger) Fatal(...interface{})          {}
