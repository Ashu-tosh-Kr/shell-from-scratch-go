package readline

import (
	"fmt"
	"io"
)

type ReadLine struct {
	Stdin  io.ReadCloser
	Stdout io.WriteCloser
	Stderr io.WriteCloser
}

func NewReadLine(Stdin io.ReadCloser, Stdout io.WriteCloser, Stderr io.WriteCloser) ReadLine {
	return ReadLine{Stdin: Stdin, Stdout: Stdout, Stderr: Stderr}
}

func (rl *ReadLine) Read() ([]byte, error) {
	buf := make([]byte, 1024)
	read := 0
	for {
		n, _ := rl.Stdin.Read(buf[read:])
		for i := read; i < read+n; i++ {
			if buf[i] == 0x03 { // Ctrl-C
				return make([]byte, 1), io.EOF

			}
		}
		rl.Echo(buf[read : read+n])

		read += n
	}

}

func (rl *ReadLine) Echo(b []byte) {
	fmt.Fprint(rl.Stdout, string(b))
}

func parseInput() {

}
