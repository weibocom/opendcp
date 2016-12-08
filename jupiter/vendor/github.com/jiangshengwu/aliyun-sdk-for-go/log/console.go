package log

import (
	"fmt"
	"os"
)

type ConsoleLogger struct {
}

func (t ConsoleLogger) log(args ...interface{}) {
	fmt.Println(args...)
}

func (t ConsoleLogger) logf(fmtString string, args ...interface{}) {
	fmt.Printf(fmtString+"\n", args...)
}

func (t ConsoleLogger) err(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}

func (t ConsoleLogger) errf(fmtString string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, fmtString+"\n", args...)
}

func (t ConsoleLogger) Debug(args ...interface{}) {
	if isDebug() {
		t.log(args...)
	}
}

func (t ConsoleLogger) Debugf(fmtString string, args ...interface{}) {
	if isDebug() {
		t.logf(fmtString, args...)
	}
}

func (t ConsoleLogger) Error(args ...interface{}) {
	t.err(args...)
}

func (t ConsoleLogger) Errorf(fmtString string, args ...interface{}) {
	t.errf(fmtString, args...)
}

func (t ConsoleLogger) Info(args ...interface{}) {
	t.log(args...)
}

func (t ConsoleLogger) Infof(fmtString string, args ...interface{}) {
	t.logf(fmtString, args...)
}

func (t ConsoleLogger) Fatal(args ...interface{}) {
	t.err(args...)
	os.Exit(1)
}

func (t ConsoleLogger) Fatalf(fmtString string, args ...interface{}) {
	t.errf(fmtString, args...)
	os.Exit(1)
}

func (t ConsoleLogger) Print(args ...interface{}) {
	t.log(args...)
}

func (t ConsoleLogger) Printf(fmtString string, args ...interface{}) {
	t.logf(fmtString, args...)
}

func (t ConsoleLogger) Warn(args ...interface{}) {
	t.log(args...)
}

func (t ConsoleLogger) Warnf(fmtString string, args ...interface{}) {
	t.logf(fmtString, args...)
}
