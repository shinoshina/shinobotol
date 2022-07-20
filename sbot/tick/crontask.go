package tick

import (
	"context"
)

type CronTask struct {
	name  string
	task  func()
	state string

	t          *timer
	ctx        context.Context
	cancel     context.CancelFunc
	autoBooted bool
}

func NewCronTask(name string, task func()) (ct *CronTask) {
	ct = new(CronTask)
	ct.name = name
	ct.task = task
	ct.state = "off"
	ct.ctx, ct.cancel = context.WithCancel(context.Background())
	ct.t = newTimer()
	ct.autoBooted = true
	return
}
func (ct *CronTask) AddRule(raw string) {
	ct.t.fromSchedule(raw)
}
func (ct *CronTask) Start() {
	ct.ctx, ct.cancel = context.WithCancel(context.Background())
	go func() {
		if ct.state == "off" {
			ct.state = "on"
			for {
				select {
				case <-ct.ctx.Done():
					return
				default:
                    ct.t.wait()
					if ct.t.mode == READY {
						// no need to use state ,just use boolean
						// dont wanna mention exceeD
						ct.task()
						ct.t.mode = WAIT_INTERVAL
					}
				}
			}
		} else if ct.state == "on" {
			return
		}
	}()
}
func (ct *CronTask) Stop() {
	if ct.state == "on" {
		ct.state = "off"
		ct.cancel()
	} else if ct.state == "off" {
		return
	}
}
