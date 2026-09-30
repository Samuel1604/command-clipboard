# Command Clipboard

A lightweight Go CLI for saving, searching, and managing frequently used terminal commands.

## Why?

There are commands I use repeatedly but don't always remember.

Instead of searching through shell history or the internet every time, Command Clipboard gives me a simple place to save commands and retrieve them when I need them.

## Features

* Save frequently used commands
* List saved commands
* Search commands by text
* Retrieve a command by ID
* Delete saved commands
* Persistent local storage
* No database or external service required

## Usage

### Save a command

```bash
cmdclip save "docker compose up -d"
```

Output:

```text
Saved command #1
```

### List commands

```bash
cmdclip list
```

Output:

```text
#1  docker compose up -d
#2  git log --oneline --graph
#3  go test ./...
```

### Find a command

```bash
cmdclip find docker
```

Output:

```text
#1  docker compose up -d
```

### Get a command

```bash
cmdclip get 1
```

Output:

```text
#1  docker compose up -d
```

### Delete a command

```bash
cmdclip delete 1
```

Output:

```text
Deleted command #1
```

## Storage

Command Clipboard stores commands locally as JSON.

The storage location is determined using Go's user configuration directory support:

```text
<user-config-directory>/cmdclip/commands.json
```

On Linux, this will typically resolve to:

```text
~/.config/cmdclip/commands.json
```

The application creates the directory automatically when it is needed.

No external database or cloud service is required.

## Architecture

Command Clipboard uses a small layered architecture:

```text
CLI
 ↓
Service
 ↓
Repository Interface
 ↓
JSON Repository
 ↓
commands.json
```

### CLI

Responsible for parsing user input and displaying results.

### Service

Contains application logic such as finding commands and retrieving commands by ID.

### Repository

Defines the persistence contract without coupling the application to a specific storage implementation.

### JSON Repository

Implements the repository using a local JSON file.

This keeps the application simple while leaving room for another storage implementation if the project eventually needs one.

## Project Structure

```text
command-clipboard/
├── cmd/
│   └── main.go
├── internal/
│   ├── cli/
│   │   └── cli.go
│   ├── command/
│   │   ├── command.go
│   │   └── command_test.go
│   ├── repository/
│   │   └── command.go
│   ├── service/
│   │   ├── command.go
│   │   └── command_test.go
│   └── storage/
│       ├── json_command_repository.go
│       └── json_command_repository_test.go
├── go.mod
└── README.md
```

## Development

Clone the repository:

```bash
git clone https://github.com/Samuel1604/command-clipboard.git
cd command-clipboard
```

Run the tests:

```bash
go test ./...
```

Build the CLI:

```bash
go build -o cmdclip ./cmd
```

Run it:

```bash
./cmdclip list
```

## Current Scope

Command Clipboard is intentionally small.

The goal of the project is to solve a real developer problem with a simple local tool while keeping the code easy to understand, test, and extend.

The current MVP focuses on:

```text
save
list
find
get
delete
```

## License

MIT
