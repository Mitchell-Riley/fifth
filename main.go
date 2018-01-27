package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/Mitchell-Riley/fith/fifth"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	c := fifth.NewCompiler(nil)
	for scanner.Scan() {
		// input := strings.Split(scanner.Text(), " ")
		l, err := fifth.NewLexer(strings.NewReader(scanner.Text()))
		if err != nil {
			log.Fatal(err)
		}
		c.Lexer = l
		c.Eval()
	}
}
