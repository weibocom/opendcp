package log

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Logger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})

	Error(...interface{})
	Errorf(string, ...interface{})

	Info(...interface{})
	Infof(string, ...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})

	Print(...interface{})
	Printf(string, ...interface{})

	Warn(...interface{})
	Warnf(string, ...interface{})
}

var (
	l = ConsoleLogger{}
)

func isDebug() bool {
	debugEnv := os.Getenv("DEBUG")
	if strings.EqualFold(debugEnv, "true") {
		showDebug, err := strconv.ParseBool(debugEnv)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing boolean value from DEBUG: %s\n", err)
			return false
		}
		return showDebug
	}
	return false
}

func Debug(args ...interface{}) {
	l.Debug(args...)
}

func Debugf(fmtString string, args ...interface{}) {
	l.Debugf(fmtString, args...)
}

func Error(args ...interface{}) {
	l.Error(args...)
}

func Errorf(fmtString string, args ...interface{}) {
	l.Errorf(fmtString, args...)
}

func Info(args ...interface{}) {
	l.Info(args...)
}

func Infof(fmtString string, args ...interface{}) {
	l.Infof(fmtString, args...)
}

func Fatal(args ...interface{}) {
	l.Fatal(args...)
}

func Fatalf(fmtString string, args ...interface{}) {
	l.Fatalf(fmtString, args...)
}

func Print(args ...interface{}) {
	l.Print(args...)
}

func Printf(fmtString string, args ...interface{}) {
	l.Printf(fmtString, args...)
}

func Warn(args ...interface{}) {
	l.Warn(args...)
}

func Warnf(fmtString string, args ...interface{}) {
	l.Warnf(fmtString, args...)
}
