package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func worker4(ctx context.Context) {
LOOP:
	for {
		fmt.Println("db connecting ...")
		time.Sleep(time.Microsecond * 100) //假设正常连接数据库的耗时
		select {
		case <-ctx.Done(): //设定时间到后deadline会发信号
			break LOOP
		default:
			fmt.Println("wait")
		}
	}
	fmt.Println("worker done!")
	wg2.Done()
}

func main5() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*50)
	wg2.Add(1)
	go worker4(ctx)
	time.Sleep(time.Second * 3)
	cancel()
	wg2.Wait()
	fmt.Println("over")
}
