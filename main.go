package main

import (
	"self-payrol/config"
	"sync"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

// main is the entry point of the application
func main() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		// log error message if .env file is not loaded properly
		log.Infof(".env is not loaded properly")
	}

	// log message after .env file is read
	log.Infof("read .env from file")

	// initialize config
	config := config.NewConfig()

	// initialize server
	server := InitServer(config)

	// create a wait group to synchronize goroutines
	wg := sync.WaitGroup{}

	// add a goroutine to wait group
	wg.Add(1)

	// start a new goroutine
	go func() {
		// signal wait group when goroutine exits
		defer wg.Done()
		// start server
		server.Run()
	}()

	// wait for all goroutines to finish
	wg.Wait()
}
