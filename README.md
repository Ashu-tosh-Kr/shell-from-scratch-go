# Shell Implementation in Go

A POSIX compliant shell implementation written in Go that's capable of interpreting shell commands, running external programs and builtin commands like cd, pwd, echo and more.

## Capabilities

### Built-in Commands
- **`echo`** - Display text to stdout
- **`exit`** - Exit the shell with optional exit code
- **`pwd`** - Print current working directory
- **`cd`** - Change directory (supports `~` for home directory)
- **`type`** - Display whether command is builtin or external program
- **`cat`** - Display file contents (also supports stdin)
- **`history`** - Show command history with optional number limit

### Shell Features
- **Command Execution** - Run external programs from PATH
- **Pipes** - Chain commands using `|` operator
- **Output Redirection** - Redirect stdout and stderr using `>`, `>>`, `2>`, `2>>`
- **Command History** - Automatically tracks all executed commands in `history.txt`

### Highlights
- **Tokenization & Parsing** - I have implemented a mini interpreter that tokenizes -> parses -> evaluates the commands that are passed by the user
- **AST-based Evaluation** - My implementation of parser uses pratt parsing create an AST
- **PATH Resolution** - Automatic lookup of external programs in system PATH

## How to Run

1. Ensure you have Go installed locally (version 1.24 or higher)
2. Run the following command to start the shell:

```sh
go run app/main.go
```

The main program is implemented in `app/main.go`.
