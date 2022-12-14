/*
Simple logrus implementation.
*/
package logging

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

// Custom hook for logrus.
type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}

	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}

	return err
}

// Return logger levels.
func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

var entry *logrus.Entry

type Logger struct {
	*logrus.Entry
}

// Return new logger instanec.
func GetLogger() *Logger {
	return &Logger{entry}
}

// Return logger with parameters.
func (log *Logger) GetLoggerWithField(k string, v interface{}) *Logger {
	return &Logger{log.WithField(k, v)}
}

// Logger constructor.
func init() {
	log := logrus.New()
	log.SetReportCaller(true)
	log.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	err := os.MkdirAll("logs", 0644)
	if err != nil {
		panic(err)
	}

	allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		panic(err)
	}

	log.SetOutput(io.Discard)

	log.AddHook(&writerHook{
		Writer:    []io.Writer{allFile, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	wrt := io.MultiWriter(os.Stdout, allFile)
	log.SetOutput(wrt)

	log.SetLevel(logrus.TraceLevel)

	entry = logrus.NewEntry(log)
}
