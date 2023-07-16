package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func main() {
	filename := "./xx.log"
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}

	tails, err := tail.TailFile(filename, config)
	if err != nil {
		fmt.Println("tail %s failed,err:%v\n", filename, err)
		return
	}

	//开始读取数据
	var (
		msg *tail.Line
		ok  bool
	)
	for {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen,filename:%s\n",
				tails.Filename)
			time.Sleep(time.Second) //读取出错，等待一秒
			continue
		}
		fmt.Println("msg:", msg.Text)
	}
}
