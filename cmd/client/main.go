package main

import (
	"log"

	"github.com/lgastineau/goodreads/cmd/client/commands"
)

func main() {
	log.Println("main hello world")
	commands.Execute()
}
