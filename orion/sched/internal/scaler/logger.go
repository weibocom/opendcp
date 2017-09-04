package scaler

import (
	"fmt"

	"weibo.com/opendcp/orion/service"
)

type logger struct {
	fid int
}

func (l *logger) Infof(format string, v ...interface{}) {
	service.Logs.Info(l.fid, fmt.Sprintf(format, v...))
}

func (l *logger) Errorf(format string, v ...interface{}) {
	service.Logs.Error(l.fid, fmt.Sprintf(format, v...))
}

func (l *logger) Warnf(format string, v ...interface{}) {
	service.Logs.Warn(l.fid, fmt.Sprintf(format, v...))
}
