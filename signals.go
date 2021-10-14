package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func processSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	for {
		<-c
		log.Println("Quitting...")
		os.Exit(0)
	}
}
