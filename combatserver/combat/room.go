package combat

import (
	"fmt"
	"sync"
	"time"
)

const (
	JoinRoom = 1
	Rool     = 2
)

type PlayerInfoEvent struct {
	Type   int32
	RoomID string
	Uid    string
}

var roolEventChan chan *PlayerInfoEvent = make(chan *PlayerInfoEvent, 2048)

func InitRoom() {
	var once sync.Once
	onceInit := func() {
		tickerSetServer := time.NewTicker(time.Second)
		done := make(chan bool)
		go func() {
			for {
				select {
				case <-done:
					return
				case <-roolEventChan:
					fmt.Println()
				case <-tickerSetServer.C:
				}
			}
		}()
	}
	once.Do(onceInit)
}
