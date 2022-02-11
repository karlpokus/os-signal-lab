package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// interrupt returns a chan that recieves interrupt signals
func interrupt() <-chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c,
		os.Interrupt,
		os.Kill,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
		syscall.SIGQUIT,
		//syscall.SIGSTOP,
	)
	return c
}

// repeat listens for signals forever
func repeat() {
	for s := range interrupt() {
		log.Printf("caught signal: %s", s)
	}
}

// onceGraceful listens for one signal and simulates a graceful exit
func onceGraceful() {
	log.Printf("caught signal: %s", <- interrupt())
	log.Println("waiting 3s")
	<-time.After(3 * time.Second)
}

func main() {
	log.Printf("signal-catcher started. pid %d", os.Getpid())
	log.Println("waiting on signal")
	onceGraceful() // repeat
	log.Println("waiting done. exiting")
}
