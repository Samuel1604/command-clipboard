package command_test

import (
	"github.com/Samuel1604/command-clipboard/internal/command"
	"testing"
)

func TestCommand(t *testing.T) {
	// Create a Command using the domain model.
	savedCommand := command.Command{
		ID:      1,
		Command: "docker compose up -d",
	}

	// Verify that the ID was stored correctly.
	if savedCommand.ID != 1 {
		t.Fatalf("expected ID 1, got %d", savedCommand.ID)
	}

	// Verify that the command text was stored correctly.
	if savedCommand.Command != "docker compose up -d" {
		t.Fatalf("expected command text to match, got %q", savedCommand.Command)
	}

}
