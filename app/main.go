package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" and "os" imports in stage 1 (feel free to remove this!)
var _ = fmt.Fprint
var _ = os.Stdout

func main() {
	// TODO: Uncomment the code below to pass the first stage
	for {
		fmt.Fprint(os.Stdout, "$ ")
		b, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
		b = b[:len(b)-1]
		if err != nil {
			log.Fatal("error reading input")
		}
		cmd := string(b)
		splittedCmd := strings.Split(cmd, " ")
		mainCmd := splittedCmd[0]
		// fmt.Print(mainCmd)
		switch mainCmd {
		case "exit":
			subCmd := "0"
			if len(cmd) == 2 {
				subCmd = splittedCmd[1]
			}
			v, err := strconv.Atoi(subCmd)
			if err != nil {
				fmt.Fprint(os.Stdout, "invalid code")
			}
			os.Exit(v)

		case "echo":
			for _, val := range splittedCmd[1:] {
				fmt.Print(val + " ")
			}
			fmt.Println()

		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", mainCmd)
		}

	}
}
