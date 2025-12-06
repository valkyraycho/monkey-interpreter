package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/valkyraycho/interpreter/evaluator"
	"github.com/valkyraycho/interpreter/lexer"
	"github.com/valkyraycho/interpreter/object"
	"github.com/valkyraycho/interpreter/parser"
)

const PROMPT = ">> "

func Start(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	env := object.NewEnvironment()

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(w, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(w, evaluated.Inspect())
			io.WriteString(w, "\n")
		}
	}

}

func printParserErrors(w io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(w, "\t"+msg+"\n")
	}
}
