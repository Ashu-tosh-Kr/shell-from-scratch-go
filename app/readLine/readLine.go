package readline

import (
	"fmt"
	"io"
	"log"
	"os"
	"unicode/utf8"

	"github.com/charmbracelet/x/term"
)

type ReadLine struct {
	Stdin  io.ReadCloser
	Stdout io.WriteCloser
	Stderr io.WriteCloser
	buf    []rune
	cursor int
	histF  *os.File
}

func NewReadLine(Stdin io.ReadCloser, Stdout io.WriteCloser, Stderr io.WriteCloser) ReadLine {
	historyFile, err := os.OpenFile("history.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer historyFile.Close()
	return ReadLine{Stdin: Stdin, Stdout: Stdout, Stderr: Stderr, buf: make([]rune, 0), cursor: 0, histF: historyFile}
}

func (rl *ReadLine) Read() ([]byte, error) {
	// Get current terminal state
	oldState, err := term.MakeRaw(os.Stdin.Fd())
	if err != nil {
		panic(err)
	}
	// Restore state on exit
	defer term.Restore(os.Stdin.Fd(), oldState)

	byteBuf := make([]byte, 1024)
	fmt.Fprint(os.Stdout, "$ ")
	for {
		n, _ := rl.Stdin.Read(byteBuf)
		runeBuf := make([]rune, 0)
		for n > 0 {
			r, size := utf8.DecodeRune(byteBuf[:n])
			runeBuf = append(runeBuf, r)
			copy(byteBuf, byteBuf[size:n])
			n -= size
		}
		done := rl.handleInput(runeBuf)
		if done {
			collector := make([]byte, 32)
			out := make([]byte, 0)
			for _, run := range rl.buf {
				utf8.EncodeRune(collector, run)
				out = append(out, byte(run))
			}
			rl.cursor = 0
			rl.buf = rl.buf[:0]
			rl.histF.Write(out)
			rl.histF.Write([]byte{'\n'})
			return out, nil
		}

	}

}

func (rl *ReadLine) handleInput(runeBuf []rune) bool {
	i := 0
	for i < len(runeBuf) {

		switch runeBuf[i] {
		case 0x03: // CTRL + C
			fmt.Fprint(rl.Stdout, "\r\n")
			rl.buf = rl.buf[:0]
			rl.redrawLine()
			i++
		case 13: // /r/n
			fmt.Fprint(rl.Stdout, "\r\n")
			return true
		case 127: // backspace
			if rl.cursor == 0 {
				return false
			}
			rl.cursor--
			if rl.cursor == len(rl.buf)-1 {
				rl.buf = rl.buf[:rl.cursor]
			} else {
				rl.buf = append(rl.buf[:rl.cursor], rl.buf[rl.cursor+1:]...)
			}
			rl.redrawLine()
			i++
		case '\x1b':
			if i+2 < len(runeBuf) && runeBuf[i+1] == '[' {
				switch runeBuf[i+2] {
				case 'A':
					break
				case 'B':
					break
				case 'C':
					rl.cursor = min(rl.cursor+1, len(rl.buf))
					rl.redrawLine()
				case 'D':
					rl.cursor = max(rl.cursor-1, 0)
					rl.redrawLine()
				}
				i += 3
			}
		default:
			tmp := make([]rune, rl.cursor)
			copy(tmp, rl.buf[:rl.cursor])
			tmp = append(tmp, runeBuf[i])
			if rl.cursor == len(rl.buf) {
				rl.buf = tmp
			} else {
				rl.buf = append(tmp, rl.buf[rl.cursor:]...)

			}
			rl.cursor++
			rl.redrawLine()
			i++
		}

	}
	return false
}

func (rl *ReadLine) redrawLine() {
	fmt.Fprint(rl.Stdout, "\r")
	line := "$ " + string(rl.buf)
	fmt.Print(line)
	fmt.Print("\x1b[K")
	wanted := len("$ ") + rl.cursor
	current := len(line)
	diff := current - wanted
	if diff > 0 {
		// Move left diff
		fmt.Printf("\x1b[%dD", diff)
	}
}
