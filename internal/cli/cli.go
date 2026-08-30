package cli

import (
	"fmt"
	"github.com/Samuel1604/command-clipboard/internal/command"
	"os"
)

// Run starts the Command Clipboard command-line interface.
//
// It receives the arguments passed to the program and decides
// what operation the user want to perform.
func Run(args []string) {
	// We need at least one argument after the program name.
	// For example:
	//
	//	cmd save "docker compose up -d"
	//
	// The arguments we receive here are:
	//
	//	["save", "docker compose up -d"]
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

		// Create a domain Command from the user's input.
		// We don't have persistence yet, so we temporarily give
		// it an ID of 0. The storage layer will eventually be
		// responsible for assigning real IDs.
		savedCommand := command.Command{
			ID:      0,
			Command: args[1],
		}

		fmt.Printf("Received command: %s\n", savedCommand.Command)
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
