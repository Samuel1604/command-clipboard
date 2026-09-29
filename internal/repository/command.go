package repository

import "github.com/Samuel1604/command-clipboard/internal/command"

// CommandRepository defines how the application interacts with
// stored commands.
//
// The interface describes what the application needs from a
// storage implementation without specifying how that storage
// actually works
type CommandRepository interface {
	// Save stores a command and returns the stored command.
	Save(cmd command.Command) (command.Command, error)
	List() ([]command.Command, error)
}
