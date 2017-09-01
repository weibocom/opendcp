package scaler

import (
	"context"
	"fmt"
	"math"
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

func dependNoticeGoroutine(ctx context.Context, expand bool, should int, dependc chan *dependNotice, ctrls map[int]*dependCtrl) {
	wg := ctx.Value("wg").(*sync.WaitGroup)
	defer wg.Done()

	var (
		//depnum  = int(math.MaxInt32)
		timeout = time.After(30 * time.Minute)
	)

	//defer close(dependc)

	if len(ctrls) == 0 {
		for i := 0; i < should; i++ {
			dependc <- &dependNotice{num: 1}
		}
		return
	}

	isDepenSuccess := true
	for _, ctrl := range ctrls {
		successDependNum := 0
		for i := 0; i < ctrl.num; i++ {
			select {
			case n := <-ctrl.dependc:
				successDependNum += n.num
			//num := int(math.Ceil(float64(n.num) / ctrl.ratio))
			//beego.Info(fmt.Sprintf("expand depend recv pool(%d) (%d->%d)", n.pid, n.num, num))
			//if num < depnum {
			//	depnum = num
			//}
			case <-timeout:
				beego.Info("wait expand depend signal timeout")
			}
		}
		if ctrl.num != successDependNum{
			isDepenSuccess = false
		}

	}

	for i := 0; i < should; i++ {
		if(isDepenSuccess) {
			dependc <- &dependNotice{num: 1}
		}else {
			dependc <- &dependNotice{num: 0}
		}

	}


	//if expand {
	//	dependExpandGoroutine(ctx, should, dependc, ctrls)
	//} else {
	//	dependShrinkGoroutine(ctx, should, dependc, ctrls)
	//}
}

func dependExpandGoroutine(ctx context.Context, should int, dependc chan *dependNotice, ctrls map[int]*dependCtrl) {
	// just make it runnable
	var (
		//depnum  = int(math.MaxInt32)
		timeout = time.After(30 * time.Minute)
	)

	//defer close(dependc)

	if len(ctrls) == 0 {
		for i := 0; i < should; i++ {
			dependc <- &dependNotice{num: 1}
		}
		return
	}

	isDepenSuccess := true
	for _, ctrl := range ctrls {
		successDependNum := 0
		for i := 0; i < ctrl.num; i++ {
			select {
			case n := <-ctrl.dependc:
				successDependNum += n.num
				//num := int(math.Ceil(float64(n.num) / ctrl.ratio))
				//beego.Info(fmt.Sprintf("expand depend recv pool(%d) (%d->%d)", n.pid, n.num, num))
				//if num < depnum {
				//	depnum = num
				//}
			case <-timeout:
				beego.Info("wait expand depend signal timeout")
			}
		}
		if ctrl.num != successDependNum{
			isDepenSuccess = false
		}

	}

	for i := 0; i < should; i++ {
		if(isDepenSuccess) {
			dependc <- &dependNotice{num: 1}
		}else {
			dependc <- &dependNotice{num: 0}
		}

	}
	//}
	//dependc <- &dependNotice{num: depnum}
}

func dependShrinkGoroutine(ctx context.Context, should int, dependc chan *dependNotice, ctrls map[int]*dependCtrl) {
	// just make it runnable
	//defer func() {
	//	for _, c := range ctrls {
	//		close(c.dependc)
	//	}
	//}()

	select {
	case notice := <-dependc:
		for _, ctrl := range ctrls {
			num := int(math.Ceil(float64(notice.num) * ctrl.ratio))
			beego.Info(fmt.Sprintf("pool(%d) shrink notice (%d->%d)", notice.pid, notice.num, num))
			ctrl.dependc <- &dependNotice{pid: notice.pid, num: num}
		}
	case <-time.After(30 * time.Minute):
		beego.Info("wait shrink depend signal timeout")
	}
}
