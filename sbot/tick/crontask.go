package tick

import (
	"context"
	"fmt"
)

type CronTask struct {
	name   string
	task   func(*Timer)
	state  string
	ctx    context.Context
	cancel context.CancelFunc

	timer *Timer
}
func NewCronTask(name string, task func( t *Timer)) (ct *CronTask) {
	ct = new(CronTask)
	ct.name = name
	ct.task = task
	ct.state = "off"
	ct.ctx, ct.cancel = context.WithCancel(context.Background())
	return
}
func (ct *CronTask) Start() {
	ct.ctx, ct.cancel = context.WithCancel(context.Background())
	go func() {
		if ct.state == "off" {
			ct.state = "on"
			for {
				select {
				case <-ct.ctx.Done():
					fmt.Println("直接结束啦？")
					return
				default:
					fmt.Println("没有啊")
					ct.task(ct.timer)
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
