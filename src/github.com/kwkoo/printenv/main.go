package main

import (
	"log"
	"os"
	"time"
)

const sleepInterval = 60

func main() {
	for {
		for _, e := range os.Environ() {
			log.Print(e)
		}
		log.Printf("sleeping for %d seconds...", sleepInterval)
		time.Sleep(time.Duration(sleepInterval) * time.Second)
	}
}
