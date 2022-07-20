package tick

import (
	"context"
)


type CronTask struct {
	name   string
	task   func()
	state  string

	rules   []string
	ctx    context.Context
	cancel context.CancelFunc
	autoBooted bool
}

func NewCronTask(name string, task func()) (ct *CronTask) {
	ct = new(CronTask)
	ct.name = name
	ct.task = task
	ct.state = "off"
	ct.ctx, ct.cancel = context.WithCancel(context.Background())

	ct.autoBooted = true
	ct.rules = make([]string, 1)
	return
}
func(ct *CronTask) AddRule(rule string){
	ct.rules = append(ct.rules, rule)
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
					ct.task()
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
