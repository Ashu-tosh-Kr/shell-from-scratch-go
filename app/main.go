package main

import (
	"fmt"
	"os"
	"strconv"
)

// Ensures gofmt doesn't remove the "fmt" and "os" imports in stage 1 (feel free to remove this!)
var _ = fmt.Fprint
var _ = os.Stdout

func main() {
	// TODO: Uncomment the code below to pass the first stage
	for {
		fmt.Fprint(os.Stdout, "$ ")
		var cmd string
		var subcmd string
		fmt.Scanln(&cmd, &subcmd)
		if cmd == "exit" {
			v, err := strconv.Atoi(subcmd)
			if err != nil {
				fmt.Fprint(os.Stdout, "invalid code")
			}
			os.Exit(v)
		}
		fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd)

	}
}
