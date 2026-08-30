package main

import "github.com/Samuel1604/command-clipboard/internal/cli"

func main() {
	// main is responsible only for stating the application
	// the CLI package handles the actual command-line behavior.
	cli.Run(cli.Args())
}
