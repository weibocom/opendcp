package cron

import (
	"context"
	"sync"

	"weibo.com/opendcp/orion/sched/internal/scaler"
)

type cronJob struct {
	ctx  context.Context
	cfgs configs
	id   int
	idx  int
}

func (j *cronJob) Run() {
	j.ctx.Value("wg").(*sync.WaitGroup).Add(1)
	go scaler.Scale(j.ctx, j.cfgs, j.id, j.idx)
}
