package main

import (
	"log"
)

func main() {
	log.Println("starting nulpunt server application")

	// setup connection to mongodb, check indexes, etc.
	initPersistency()
	log.Println("persistency set up")

	// start a http server
	initHTTPServer()
	log.Println("http server set up")

	// all seems good
	log.Println("running..")
	// wait forever
	select {}
}
