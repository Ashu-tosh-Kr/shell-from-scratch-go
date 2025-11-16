# Shell Implementation in Go

A simple shell implementation written in Go that's capable of interpreting shell commands, running external programs and builtin commands like cd, pwd, echo and more.

## Terminal vs Shell: What's the Difference?

**Terminal** = The window/app that displays input and output  
**Shell** = The program that interprets commands(input) and runs programs like ls, cat, cd etc to get ouput

**Example:**
1. You open **Terminal.app** (the terminal)
2. **Terminal.app** runs the shell program and prints it's output `$` prompt from **bash/zsh** (the shell)
3. You type `echo hello`
4. **Terminal.app** sends this to **The shell** which understands and executes the command and sends it's output to **Terminal.app**
5. **The terminal** displays the output "hello"

**This Go program is a shell** - it reads your commands, parses them, and executes them. You run it inside a terminal for display.

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
