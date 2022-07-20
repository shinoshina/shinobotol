package tick


var (
	TestRetrive = retrive
	TestCompletion = completion
	TestFrom = from
	TestInterval = interval
)

var (
	TestFromSchedule = newTimer().fromSchedule
)

func MockInitTimer(raw string){
	t := newTimer()
	t.fromSchedule(raw)
	t.wait()
}

