package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		PrintHelp()
		log.Fatalln("no command given")
	}

	cmds := NewCommands()
	progState, err := NewState()
	if err != nil {
		log.Fatalln(err)
	}

	err = cmds.run(progState, command{name: os.Args[1], params: os.Args[2:]})
	if err != nil {
		log.Fatalln(err)
	}
}
