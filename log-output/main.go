package main

import (
	"crypto/rand"
	"log"
	"os"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	value := rand.Text()
	for {
		logger.Println(value)
		time.Sleep(5 * time.Second)
	}
}
