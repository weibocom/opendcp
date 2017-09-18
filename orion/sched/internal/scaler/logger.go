package scaler

import (
	"fmt"

	"weibo.com/opendcp/orion/service"
)

type logger struct {
	fid int
	nid int
}

func (l *logger) Infof(format string, v ...interface{}) {
	service.Logs.Info(l.fid, l.nid, fmt.Sprintf(format, v...))
}

func (l *logger) Errorf(format string, v ...interface{}) {
	service.Logs.Error(l.fid, l.nid, fmt.Sprintf(format, v...))
}

func (l *logger) Warnf(format string, v ...interface{}) {
	service.Logs.Warn(l.fid, l.nid, fmt.Sprintf(format, v...))
}
