package main

import (
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

func worker(ch <-chan int) {
	defer wg.Done()
	for {
		select {
		case <-ch:
			goto BREAK
		default:
			fmt.Println("HHH")
		}
		time.Sleep(1 * time.Second)
	}
BREAK:
}

func main() {
	ch := make(chan int)
	wg.Add(1)
	go worker(ch)
	time.Sleep(1 * time.Second)
	ch <- 1
	wg.Wait()
	fmt.Println("over")
}
