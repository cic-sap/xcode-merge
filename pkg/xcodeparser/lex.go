package xcodeparser

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//go:generate go run golang.org/x/tools/cmd/goyacc -l -o parser.go parser.y

// Parse parses the input and returs the result.
func Parse(input []byte) (map[string]interface{}, error) {
	arr := strings.Split(strings.TrimSpace(string(input)), "\n")
	if strings.HasPrefix(arr[0], "//") {
		arr = arr[1:]
	}
	input = []byte(strings.Join(arr, "\n"))
	l := newLex(input)
	println("lex", l, l.err)
	v := yyParse(l)
	if l.err != nil {
		fmt.Println("lex error:", v, l.pos, l.err)
	}
	return l.result, l.err
}

type lex struct {
	input  []byte
	pos    int
	result map[string]interface{}
	err    error
}

func newLex(input []byte) *lex {
	return &lex{
		input: input,
	}
}

// Lex satisfies yyLexer.
func (l *lex) Lex(lval *yySymType) int {
	return l.scanNormal(lval)
}

func (l *lex) skipComment(lval *yySymType) {
	for b := l.next(); b != 0; b = l.next() {
		if b == '*' {
			b = l.next()
			if b == '/' {
				return
			} else {
				l.backup()
			}
		}
	}
}
func (l *lex) scanNormal(lval *yySymType) int {
	for b := l.next(); b != 0; b = l.next() {
		switch {
		case unicode.IsSpace(rune(b)):
			continue
		case b == '/':
			b = l.next()
			if b == '*' {
				l.skipComment(lval)
				continue
			}

		case b == '"':

			return l.scanString(lval)
		case unicode.IsDigit(rune(b)) || b == '+' || b == '-':
			l.backup()
			return l.scanNum(lval)
		case !l.isEnd(b):
			l.backup()
			return l.scanLiteral(lval)
		default:
			return int(b)
		}
	}
	return 0
}

var escape = map[byte]byte{
	'"':  '"',
	'\\': '\\',
	'/':  '/',
	'b':  '\b',
	'f':  '\f',
	'n':  '\n',
	'r':  '\r',
	't':  '\t',
}

func (l *lex) scanString(lval *yySymType) int {
	println("start scanString")
	buf := bytes.NewBuffer(nil)
	for b := l.next(); b != 0; b = l.next() {
		switch b {
		case '\\':
			b2 := escape[l.next()]
			if b2 == 0 {
				return LexError
			}
			buf.WriteByte(b2)
		case '"':
			lval.val = buf.String()
			fmt.Sprintf("end scanString:%s", lval.val)
			return String
		default:
			buf.WriteByte(b)
		}
	}
	return LexError
}

func (l *lex) scanNum(lval *yySymType) int {
	buf := bytes.NewBuffer(nil)
	for {
		b := l.next()
		switch {
		case unicode.IsDigit(rune(b)):
			buf.WriteByte(b)
		case strings.IndexByte(".+-eE", b) != -1:
			buf.WriteByte(b)
		default:
			l.backup()
			fmt.Println("is not number", string(b))
			isEnd := l.isEnd(b)
			val, err := strconv.ParseFloat(buf.String(), 64)
			if !isEnd || err != nil {
				fmt.Println("get num error", err, buf.String())
				for {
					b := l.next()
					if b == '=' || b == ';' || unicode.IsSpace(rune(b)) {
						l.backup()
						break
					} else {
						buf.WriteByte(b)
					}
				}
				lval.val = buf.String()
				fmt.Println("get raw string", buf.String())
				return Literal
			}
			lval.val = val
			fmt.Println("get Number", buf.String(), val)
			return Number
		}
	}
}

func (l *lex) isEnd(b byte) bool {
	return b == '=' || b == ',' ||
		b == '(' ||
		b == ')' ||
		b == ';' ||
		b == '}' ||
		unicode.IsSpace(rune(b)) ||
		b == '{'

}
func (l *lex) scanLiteral(lval *yySymType) int {
	buf := bytes.NewBuffer(nil)
	println("start scanLiteral")
	for {
		b := l.next()
		isEnd := l.isEnd(b)
		switch {
		case !isEnd:
			buf.WriteByte(b)
		default:
			l.backup()
			//val, ok := literal[buf.String()]
			//if !ok {
			//	return LexError
			//}
			fmt.Println("literal:", buf.String())
			lval.val = buf.String()
			return Literal
		}
	}
}

func (l *lex) backup() {
	if l.pos == -1 {
		return
	}
	l.pos--
}

func (l *lex) next() byte {
	if l.pos >= len(l.input) || l.pos == -1 {
		l.pos = -1
		return 0
	}
	l.pos++
	return l.input[l.pos-1]
}

// Error satisfies yyLexer.
func (l *lex) Error(s string) {
	l.err = errors.New(s)
}
