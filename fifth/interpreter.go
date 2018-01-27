package fifth

import (
	"log"
	"strconv"
)

const (
	TRUE  = -1 // all bits set
	FALSE = 0  // alls bits not set
)

type interpreter struct {
	Lexer       *lexer
	DataStack   *stack
	ReturnStack *stack
	position    int

	Core        map[string]func()
	userDefined map[string]func()
}

func NewInterpreter(l *lexer) *interpreter {
	i := &interpreter{
		Lexer:       l,
		DataStack:   new(stack),
		ReturnStack: new(stack),
		userDefined: make(map[string]func()),
	}
	i.generateCore()
	return i
}

func (i *interpreter) next() string {
	if i.position == len(i.Lexer.tokens) {
		return ""
	}

	next := i.Lexer.tokens[i.position:][0]
	i.position++
	return next
}

// evaluate the tokens
func (i *interpreter) Eval() {
	for tok := i.next(); tok != ""; tok = i.next() {
		switch tok {
		case "(":
			i.evalComment()
		// user-defined functions
		case ":":
			i.evalUserDefined()
		case "true":
			i.DataStack.Push(TRUE)
		case "false":
			i.DataStack.Push(FALSE)
		default:
			c, err := strconv.Atoi(tok)
			if err == nil {
				i.DataStack.Push(c)
				break
			}

			if k, ok := i.Core[tok]; ok {
				k()
			} else if k, ok := i.userDefined[tok]; ok {
				k()
			} else {
				log.Fatal("unknown token ", tok)
			}
		}
	}
}

// maybe this should be a part of the lexer instead?
// I first discovered the balanced parentheses checker algorithm in
// Data Structures: Abstraction and Design Using Java (2nd Edition)
// ISBN: 978-0470128701
func (i *interpreter) evalComment() {
	s := new(stack)
	s.Push('(')

	for t := i.next(); t != ""; t = i.next() {
		switch t {
		case "(":
			s.Push('(')
		case ")":
			s.Pop()
			// if the stack is empty, then we can stop looping because
			// the parentheses are balanced. if it still has elements in
			// it, it might be a nested comment, so we should continue
			// until we reach the end. if we reach the end and the stack
			// is still full, then it doesn't matter that we skipped
			// past a bunch of tokens because the comment parsing will
			// have failed anyway
			if s.isEmpty() {
				return
			}
		}
	}

	if !s.isEmpty() {
		log.Fatal("Unbalanced comment parentheses")
	}
}

func (i *interpreter) evalUserDefined() {
	name := i.next()
	imp := []string{}

	for t := i.next(); t != ""; t = i.next() {
		switch t {
		case "(":
			i.evalComment()
		case ";":
			i.userDefined[name] = func() {
				for _, word := range imp {
					// repition here wih Eval, should move all of this
					// into it's own function
					c, err := strconv.Atoi(word)
					if err == nil {
						i.DataStack.Push(c)
						break
					}

					if word == "true" {
						i.DataStack.Push(TRUE)
						break
					}
					if word == "false" {
						i.DataStack.Push(FALSE)
						break
					}

					if k, ok := i.Core[word]; ok {
						k()
					} else if k, ok := i.userDefined[word]; ok {
						k()
					} else {
						log.Fatal("unknown token ", word)
					}
				}
			}
			return
		default:
			imp = append(imp, t)
		}
	}
	log.Fatal("Missing ';' from word definition")
}
