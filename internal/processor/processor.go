package processor

import (
	"fmt"
	"proj-13-design/internal/manager"

	"github.com/rs/zerolog/log"
)

type Processor struct {
	manager *manager.Manager
}

func NewProcessor(manager *manager.Manager) *Processor {
	return &Processor{manager: manager}
}

func (p *Processor) Save(name, password string) {
	err := p.manager.WritePassword(name, password)

	if err != nil {
		log.Error().Err(err).Msg("Failed to write password")
	}
}

func (p *Processor) Get(name string) {
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
