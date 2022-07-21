package tick

import (
	"log"

	"github.com/robfig/cron/v3"
)

type CronTask struct {
	name  string
	task  func()
	state string

	t          *timer
	c          *cron.Cron
	autoBooted bool
}

func NewCronTask(name string, rule string, task func()) (ct *CronTask) {
	ct = new(CronTask)

	ct.c = cron.New(cron.WithSeconds(),cron.WithLogger(cron.VerbosePrintfLogger(log.Default())))
	ct.c.AddFunc(rule,task)
	ct.name = name
	ct.state = "off"
	ct.autoBooted = true
	return
}
func (ct *CronTask) AddRule(raw string) {
	ct.t.fromSchedule(raw)
}
func (ct *CronTask) Start(){
	if ct.state == "off"{
		ct.state = "on"
		ct.c.Start()
	}
}
func (ct *CronTask) Stop(){
	if ct.state == "on"{
		ct.state = "off"
		ct.c.Stop()
	}
}
// func (ct *CronTask) Start() {
// 	ct.ctx, ct.cancel = context.WithCancel(context.Background())
// 	go func() {
// 		if ct.state == "off" {
// 			ct.state = "on"
// 			for {
// 				select {
// 				case <-ct.ctx.Done():
// 					return
// 				default:
// 					ct.t.wait()

// 					// no need to use state ,just use boolean
// 					// dont wanna mention exceeD
// 					ct.task()
// 				}
// 			}
// 		} else if ct.state == "on" {
// 			return
// 		}
// 	}()
// }
// func (ct *CronTask) Stop() {
// 	if ct.state == "on" {
// 		ct.state = "off"
// 		ct.cancel()
// 	} else if ct.state == "off" {
// 		return
// 	}
// }
