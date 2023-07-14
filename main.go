package main

import (
	"context"
	"fmt"
	"gin_bubble/dao"
	"gin_bubble/models"
	"gin_bubble/routers"
	"sync"
	"time"
)

func main1() {
	//创建数据库
	//sql: 	create database bubble;
	//连接数据库
	if err := dao.InitMysql(); err != nil {
		panic(err)
	}
	defer dao.DB.Close()
	//绑定模型
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()

	r.Run(":9090")
}

var wg sync.WaitGroup

func worker1(ch <-chan int) {
	defer wg.Done()
BREAK:
	for {
		select {
		case <-ch:
			goto BREAK
		default:
			fmt.Println("HHH")
		}
		time.Sleep(1 * time.Second)
	}
}

func main2() {
	ch := make(chan int)
	wg.Add(1)
	go worker1(ch)
	time.Sleep(1 * time.Second)
	ch <- 1
	wg.Wait()
	fmt.Println("over")
}

func worker(ctx context.Context) {
	defer wg.Done()
	go worker2(ctx)
BREAK:
	for {
		fmt.Println("worker")
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			break BREAK
		default:
		}

	}
}

func worker2(ctx context.Context) {
	defer wg.Done()
BREAK:
	for {
		fmt.Println("worker2")
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			break BREAK
		default:
		}

	}
}

func main3() {
	ctx, cancl := context.WithCancel(context.Background())
	wg.Add(2)
	go worker(ctx)
	time.Sleep(5 * time.Second)
	cancl()
	wg.Wait()
	fmt.Println("over")
}
