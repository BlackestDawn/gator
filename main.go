package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/BlackestDawn/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	if len(os.Args) <= 1 {
		PrintHelp()
		log.Fatalln("no command given")
	}

	cmds := NewCommands()
	progState, err := NewState()
	if err != nil {
		log.Fatalf("error creating program state: %v", err)
	}

	db, err := sql.Open("postgres", progState.conf.DbURL)
	if err != nil {
		log.Fatalf("error connecting to DB: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)
	progState.db = dbQueries

	err = cmds.run(progState, command{name: os.Args[1], params: os.Args[2:]})
	if err != nil {
		log.Fatalln(err)
	}
}
