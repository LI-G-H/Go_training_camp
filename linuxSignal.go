package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	signals := make(chan os.Signal, 1)
	// 监听终端信号量
	signal.Notify(signals)

	// 接受 ctrl + C 信号
	s := <-signals
	fmt.Println("Get Signal : ", s)
}
