package main

import (
	"context"
	"fmt"
	"time"
)

func main4() {
	d := time.Now().Add(100 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()

	for {
		fmt.Println("go")
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		}
	}

}
