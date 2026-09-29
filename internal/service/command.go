package service

import (
	"github.com/Samuel1604/command-clipboard/internal/command"
	"github.com/Samuel1604/command-clipboard/internal/repository"
)

// CommandService contains the application Logic for working
// with saved commands.
//
// It depends on the repository interface rather than a concrete
// storage implementation.
type CommandService struct {
	repository repository.CommandRepository
}

// NewCommandService creates a new CommandService.
//
// The repository is passed into the service so the service does
// not need to know how commands are persisted.
func NewCommandService(repo repository.CommandRepository) *CommandService {
	return &CommandService{
		repository: repo,
	}
}

// Save creates and stores a new command.
func (s *CommandService) Save(text string) (command.Command, error) {
	// The service creates the domain object.
	//
	// The ID is temporarily zero because we haven't implemented
	// ID generation in our storage layer yet.
	cmd := command.Command{
		ID:      0,
		Command: text,
	}

	return s.repository.Save(cmd)
}

// List returns all saved commands.
func (s *CommandService) List() ([]command.Command, error){
	return s.repository.List()	
}
