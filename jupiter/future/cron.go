package future

import (
	"github.com/astaxie/beego"
	"time"
	"regexp"
	"errors"
	"strconv"
)

const DETAIL_INTERVAL = "00时00分60分钟"

type Handle func() error

type CronDetail struct {
	TickerDuration 	time.Duration
	Ticker  	*time.Ticker
	Stop           	bool
}

type CronFuture struct {
	Description string
	Handle      Handle
	Schedule    string
	CronDetail
}

func (cf *CronFuture) Run() error {
	beego.Info("Cron task [", cf.Description, "] begin...")
	for {
		<- cf.Ticker.C
		if !cf.Stop{
			if err := cf.Handle(); err != nil {
				return err
			}
			cf.UpdateTicker()
			continue
		}
		cf.Ticker.Stop()
		break
	}
	return nil
}

func (cf *CronFuture) Success() {
	beego.Info("Cron task execute success!")
}

func (cf *CronFuture) Failure(err error) {
	beego.Error("Cron task execute failed: ", err)
	Exec.Submit(cf)
}

func (cf *CronFuture) ShutDown() {
	cf.Stop = true
	beego.Warn("Cron task [",cf.Description,"] stop!")
}

func (cf *CronFuture)UpdateTicker()  {
	ticker, err := NewTicker(cf.Schedule)
	if err != nil {
		beego.Error("Cron task update ticker err: ", err)
		return
	}
	cf.Ticker = ticker
}

//解析定时任务的参数
func ParseScheduleTime(scheduleTime string) ([]int, error) {
	time_arr := make([]int, 10)
	var err error
	timeNumber := regexp.MustCompile("[0-9]+")
	arr := timeNumber.FindAllString(scheduleTime, -1)
	if len(arr) != 3 {
		return time_arr, errors.New("Parse config err!")
	}
	for i:=0; i<len(arr); i++{
		time_arr[i], err = strconv.Atoi(arr[i])
		if err != nil {
			return time_arr, err
		}
	}
	return time_arr, nil
}

func NewCronbFuture(detail string, scheduleTime string, handle Handle) *CronFuture {
	ticker, err := NewTicker(scheduleTime)
	if err != nil {
		beego.Error("Create ticker err: ", err)
		return nil
	}

	cron := &CronFuture{
		Description: 	detail,
		Handle: 	handle,
		Schedule: 	scheduleTime,
	}
	cron.Stop = false
	cron.Ticker = ticker
	return cron
}

func getNextTickDuration(hour, minute,interval int ) time.Duration {
	now := time.Now()
	nextTick := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, time.Local)
	for ;!nextTick.After(now); {  		//如果任务启动时间早于当前时间，更新启动时间
		nextTick = nextTick.Add(time.Duration(interval)*time.Minute)
	}
	beego.Info("The cron task execute next time is ", nextTick)
	return nextTick.Sub(time.Now())
}

func NewTicker(scheduleTime string) (*time.Ticker, error) {
	arr, err := ParseScheduleTime(scheduleTime)
	if err != nil {
		return nil, err
	}
	ticker := time.NewTicker(getNextTickDuration(arr[0], arr[1], arr[2]))
	return ticker, nil
}


