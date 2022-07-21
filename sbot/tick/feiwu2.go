package tick

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	format = "2006-01-02 15:04:05"
)

const (
	INITIAL        = 0
	NO_SCHEDULED   = 1
	WAIT_SCHEDULED = 2
	WAIT_INTERVAL  = 3
	READY          = 4
)

type timer struct {
	pit        []string
	pitOrigin  string
	pitTime    time.Time
	fixd       duration
	minCheckPt int
	mode       int
	rtimer     *time.Timer
	noInterval bool
	noSchedule bool
}

type duration struct {
	sec  int
	min  int
	hour int
}

func CurrentTime() (ct string) {

	s := time.Time{}
	s.Sub(time.Now())
	tunix := time.Now().Unix()
	ct = time.Unix(tunix, 0).Format("2006-01-02 15:04:05")

	return
}
func newTimer() (t *timer) {
	t = new(timer)
	t.pit = make([]string, 3)
	t.mode = WAIT_SCHEDULED
	t.minCheckPt = 3
	return
}

//好屎好似好事好诗
func (t *timer) fromSchedule(raw string) {
	fu, iu := handleRaw(raw)
	for i := 0; i < 3; i++ {
		t.pit[i] = fu[i]
	}
	t.pitOrigin = strings.Split(time.Now().Local().Format(format), " ")[0] + " " + t.pit[2] + ":" + t.pit[1] + ":" + t.pit[0]
	if iu[0]+iu[1]+iu[2] == 0 {
		t.noInterval = true
	} else {
		t.fixd.sec = iu[0]
		t.fixd.min = iu[1]
		t.fixd.hour = iu[2]
		t.noInterval = false
	}
	if t.pit[2] == "*" {
		t.minCheckPt = 2
	}
	if t.pit[1] == "*" {
		t.minCheckPt = 1
	}
	if t.pit[0] == "*" {
		if t.minCheckPt == 1 {
			t.noSchedule = true
		}
	} else {
		t.noSchedule = false
	}
}
func (t *timer) waitSchedule() {

	t.resetPitTimeRaw()
	if t.pitTime.After(time.Now().Local()) {
		fmt.Println("waiting schedule :")
		fmt.Println(t.pitTime.Sub(time.Now().Local()))
		t.rtimer = time.NewTimer(t.pitTime.Sub(time.Now().Local()))
		<-t.rtimer.C
	}

}
func (t *timer) waitInterval() {

	if !t.noInterval && !t.noSchedule {
		t.resetPitTimeComple()
		fmt.Println("judge current time for interval job:")
		fmt.Println(t.pitTime.Local())
		if !t.pitTime.After(time.Now().Local().Add(time.Duration(t.fixd.hour)*time.Hour + time.Duration(t.fixd.min)*
			time.Minute + time.Duration(t.fixd.sec)*time.Second)) {
			fmt.Println("no free execute chance :")
			fmt.Println(t.pitTime.Sub(time.Now().Local()))
			t.rtimer = time.NewTimer(t.pitTime.Sub(time.Now().Local()))
			<-t.rtimer.C
		}
	}
	t.rtimer = time.NewTimer(time.Duration(t.fixd.hour)*time.Hour + time.Duration(t.fixd.min)*
		time.Minute + time.Duration(t.fixd.sec)*time.Second)
	<-t.rtimer.C
}

//初次鉴定:屎
func (t *timer) wait() {

	if !t.noSchedule {
		t.waitSchedule()
	}
	if !t.noInterval {
		t.waitInterval()
	}
}

func fromStringValue(tick int) string {
	if tick < 10 {
		return "0" + strconv.Itoa(tick)
	} else {
		return strconv.Itoa(tick)
	}
}

func (t *timer) resetPitTimeComple() {
	if t.minCheckPt == 3 {
		t.pitTime = t.pitTime.Add(24*time.Hour - time.Duration(t.pitTime.Hour()))
		fmt.Println("nihao")
		fmt.Println(t.pitTime)
	} else if t.minCheckPt == 2 {
		t.pitTime = t.pitTime.Add(60*time.Minute - time.Duration(t.pitTime.Minute()))
	} else if t.minCheckPt == 1 {
		t.pitTime = t.pitTime.Add(60*time.Second - time.Duration(t.pitTime.Second()))
	}
}

func (t *timer) resetPitTimeRaw() {
	pitSchedulestr := strings.Split(time.Now().Local().Format(format), " ")[0] + " "
	if t.minCheckPt == 3 {
		pitSchedulestr += t.pit[2] + ":" + t.pit[1] + ":" + t.pit[0]
	} else if t.minCheckPt == 2 {
		pitSchedulestr += fromStringValue(time.Now().Local().Hour()) +
			":" + t.pit[1] + ":" + t.pit[0]
	} else if t.minCheckPt == 1 {
		pitSchedulestr += fromStringValue(time.Now().Local().Hour()) +
			":" + fromStringValue(time.Now().Local().Minute()) + ":" + t.pit[0]
	}
	t.pitTime, _ = time.ParseInLocation(format, pitSchedulestr, time.Local)

}
