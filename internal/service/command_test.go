package service_test

import (
	"testing"

	"github.com/Samuel1604/command-clipboard/internal/command"
	"github.com/Samuel1604/command-clipboard/internal/service"
)

// fakeCommandRepository is a small in-memory implementation
// of CommandRepository used only for testing.
//
// It doesn't write to disk.
type fakeCommandRepository struct {
	saved    command.Command
	commands []command.Command
}

// Save satisfies the CommandRepository interface.
//
// Because fakeCommandRepository has a Save method with the
// correct signature, Go automatically considers it a
// CommandRepository. We don't need to explicitly declare that.
func (f *fakeCommandRepository) Save(cmd command.Command) (command.Command, error) {
	f.saved = cmd

	return cmd, nil
}

// List returns the commands stored in our fake repository.
func (f *fakeCommandRepository) List() ([]command.Command, error) {
	return f.commands, nil
}

func TestCommandServiceSave(t *testing.T) {
	// Create our fake repository.
	repo := &fakeCommandRepository{}

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
}

func TestCommandServiceList(t *testing.T) {
	// Start with one command already in the fake repository.
	repo := &fakeCommandRepository{
		commands: []command.Command{
			{
				ID:      1,
				Command: "docker compose up -d",
			},
		},
	}

	svc := service.NewCommandService(repo)

	// Ask the service for all commands.
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

func TestCommandServiceFind(t *testing.T) {
	// Give the fake repository multiple commands so we can test
	// whether Find correctly filters them.
	repo := &fakeCommandRepository{
		commands: []command.Command{
			{
				ID:      1,
				Command: "docker compose up -d",
			},
			{
				ID:      2,
				Command: "git status",
			},
			{
				ID:      3,
				Command: "docker ps -a",
			},
		},
	}

	svc := service.NewCommandService(repo)

	// Search for commands containing "docker".
	matches, err := svc.Find("docker")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(matches) != 2 {
		t.Fatalf("expected 2 matches, got %d", len(matches))
	}

	if matches[0].ID != 1 {
		t.Fatalf("expected first match to have ID 1, got %d", matches[0].ID)
	}

	if matches[1].ID != 3 {
		t.Fatalf("expected second match to have ID 3, got %d", matches[1].ID)
	}
}