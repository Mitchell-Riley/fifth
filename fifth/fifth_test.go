package fifth

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestInterpret(t *testing.T) {
	file, err := os.Open("test.forth")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	l, err := NewLexer(file)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	l.lex()

	i := NewInterpreter(l)
	i.Eval()

	fmt.Println(i.Lexer.tokens)
	fmt.Println(i.DataStack)
}

func TestREPL(t *testing.T) {
	t.Skip()
	reader := bufio.NewReader(os.Stdin)
	l, err := NewLexer(reader)
	if err != nil {
		log.Fatal(err)
	}
	i := NewInterpreter(l)
	reader.ReadString('\n')
	i.Eval()
}
