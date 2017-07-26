package scaler

import (
	"fmt"

	"weibo.com/opendcp/orion/service"
	"weibo.com/opendcp/orion/utils"
)

func newLogger(fid int) *logger {
	return &logger{
		fid: fid,
		cid: utils.GetCorrelationId(fid, 0),
	}
}

type logger struct {
	fid int
	cid string
}

func (l *logger) Infof(format string, v ...interface{}) {
	service.Logs.Info(l.fid, 0, l.cid, fmt.Sprintf(format, v...))
}
