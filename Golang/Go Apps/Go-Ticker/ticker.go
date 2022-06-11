package main

import (
	"log"
	"os"
	"os/signal"
	"time"
)

//ticker and timeout duration
const (
	duration = 1 * time.Second
	timeout  = 3 * time.Second
)

func main() {
	log.Println("application started")

	//create channel for interupt signal
	is := make(chan os.Signal, 1)
	signal.Notify(is, os.Interrupt)

	//create channel for stop signal
	ss := make(chan struct{})

	//create ticker
	t := time.NewTicker(duration)

	//create timeout
	to := time.After(timeout)

	//start listening timer ticks and stop signal
	go func() {
		defer close(ss)
		for {
			select {
			case <-t.C:
				log.Println("ticker signal received")
			case <-ss:
				log.Println("stop signal received")
				return
			}
		}
	}()

	//wait for timeout or interupt signal
	select {
	case <-is:
		log.Println("interupt signal received")
	case <-to:
		log.Println("application timed out")
	}

	//stop ticker
	log.Println("ticker stopped")
	t.Stop()

	//send stop signal
	log.Println("stop signal sent")
	ss <- struct{}{}

	//wait until stop signal channel is closed (goroutine is finished)
	<-ss
	log.Println("stop signal channel closed")

	//close interupt signal channel
	close(is)
	log.Println("interupt signal channel closed")

	log.Println("application stopped")
}
