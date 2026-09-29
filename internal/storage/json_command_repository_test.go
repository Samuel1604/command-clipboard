package storage_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Samuel1604/command-clipboard/internal/command"
	"github.com/Samuel1604/command-clipboard/internal/storage"
)

func TestJSONCommandRepositorySave(t *testing.T) {
	// t.TempDir creates a temporary directory for this test.
	//
	// Go automatically removes it when the test finishes.
	tempDir := t.TempDir()

	filePath := filepath.Join(tempDir, "commands.json")

	repo := storage.NewJSONCommandRepository(filePath)

	// Save the first command.
	first, err := repo.Save(command.Command{
		Command: "docker compose up -d",
	})

	if err != nil {
		t.Fatalf("unexpected error saving first command: %v", err)
	}

	if first.ID != 1 {
		t.Fatalf("expected first command ID to be 1, got %d", first.ID)
	}

	// Save a second command.
	second, err := repo.Save(command.Command{
		Command: "go test ./...",
	})

	if err != nil {
		t.Fatalf("unexpected error saving second command: %v", err)
	}

	if second.ID != 2 {
		t.Fatalf("expected second command ID to be 2, got %d", second.ID)
	}

	// Verify that the JSON file was actually created.
	if _, err := os.Stat(filePath); err != nil {
		t.Fatalf("expected JSON file to exist: %v", err)
	}
}
