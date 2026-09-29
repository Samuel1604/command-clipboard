package main

import (
	"fmt"
	"os"
	"path/filepath"
	
	"github.com/Samuel1604/command-clipboard/internal/cli"
	"github.com/Samuel1604/command-clipboard/internal/service"
	"github.com/Samuel1604/command-clipboard/internal/storage"	
)

func main() {
	// main is responsible only for stating the application
	//
	// os.UserConfigDir returns the operating-system-specific
	// directory intended for user application configuration.
	//
	// On Linux this will typically resolve to something like:
	//
	// 	~/.config
	//
	// We keep Command Clipboard's data inside it's own directory
	// so it doesn't pollute the user's current working directory.
	configDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Printf("Error determining config directory: %v\n", err)
		return
	}
	
	appDir := filepath.Join(configDir, "cmdclip")

	// Create the application directory if it doesn't exist.
	if err := os.MkdirAll(appDir, 0o700); err != nil {
		fmt.Printf("Error creating application directory: %v\n", err)
		return
	}

	dataFile := filepath.Join(appDir, "commands.json")

	// Create the JSON repository.
	repository := storage.NewJSONCommandRepository(dataFile)

	// Give the repository to the service.
	commandService := service.NewCommandService(repository)

	// Start the CLI with the application service.
	cli.Run(cli.Args(), commandService)
}
