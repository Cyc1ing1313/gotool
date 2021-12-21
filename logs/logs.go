package logs

import (
	"io"
	"log"
	"os"
)

var (
	logger *log.Logger
	level  int
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
)

func init() {
	logger = log.New(os.Stdout, "", log.LstdFlags)
	level = 0
}

func SetOutPut(out io.Writer) {
	logger.SetOutput(out)
}

func Log(format string, args ...interface{}) {
	if level <= INFO {
		logger.Printf(format, args...)
	}
}

func Warn(format string, args ...interface{}) {
	if level <= WARN {
		logger.Printf(format, args...)
	}
}

func Error(prefix string, err error) {
	if level <= ERROR {
		logger.Printf(prefix+": %+v", err)
	}
}
