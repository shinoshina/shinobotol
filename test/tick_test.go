package test

import (
	"fmt"
	"shinobot/sbot/tick"
	"testing"
)

func TestPrivateRetrive(t *testing.T) {
	r := tick.TestRetrive("*/10 */20 */1 * * *")
	for _, v := range r {
		fmt.Print(v)
	}
}

func TestCompletion(t *testing.T) {
	r := tick.TestRetrive("*/10 */20 */1 * * *")
	pr := tick.TestCompletion(r)
	fmt.Println(pr)

}

func TestInterval(t *testing.T) {

	r := tick.TestRetrive("20/3 03 19 * * *")
	pr := tick.TestCompletion(r)
	iu := tick.TestInterval(pr)
	fmt.Println(iu)
}

func TestFrom(t *testing.T) {
	r := tick.TestRetrive("20/3 03 19 * * *")
	pr := tick.TestCompletion(r)
	fu := tick.TestFrom(pr)
	fmt.Println(fu)

}

func TestFromSchedule(t *testing.T) {
	tick.TestFromSchedule("20/3 03 19 * * *")
}

func TestInitTimer(t *testing.T) {
	tick.MockInitTimer("20/3 09 19 * * *")
}

func TestTaskLoop(t *testing.T) {
	ct := tick.NewCronTask("wuyuzi", func() {
		fmt.Println("测试")
	})
	ct.AddRule("10/3 23 20 * * *")

	ct.Start()

}
