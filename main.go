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

func main() {
	log.Printf("signal-catcher started. pid %d", os.Getpid())
	log.Println("waiting on signal")
	s := <-interrupt()
	log.Printf("caught signal: %s", s)
	log.Println("waiting 3s")
	<-time.After(3 * time.Second)
	log.Println("waiting done. exiting")
}
