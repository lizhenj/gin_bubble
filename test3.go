package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

type TraceCode string

var wg3 sync.WaitGroup

func worker5(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(TraceCode)
	if !ok {
		fmt.Println("invalid trace code")
	}
	log.Printf("%s worker func...", traceCode)
LOOP:
	for {
		fmt.Printf("worker,trace code:%s\n", traceCode)
		time.Sleep(time.Microsecond * 1)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	fmt.Println("worker done")
	wg3.Done()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*1000)
	//在系统的入口中设置trace code 传递给后续启动的goroutine实现日志数据聚合
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), TraceCode("9859777115"))
	log.Printf("%s main函数", "9859777115")
	wg3.Add(1)
	go worker5(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg3.Wait()
	fmt.Println("over...")
}
