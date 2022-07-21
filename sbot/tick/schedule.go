package tick

import (
	"errors"
	"time"
)

const(
	
	Second time.Duration = time.Second
	
	Minute time.Duration = time.Minute

	Hour time.Duration = time.Hour

	Day time.Duration = time.Hour * 24
	
	Week time.Duration = Day * 7
)

type Schedule interface {
	Next(t time.Time) time.Time
}


type AtSchedule interface {
	At(t string) Schedule
	Schedule
}

func Every(p time.Duration) AtSchedule {

	if p < time.Second {
		p = Second
	}

	p = p - time.Duration(p.Nanoseconds())%time.Second 

	return &periodicSchedule{
		period: p,
	}
}

type periodicSchedule struct {
	period time.Duration
}


func (ps periodicSchedule) Next(t time.Time) time.Time {
	return t.Truncate(time.Second).Add(ps.period)
}


func (ps periodicSchedule) At(t string) Schedule {
	if ps.period < Day {
		panic("period must be at least in days")
	}


	h, m, err := parse(t)

	if err != nil {
		panic(err.Error())
	}

	return &atSchedule{
		period: ps.period,
		hh:     h,
		mm:     m,
	}
}

func parse(hhmm string) (hh int, mm int, err error) {

	hh = int(hhmm[0]-'0')*10 + int(hhmm[1]-'0')
	mm = int(hhmm[3]-'0')*10 + int(hhmm[4]-'0')

	if hh < 0 || hh > 24 {
		hh, mm = 0, 0
		err = errors.New("invalid hh format")
	}
	if mm < 0 || mm > 59 {
		hh, mm = 0, 0
		err = errors.New("invalid mm format")
	}

	return
}

type atSchedule struct {
	period time.Duration
	hh     int
	mm     int
}


func (as atSchedule) reset(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), as.hh, as.mm, 0, 0, time.Local)
}


func (as atSchedule) Next(t time.Time) time.Time {
	next := as.reset(t)
	if t.After(next) {
		return next.Add(as.period)
	}
	return next
}
