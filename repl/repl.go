package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/racg0092/rhombifer/lexer"
	"github.com/racg0092/rhombifer/tokens"
)

const PROMPT = ">> "

func Start(i io.Reader, o io.Writer) {
	scanner := bufio.NewScanner(i)

	for {
		fmt.Fprintf(o, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != tokens.EOF; tok = l.NextToken() {
			fmt.Fprintf(o, "%+v\n", tok)
		}
	}
}
