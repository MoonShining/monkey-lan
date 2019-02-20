package repl

import (
	"bufio"
	"fmt"
	"github.com/MoonShining/monkey/lexer"
	"github.com/MoonShining/monkey/token"
	"io"
)

const HEADER = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(HEADER)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for t := l.NextToken(); t.Type != token.EOF; t = l.NextToken() {
			fmt.Printf("%+v\n", t)
		}
	}
}
