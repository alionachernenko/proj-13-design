package main

import (
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

	manager := manager.NewManager("password.txt")
	processor := processor.NewProcessor(args, manager)

	switch command := args[1]; command {
	case "save":
		processor.Save()
	case "list":
		processor.List()
	case "get":
		processor.Get()
	}
}
