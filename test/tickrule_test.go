package test

import (
	"fmt"
	"shinobot/sbot/tick"
	"testing"
)

func TestPrivateRetrive(t *testing.T) {
	r := tick.TestRetrive("/10 */20 */1 */SHABI */HI *")
	for _, v := range r {
		fmt.Print(v)
	}
}

func TestCompletion(t *testing.T) {
	r := tick.TestRetrive("10 20 */1 */SHABI */HI *")
	pr := tick.TestCompletion(r)
	fmt.Println(pr)

}

func TestInterval(t *testing.T) {

	r := tick.TestRetrive("10 20 */1 */SHABI */HI *")
	pr := tick.TestCompletion(r)
	iu := tick.TestInterval(pr)
	fmt.Println(iu)
}

func TestFrom(t *testing.T) {
	r := tick.TestRetrive("10 20 */1 */SHABI */HI *")
	pr := tick.TestCompletion(r)
	fu := tick.TestFrom(pr)
	fmt.Println(fu)

}
