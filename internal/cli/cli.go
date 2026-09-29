package cli

import (
	"fmt"
	"os"

	"github.com/Samuel1604/command-clipboard/internal/service"
)

// Run starts the Command Clipboard command-line interface.
//
// The service is passed into the CLI so the CLI doesn't need to
// know how commands are stored.
func Run(args []string, commandService *service.CommandService) {
	// We needd at least one argument after the program name.
	if len(args) < 1 {
		fmt.Println("Usage: cmd <command> [arguments...]")
		return
	}

	// The first argument is the subcommand the user wants to run.
	subcommand := args[0]

	switch subcommand {
	case "save":
		// The save command needs another argument:
		// the command we want to store.
		if len(args) < 2 {
			fmt.Println("Usage: cmd save <command>")
			return
		}

		// Ask the application service to save the command
		//
		// The CLI doesn't know whether the command is stored in
		// JSON, SQLite, PostgreSQL, or somewhere else.
		savedCommand, err  := commandService.Save(args[1])

		if err != nil {
			fmt.Printf("Error saving command: %v\n", err)
			return
		}

		fmt.Printf("Saved command #%d\n", savedCommand.ID)
	
	case "list":
		commands, err := commandService.List()
		if err != nil {
			fmt.Printf("Error listing commands: %v\n", err)
			return
		}

		if len(commands) == 0 {
			fmt.Println("No commands saved.")
			return
		}

		for _, cmd := range commands {
			fmt.Printf("#%d %s\n", cmd.ID, cmd.Command)
		}

	default:
		fmt.Printf("Unknown command: %s\n", subcommand)
	}
}

// Args return the command-line arguments without the program path.
//
// os.Args contains the program path at index 0.
// We don't want the CLI package to care about that detail,
// so we remove it here.
func Args() []string {
	if len(os.Args) <= 1 {
		return []string{}
	}

	return os.Args[1:]
}
