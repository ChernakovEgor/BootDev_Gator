package main

import (
	// "fmt"
	"database/sql"
	"log"
	"os"

	"github.com/ChernakovEgor/gator/internal/config"
	"github.com/ChernakovEgor/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.Read()
	if err != nil {
		log.Fatalf("could not read config: %v", err)
	}

	db, err := sql.Open("postgres", config.URL)
	if err != nil {
		log.Fatalf("could not open connection to db: %v", err)
	}

	dbQueries := database.New(db)

	programState := state{&config, dbQueries}
	m := make(map[string]func(*state, command) error)

	commands := commands{m}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAgg)
	commands.register("addfeed", handlerAddFeed)
	commands.register("feeds", handlerFeeds)

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
