package main

import (
	// "fmt"
	"log"
	"os"

	"github.com/ChernakovEgor/gator/internal/config"
)

func main() {
	config, err := config.Read()
	if err != nil {
		log.Fatalf("could not read config: %v", err)
	}

	programState := state{&config}
	m := make(map[string]func(*state, command) error)
	commands := commands{m}

	commands.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		log.Fatalf("too few arguments")
		os.Exit(1)
	}

	command := command{args[1], args[2:]}

	err = commands.run(&programState, command)
	if err != nil {
		log.Fatalf("could not run command: %v", err)
	}
}
