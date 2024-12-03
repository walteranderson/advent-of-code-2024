package main

import (
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Provide input")
	}
	filename := os.Args[1]
	input, err := getFileContent(filename)
	if err != nil {
		log.Fatal(err)
	}
	l := NewLexer(input)
	for l.shouldContinue() {
		// log.Printf("pos=%d, ch=%s\n", l.position, string(l.ch))
		switch l.ch {
		case 'm':
			if l.enabled {
				l.readMul()
			} else {
				l.readChar()
			}
		case 'd':
			l.readConditional()
		default:
			l.readChar()
		}
	}
	if len(l.muls) == 0 {
		log.Fatal("no multiplication instructions")
	}
	result := 0
	for _, mul := range l.muls {
		result += mul.left * mul.right
	}
	log.Printf("Result: %d\n", result)
}

type Mul struct {
	left  int
	right int
}

type Lexer struct {
	enabled      bool
	input        string
	position     int
	readPosition int
	ch           byte
	muls         []Mul
}

func NewLexer(input string) *Lexer {
	l := &Lexer{enabled: true, input: input, muls: make([]Mul, 0)}
	l.readChar()
	return l
}

func (l *Lexer) readConditional() {
	// so far: d
	l.readChar()
	if l.ch != 'o' {
		return
	}
	l.readChar()
	if l.ch == 'n' {
		l.readDisable()
	} else if l.ch == '(' {
		l.readEnable()
	}
}

func (l *Lexer) readDisable() {
	// so far: don
	l.readChar()
	if l.ch != '\'' {
		return
	}
	l.readChar()
	if l.ch != 't' {
		return
	}
	l.readChar()
	if l.ch != '(' {
		return
	}
	l.readChar()
	if l.ch != ')' {
		return
	}
	l.enabled = false
}

func (l *Lexer) readEnable() {
	// so far: do(
	l.readChar()
	if l.ch != ')' {
		return
	}
	l.enabled = true
}

func (l *Lexer) readMul() {
	// so far: m
	l.readChar()
	if l.ch != 'u' {
		return
	}
	l.readChar()
	if l.ch != 'l' {
		return
	}
	l.readChar()
	if l.ch != '(' {
		return
	}
	l.readChar()
	left, err := l.getParam()
	if err != nil {
		return
	}
	if l.ch != ',' {
		return
	}
	l.readChar()
	right, err := l.getParam()
	if err != nil {
		return
	}
	if l.ch != ')' {
		return
	}
	if left > 999 || right > 999 {
		return
	}
	log.Printf("Adding mul instruction: %d * %d\n", left, right)
	mul := Mul{left: left, right: right}
	l.muls = append(l.muls, mul)
}

func (l *Lexer) getParam() (int, error) {
	result := 0
	for l.isDigit() {
		i, err := l.byteToInt()
		if err != nil {
			return -1, err
		}
		result = result*10 + i
		l.readChar()
	}
	return result, nil
}

func (l *Lexer) isDigit() bool {
	return l.ch >= '0' && l.ch <= '9'
}

func (l *Lexer) byteToInt() (int, error) {
	return strconv.Atoi(string(l.ch))
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) shouldContinue() bool {
	return l.ch != 0
}

func getFileContent(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
