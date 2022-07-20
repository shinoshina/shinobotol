package tick

import "time"

type timer struct {
	schedules       []schedule
	currentSchedule int
}
type pitUnit struct {
	sec  int
	min  int
	hour int
	day  int
}
type schedule struct {
	pit pitUnit 
	duration pitUnit
}

func CurrentTime() (ct string) {

	s := time.Time{}
	s.Sub(time.Now())
	tunix := time.Now().Unix()
	ct = time.Unix(tunix, 0).Format("2006-01-02 15:04:05")

	return
}
func newTimer()(t *timer){
	t = new(timer)
	t.schedules = make([]schedule, 1)
	t.currentSchedule = -1
	return
}
func(t *timer)fromSchedule(raw string){

	// sc := schedule{}
	// fu,iu := handleRaw(raw)
    

}
