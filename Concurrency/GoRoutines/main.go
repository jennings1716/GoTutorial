package main

import (
	"fmt"
	"time"
)

func printNow(done chan bool) {
	fmt.Println("PRINTING")
	done <- true
}

func pauseNow(done chan bool) {
	fmt.Println("PAUSE...")
	time.Sleep(3 * time.Second)
	done <- true
}

func main() {
	done := make(chan bool)

	go printNow(done)
	<-done

	go printNow(done)
	<-done

	go pauseNow(done)
	<-done

	go printNow(done)
	<-done
}
