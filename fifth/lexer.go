package fifth

import (
	"errors"
	"io"
	"io/ioutil"
	"unicode"
)

type lexer struct {
	input    []byte
	tokens   []string
	start    int
	position int
}

func NewLexer(file interface{}) (*lexer, error) {
	tokens := []string{}

	switch file.(type) {
	case io.Reader:
		data, err := ioutil.ReadAll(file.(io.Reader))
		if err != nil {
			return nil, err
		}

		return &lexer{data, tokens, 0, 0}, nil
	case []byte:
		return &lexer{file.([]byte), tokens, 0, 0}, nil
	}
	return nil, errors.New("unknown argument type")
}

func (l *lexer) next() byte {
	if l.position == len(l.input) {
		return 0
	}

	next := l.input[l.position:][0]
	l.position++
	return next
}

func (l *lexer) peek() byte {
	s := l.next()
	if l.position == len(l.input) {
		return 0
	}

	l.position--
	return s
}

// ignore the current character
func (l *lexer) ignore() {
	l.start = l.position
}

func (l *lexer) tokenAppend() {
	l.tokens = append(l.tokens, string(l.input[l.start:l.position]))
	l.start = l.position
}

func (l *lexer) lex() {
	for tok := l.next(); tok != 0; tok = l.next() {
		switch v := tok; {
		case v == 0:
			return
		case v == ' ', v == '\r', v == '\n':
			l.ignore()
		case isAlphanumeric(rune(tok)):
			for isAlphanumeric(rune(l.peek())) {
				l.next()
			}
			l.tokenAppend()
		default:
			l.tokenAppend()
			// panic(fmt.Sprintf("unknown token %s", string(tok)))
		}
	}
}

func isAlphanumeric(r rune) bool {
	return unicode.IsDigit(r) || unicode.IsLetter(r)
}
