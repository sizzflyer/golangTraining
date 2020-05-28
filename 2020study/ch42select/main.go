package main

import (
	"fmt"
	"time"
)

func push(c chan int) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		c <- i
		i++
	}
}

func main() {
	c := make(chan int)
	go push(c)
	timerChan := time.After(10 * time.Second)
	tickTimerChan := time.Tick(2 * time.Second)
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		// default:	//1초동안 +1씩 찍고 입력 값이 없을 때는 default
		// 	fmt.Println("Idle")
		// 	time.Sleep(1 * time.Second)
		// }
		case <-timerChan:
			fmt.Println("timeout")
			return
		case <-tickTimerChan:
			fmt.Println("Tick")
		}
	}
}
