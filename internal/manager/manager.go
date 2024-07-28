package manager

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

type Passwords map[string]string

type PasswordManager interface {
	WritePassword(name, password string) error
	GetPasswords() (Passwords, error)
	GetPassword(name string) string
}

type Manager struct {
	filename string
}

func NewManager(filename string) *Manager {
	return &Manager{filename: filename}
}

func (m *Manager) WritePassword(name, password string) error {
	file, err := os.OpenFile(m.filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModeAppend)

	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}

	defer file.Close()

	_, err = file.WriteString(name + "=" + password + "\n")

	if err != nil {
		return fmt.Errorf("failed to write password: %v", err)
	}

	return nil
}

func (m *Manager) GetPasswords() (Passwords, error) {
	fileContent, err := os.ReadFile(m.filename)

	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	passwords := make(map[string]string)

	lines := strings.Split(string(fileContent), "\n")

	for _, l := range lines {
		l = strings.TrimSpace(l)

		if l == "" {
			continue
		}

		parts := strings.Split(l, "=")

		name := parts[0]
		password := parts[1]

		passwords[name] = password
	}

	return passwords, nil
}

func (m *Manager) GetPassword(name string) string {
	passwords, err := m.GetPasswords()

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get passwords")
	}

	return passwords[name]
}
