package test

import (
	"bytes"
	log "fist_test/helpers"
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func GetGid() (gid uint64) {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}

// 生产消费模型
func RunGoroutine1() {

	c := make(chan int, 10)
	chanCancel := make(chan int)
	for i := 0; i < 5; i++ {
		go producer(c, chanCancel)
	}

	for i := 0; i < 10; i++ {
		go comsumer(c, chanCancel)
	}

	time.Sleep(5 * time.Second)

	fmt.Println("退出")

	close(chanCancel) //发送关闭信号
	//close(c)          //关闭通道
	// runtime.Goexit()
}

// 生产者
func producer(c chan int, chanCancel <-chan int) {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("生产者 err: %v\n", err)
		}

	}()
	var gid = GetGid()
	for {

		for i := 0; i < 10; i++ {
			v := i

			select {
			case _, ok := <-chanCancel: //结束信号
				if !ok {
					//关闭信号
					fmt.Printf("生产者退出，id:%v\n", gid)
					return
				}
			case c <- v:
				msg := fmt.Sprintf("生产v: %v,goid:%v\n", v, gid)
				log.Logger.Info().Msg(msg)

			default:
				log.Logger.Info().Msg("producer default")
			}

			time.Sleep(100 * time.Millisecond)

		}

	}
}

// 消费者
func comsumer(c chan int, chanCancel <-chan int) {
	var gid = GetGid()
	for {
		select {
		case _, ok := <-chanCancel:
			if !ok {
				fmt.Printf("消费者退出:%v\n", gid)
				return
			}
		case v, ok := <-c:
			if !ok {
				fmt.Printf("goroutine %v exit\n", gid)
				break
			}
			msg := fmt.Sprintf("消费v: %v,goid:%v\n", v, gid)
			log.Logger.Info().Msg(msg)
			// time.Sleep(1000 * time.Millisecond)
		}

	}
}
