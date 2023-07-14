package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

//服务端

func indexHandler(w http.ResponseWriter, r *http.Request) {
	number := rand.Intn(2)
	fmt.Println(number)
	if number == 0 {
		time.Sleep(10 * time.Second)
		fmt.Fprintf(w, "slow response")
		return
	}
	fmt.Fprint(w, "quick response")
}

func main() {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}
