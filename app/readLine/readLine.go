package readline

import (
	"fmt"
	"io"
	"unicode/utf8"
)

type ReadLine struct {
	Stdin  io.ReadCloser
	Stdout io.WriteCloser
	Stderr io.WriteCloser
	buf    []rune
	cursor int
}

func NewReadLine(Stdin io.ReadCloser, Stdout io.WriteCloser, Stderr io.WriteCloser) ReadLine {
	return ReadLine{Stdin: Stdin, Stdout: Stdout, Stderr: Stderr, buf: make([]rune, 4096), cursor: 2}
}

func (rl *ReadLine) Read() ([]byte, error) {
	byteBuf := make([]byte, 1024)
	for {
		n, _ := rl.Stdin.Read(byteBuf)
		runeBuf := make([]rune, 0)
		for n > 0 {
			r, size := utf8.DecodeRune(byteBuf[:n])
			runeBuf = append(runeBuf, r)
			copy(byteBuf, byteBuf[size:n])
			n -= size
		}
		// fmt.Println("a")
		rl.handleInput(runeBuf)

	}

}

func (rl *ReadLine) handleInput(runeBuf []rune) ([]byte, error) {
	for _, run := range runeBuf {
		switch run {
		case 0x03:
			return make([]byte, 1), io.EOF
		// case :
		default:
			// fmt.Println("b")
			rl.buf = append(rl.buf, run)
			rl.cursor++
			rl.redrawLine()
		}

	}
	// rl.echo(rl.buf)
	return make([]byte, 2), nil
}

func (rl *ReadLine) redrawLine() {
	fmt.Fprint(rl.Stdout, "\r")
	line := "$ " + string(rl.buf)
	fmt.Print(line)
	// fmt.Print("\x1b[K")
	// wanted := len("$ ") + rl.cursor
	// current := len(line)
	// diff := current - wanted
	// // fmt.Println("c")
	// if diff > 0 {
	// 	// Move left diff
	// 	fmt.Printf("\x1b[%dD", diff)
	// }
}
