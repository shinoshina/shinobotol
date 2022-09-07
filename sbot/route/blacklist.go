package route

import (
	"fmt"
	"strconv"
	"sync"
)

type blacklist struct{
	blm map[float64]bool
	blf bool
	mutex sync.Mutex
}
func NewBlackList()(bl* blacklist){
	bl = new(blacklist)
	bl.blm = make(map[float64]bool)
	bl.blf = false
	return
}
func (bl *blacklist) add(sid string) {
	fid, err := strconv.ParseFloat(sid, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	bl.mutex.Lock()
	bl.blm[fid] = true
	bl.mutex.Unlock()
}
func (bl *blacklist) addList(sids []string) {
	for _, sid := range sids {
		bl.add(sid)
	}
}
func (bl *blacklist) remove(sid string) {
	fid, err := strconv.ParseFloat(sid, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	bl.mutex.Lock()
	bl.blm[fid] = false
	bl.mutex.Lock()
}
