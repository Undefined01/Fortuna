package main

import (
	"github.com/Undefined01/fortuna/backend/server"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	server.Start()
}
