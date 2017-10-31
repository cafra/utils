package utils

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

var (
	tps             int64 = 0
	goroutine_count int   = 1000
	ops_count       int   = 1000
)

type TestFun func() error

func CurrentTest(do TestFun, args ...int) {
	fmt.Println("======== START ======")

	exit := make(chan bool)

	if len(args) > 0 && args[0] > 0 {
		goroutine_count = args[0]
	}
	if len(args) > 1 && args[1] > 0 {
		ops_count = args[1]
	}

	for cs := 0; cs < goroutine_count; cs++ {
		go func() {
			for ps := 0; ps < ops_count; ps++ {
				if err := do(); err != nil {
					fmt.Errorf("err=%v", err)
					os.Exit(0)
				}
				atomic.AddInt64(&tps, 1)
			}
		}()
	}
	go listen(exit)

	<-exit
	fmt.Println("======== END ======")
}

func listen(exit chan bool) {
	ticket := time.NewTicker(time.Second)
	for range ticket.C {
		tp := atomic.SwapInt64(&tps, 0)
		fmt.Println("tps=", tp)
		//if tp == 0 {
		//	exit <- true
		//	return
		//}
	}
}
