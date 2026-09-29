package service_test

import (
	"github.com/Samuel1604/command-clipboard/internal/command"
	"github.com/Samuel1604/command-clipboard/internal/service"
	"testing"
)

// fakeCommandRepository is a small in-memory implementation
// of CommandRepository used only for testing.
//
// It doesn't write to disk.
type fakeCommandRepository struct {
	saved command.Command
	commands []command.Command

}

// Save statistics the CommandRepository interface.
//
// Because fakeCommanReposioty has a Save method with the
// correct signature, Go automatically considers it a
// CommandRepository. We don't need to explicitly declare that.
func (f *fakeCommandRepository) Save(cmd command.Command) (command.Command, error) {
	f.saved = cmd

	return cmd, nil
}

func (f *fakeCommandRepository) List() ([]command.Command, error) {
	return f.commands, nil
}

func TestCommandServiceSave(t *testing.T) {
	// Create our fake repository.
	repo := &fakeCommandRepository{}

	repo.saved = command.Command{
		ID: 1,
		Command: "docker compose up -d",
	}

	repo.commands = []command.Command{
		{
			ID:		1,
			Command: "docker compose up -d",
		},
	}
	// Give the service the repository.
	svc := service.NewCommandService(repo)

	// Ask the service to save a command.
	saved, err := svc.Save("docker compose up -d")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify that the command returned by the service is correct.
	if saved.Command != "docker compose up -d" {
		t.Fatalf("expected command to be saved, got %q", saved.Command)
	}

	// Verify that the repository actually received the command.
	if repo.saved.Command != "docker compose up -d" {
		t.Fatalf("expected repository to receive command, got %q", repo.saved.Command)
	}

	// Ask the service to list commands.
	commands, err := svc.List()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(commands) != 1 {
		t.Fatalf("expected 1 command, got %d", len(commands))
	}

	if commands[0].Command != "docker compose up -d" {
		t.Fatalf("expected saved command, got %q", commands[0].Command)
	}
}
