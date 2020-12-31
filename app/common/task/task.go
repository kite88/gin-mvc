package task

import (
	"fmt"
	"time"
)

func Start() {
	go missionA() //协程开启
	go missionB() //协程开启
}

func missionA() {
	fmt.Println("定时器A任务启动")
	for ; ; {
		timer := time.NewTimer(5 * time.Second)
		t := <-timer.C
		fmt.Printf("定时器A任务正在执行:%v\n", t.Format("2006-01-02 15:04:05"))

	}
}

func missionB() {
	fmt.Println("定时器B任务启动")
	for ; ; {
		timer := time.NewTimer(10 * time.Second)
		t := <-timer.C
		fmt.Printf("定时器B任务正在执行:%v\n", t.Format("2006-01-02 15:04:05"))

	}
}
