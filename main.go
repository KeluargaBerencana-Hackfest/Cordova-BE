package main

import (
	"log"
	"os"
	"sync"

	"github.com/Ndraaa15/cordova/app"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Printf("[cordova-main] failed to load .env file. Error : %v\n", err)
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		os.Exit(app.RunServer())
	}()

	wg.Wait()
}
