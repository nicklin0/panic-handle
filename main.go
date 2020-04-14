package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("panic-handle")

	go somethingPanic()
	for true {
		time.Sleep(time.Duration(1) * time.Millisecond)
	}
}

func somethingPanic() {
	fmt.Println("something panic later")
	time.Sleep(time.Duration(1) * time.Millisecond)

	// try recover in defer and launch a new goroutine to restart
	defer func() {
		fmt.Println("in defer")
		err := recover()
		if err != nil {
			fmt.Printf("in recover from panic = %s\n", err)
			go somethingPanic()
		}
	}()

	// do something possible panic here
	panic("it is panic!")
}
