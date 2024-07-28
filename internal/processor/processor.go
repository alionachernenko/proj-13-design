package processor

import (
	"fmt"
	"proj-13-design/internal/manager"

	"github.com/rs/zerolog/log"
)

type Processor struct {
	args    []string
	manager *manager.Manager
}

func NewProcessor(args []string, manager *manager.Manager) *Processor {
	return &Processor{args: args, manager: manager}
}

func (p *Processor) Save() {
	name := p.args[2]
	password := p.args[3]

	err := p.manager.WritePassword(name, password)

	if err != nil {
		log.Error().Err(err).Msg("Failed to write password")
	}
}

func (p *Processor) Get() {
	name := p.args[2]
	password := p.manager.GetPassword(name)

	fmt.Println(password)
}

func (p *Processor) List() {
	passwords, err := p.manager.GetPasswords()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get passwords")
		return
	}

	for k, v := range passwords {
		fmt.Printf("%v=%v\n", k, v)
	}
}
