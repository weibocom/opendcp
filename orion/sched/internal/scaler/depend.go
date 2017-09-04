package scaler

import (
	"context"
	"sync"
	"time"

	"github.com/astaxie/beego"
)

type dependNotice struct {
	pid int
	num int
}

type dependCtrl struct {
	base, num int
	elastic   int
	ratio     float64
	dependc   chan *dependNotice
}

func dependNoticeGoroutine(ctx context.Context, should int, dependc chan *dependNotice, ctrls map[int]*dependCtrl) {
	wg := ctx.Value("wg").(*sync.WaitGroup)
	defer wg.Done()

	var (
		isDepenSuccess = true
		timeout        = time.After(30 * time.Minute)
	)
	defer func() {
		for _, ctrl := range ctrls {
			close(ctrl.dependc)
		}
	}()

	if len(ctrls) == 0 {
		for i := 0; i < should; i++ {
			dependc <- &dependNotice{num: 1}
		}
		return
	}
	for _, ctrl := range ctrls {
		successDependNum := 0
		for i := 0; i < ctrl.num; i++ {
			select {
			case n := <-ctrl.dependc:
				successDependNum += n.num
			case <-timeout:
				beego.Error("wait expand depend signal timeout!")
			}
		}
		if ctrl.num != successDependNum {
			isDepenSuccess = false
		}
	}
	for i := 0; i < should; i++ {
		if isDepenSuccess {
			dependc <- &dependNotice{num: 1}
		} else {
			dependc <- &dependNotice{num: 0}
		}
	}
}
