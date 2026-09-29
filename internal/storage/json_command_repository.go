package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Samuel1604/command-clipboard/internal/command"
	"github.com/Samuel1604/command-clipboard/internal/repository"
)

var _ repository.CommandRepository = (*JSONCommandRepository)(nil)

// JSONCommandRepository stores commands inside a JSON file.
type JSONCommandRepository struct {
	filePath string
}

// NewJSONCommandRepository creates a JOSN-backed command repository.
func NewJSONCommandRepository(filePath string) *JSONCommandRepository {
	return &JSONCommandRepository{
		filePath: filePath,
	}
}

// Save stores a command in the JSON file.
func (r *JSONCommandRepository) Save(cmd command.Command) (command.Command, error) {
	commands, err := r.load()
	if err != nil {
		return command.Command{}, err
	}

	// Generate the next ID.
	cmd.ID = nextID(commands)

	commands = append(commands, cmd)

	if err := r.save(commands); err != nil {
		return command.Command{}, err
	}

	return cmd, nil
}

// load reads all commands from the JSON file.
//
// If the file doesn't exist yet, we return an empty list.
// This allows the first `save` operation to create the file.
func (r *JSONCommandRepository) load() ([]command.Command, error) {
	data, err := os.ReadFile(r.filePath)

	if err != nil {
		if os.IsNotExist(err) {
			return []command.Command{}, nil
		}

		return nil, fmt.Errorf("read commands: %w", err)
	}

	var commands []command.Command

	if err := json.Unmarshal(data, &commands); err != nil {
		return nil, fmt.Errorf("decode commands: %w", err)
	}

	return commands, nil
}

// save writes all commands to the JSON file.
func (r *JSONCommandRepository) save(commands []command.Command) error {
	data, err := json.MarshalIndent(commands, "", " ")
	if err != nil {
		return fmt.Errorf("encode commands: %w", err)
	}

	if err := os.WriteFile(r.filePath, data, 0o600); err != nil {
		return fmt.Errorf("write commands: %w", err)
	}

	return nil
}

// nextID returns the next available command ID.
func nextID(commands []command.Command) int {
	maxID := 0

	for _, cmd := range commands {
		if cmd.ID > maxID {
			maxID = cmd.ID
		}
	}

	return maxID + 1
}
