package fifth

import (
	"fmt"
	"log"
)

// s is the stack that the core words will act upon
func (i *interpreter) generateCore() {
	d := i.DataStack
	r := i.ReturnStack
	i.Core = map[string]func(){
		"+": func() {
			d.Push(d.Pop().data + d.Pop().data)
		},
		"-": func() {
			d.Push(d.Pop().data - d.Pop().data)
		},
		"*": func() {
			d.Push(d.Pop().data * d.Pop().data)
		},
		".": func() {
			fmt.Println(d.Pop())
		},
		"0<": func() {
			if d.Pop().data < 0 {
				d.Push(TRUE)
			} else {
				d.Push(FALSE)
			}
		},
		"0=": func() {
			if d.Pop().data == 0 {
				d.Push(TRUE)
			} else {
				d.Push(FALSE)
			}
		},
		"1+": func() {
			d.Push(d.Pop().data + 1)
		},
		"1-": func() {
			d.Push(d.Pop().data - 1)
		},
		"2DROP": func() {
			d.Pop()
			d.Pop()
		},
		"2DUP": func() {
			x := d.Pop()
			y := d.Pop()

			d.Push(y)
			d.Push(x)
			d.Push(y)
			d.Push(x)
		},
		"2SWAP": func() {
			w := d.Pop()
			x := d.Pop()
			y := d.Pop()
			z := d.Pop()

			d.Push(x)
			d.Push(w)
			d.Push(z)
			d.Push(y)
		},
		"<": func() {
			n := d.Pop()
			m := d.Pop()
			if m.data < n.data {
				d.Push(TRUE)
			} else {
				d.Push(FALSE)
			}
		},
		"=": func() {
			d.Push(d.Pop().data == d.Pop().data)
		},
		">": func() {
			n := d.Pop()
			m := d.Pop()
			if m.data > n.data {
				d.Push(TRUE)
			} else {
				d.Push(FALSE)
			}
		},
		">R": func() {
			r.Push(d.Pop())
		},
		// duplicate x if it is non-zero
		"?DUP": func() {
			x := d.Pop()

			d.Push(x)
			if x.data == 0 {
				d.Push(x)
			}
		},
		"DEPTH": func() {
			d.Push(d.size())
		},
		"DROP": func() {
			d.Pop()
		},
		"R>": func() {
			d.Push(r.Pop())
		},
		"ROT": func() {
			x := d.Pop()
			y := d.Pop()
			z := d.Pop()
			d.Push(y)
			d.Push(x)
			d.Push(z)
		},
		"SWAP": func() {
			x := d.Pop()
			y := d.Pop()
			d.Push(x)
			d.Push(y)
		},

		// extended words
		"NIP": func() {
			x := d.Pop()
			d.Pop()
			d.Push(x)
		},
		"TUCK": func() {
			x := d.Pop()
			y := d.Pop()
			d.Push(x)
			d.Push(y)
			d.Push(x)
		},
		"ROLL": func() {
			u := d.Pop()
			if d.size() < u.data+2 {
				log.Fatal("ambiguous condition!")
			}

		},
	}
}
