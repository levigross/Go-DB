package dhash

import (
	"../timenet"
	"time"
)

type timerServer timenet.Timer

func (self *timerServer) ActualTime(x int, result *time.Time) error {
	*result = (*timenet.Timer)(self).ActualTime()
	return nil
}
