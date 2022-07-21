package tick

import (
	"fmt"
	"strconv"
	"strings"
)

//   *   *   *   *   *   *   //
//  sec min hor day mon wek  //
//  /1  /30   /5   *   *   *   //
type (
	raw          []string
	intervalUnits []int
	fromUnits     []string

	proRule struct {
		units []unit
	}

	unit struct {
		from     string
		interval int
	}
)

func retrive(rule string) (r raw) {
	r = make(raw, 6)
	r = strings.Split(rule, " ")
	return
}

func completion(r raw) (pr proRule) {
	pr.units = make([]unit, 6)
	for i, v := range r {
		runit := strings.Split(v, "/")
		if len(runit) == 1 {
			//judge from num legal
			pr.units[i].from = runit[0]
			pr.units[i].interval = -1
		} else {
			pr.units[i].from = runit[0]
			inum, err := strconv.Atoi(runit[1])
			if err != nil {
				fmt.Println(err)
				pr.units[i].interval = -1
			} else {
				pr.units[i].interval = inum
			}
		}
	}
	return
}

func interval(pr proRule) intervalUnits {
	iu := make(intervalUnits, 6)
	for i := len(pr.units) - 1; i >= 0; i-- {
		if pr.units[i].interval != -1 {
			iu[i] = pr.units[i].interval
		} else {
			iu[i] = 0
		}
	}
	return iu
}
func from(pr proRule) fromUnits {

	fu := make(fromUnits, 6)
	for i := len(pr.units) - 1; i >= 0; i-- {
		fu[i] = pr.units[i].from
	}
	return fu
}
func handleRaw(raw string) (fromUnits, intervalUnits) {
	pr := completion(retrive(raw))
	fu, iu := from(pr), interval(pr)
	return fu, iu
}
