package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/mchmarny/thingz-server/server"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile)
}

func main() {

	// make sure we can shutdown gracefully
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	doneDone := make(chan bool)
	errCh := make(chan error, 1)

	go func() {
		router := server.NewRouter()
		errCh <- http.ListenAndServe(":8080", router)
	}()

	go func() {

		for {
			select {
			case err := <-errCh:
				log.Printf("Error: %v", err)
			case ex := <-sigCh:
				log.Println("Shutting down due to: ", ex)
				doneDone <- true
			default:
				// nothing to do
			}
		}

	}()
	<-doneDone

}
