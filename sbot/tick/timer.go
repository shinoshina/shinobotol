package tick

import (
	"time"
)


type Timer struct{
	timer time.Timer
	schduledt int64
	Counter int
	ptfrom *CronTask
	set bool
}

type Time struct{
	Date string
	Weekday string
	Clock string
	Increament *Time
}

func CurrentTime() (ct string) {

	tunix := time.Now().Unix()
	ct = time.Unix(tunix, 0).Format("2006-01-02 15:04:05")
	return

}

func (t *Timer) WaitUntil(tformat string){

	// ct := time.Unix(time.Now().Unix(),0).Format("xxxx-xx-xx 15:04:05")

}
func (t *Timer) Wait(duration string){

}
func(t *Timer) Stop(){
	t.ptfrom.Stop()
}

