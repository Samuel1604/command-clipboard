package command

// Command represents a command saved by the user.
//
// This is a domain model. It describes what a saved command
// looks like without caring about where the command is stored
// or how the user interacts with the application.
type Command struct {
	ID      int    `json:"id"`
	Command string `json:"command"`
}
