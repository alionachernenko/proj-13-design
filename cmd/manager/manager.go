package main

import (
	"fmt"
	"os"
	"proj-13-design/internal/manager"
	"proj-13-design/internal/processor"

	"github.com/rs/zerolog/log"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatal().Msg("Expected minimum 1 argument: <command>")
		return
	}

	manager := manager.NewManager("passwords.txt")
	processor := processor.NewProcessor(manager)

	switch command := args[1]; command {
	case "save":
		name := args[2]
		password := args[3]

		processor.Save(name, password)
	case "list":
		processor.List()
	case "get":
		name := args[2]
		processor.Get(name)
	default:
		fmt.Println("Invalid command. Available commands:\nsave <name> <password>\tsave password with certain name\nlist\tshow all saved passwords\nget <name>\tget password by name")
	}
}
