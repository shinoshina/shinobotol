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
	intervalUnit []int
	fromUnit     []int

	proRule struct {
		units []unit
	}

	unit struct {
		from     int
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
			if runit[0] == "*" {
				pr.units[i].from = -1
			} else {
				pr.units[i].from, _ = strconv.Atoi(runit[0])
			}
			pr.units[i].interval = -1
		} else {
			if runit[0] == "*" || runit[0] == "" {
				pr.units[i].from = -1
			}
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

func interval(pr proRule) intervalUnit {
	iu := make(intervalUnit, 6)
	for i := len(pr.units) - 1; i >= 0; i-- {
		if pr.units[i].interval != -1 {
			iu[i] = pr.units[i].interval
		}else {
			iu[i] = -1
		}
	}
	return iu
}
func from(pr proRule) fromUnit {

	fu := make(fromUnit, 6)
	for i := len(pr.units) - 1; i >= 0; i-- {
		if pr.units[i].from != -1 {
			fu[i] = pr.units[i].from
		}else {
			fu[i] = -1
		}
	}
	return fu
}
func handleRaw(raw string) (fromUnit, intervalUnit) {
	pr := completion(retrive(raw))
	fu, iu := from(pr), interval(pr)
	return fu, iu
}
