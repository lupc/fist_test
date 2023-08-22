package test

import (
	"fmt"
	"time"
)

func RunTimer() {
	t1 := time.NewTicker(time.Second * 5)
	t2 := time.NewTicker(time.Second * 20)
	go executeJob(t1, t2)
	select {}
}

func executeJob(t1 interface{}, t2 interface{}) {
	println("executeJob called")
	for {
		select {
		case tick1 := <-t1.(*time.Ticker).C:
			fmt.Printf("t1 tick1: %vtime:%v\n", tick1, time.Now())
		case tick2 := <-t2.(*time.Ticker).C:
			fmt.Printf("t2 tick2: %v,time:%v\n", tick2, time.Now())
		}

	}
}
