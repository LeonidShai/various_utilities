package interpreter

import (
	"strconv"
	"strings"
	"unicode"
)

type parser struct {
	commands []Interpreter
	opers    []string
	vals     []int
}

func NewParser() parser {
	return parser{
		commands: make([]Interpreter, 0),
		opers:    make([]string, 0),
		vals:     make([]int, 0),
	}
}

func (p *parser) Parse(expression string) {
	p.commands = make([]Interpreter, 0)
	p.opers = make([]string, 0)
	p.vals = make([]int, 0)

	tokens := strings.Split(expression, " ")
	for _, token := range tokens {
		if isInt(token) {
			val, _ := strconv.Atoi(token)
			p.vals = append(p.vals, val)
		}

		if token == "+" || token == "-" || token == "*" {
			p.opers = append(p.opers, token)
		}
	}

	for _, oper := range p.opers {
		switch oper {
		case "+":
			p.commands = append(p.commands, &AddValue{})
		case "-":
			p.commands = append(p.commands, &SubValue{})
		}
	}
}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func (p *parser) Result() int {
	res := p.commands[0].Interpret(p.vals[0], p.vals[1])
	if len(p.commands) > 1 {
		for i := 1; i < len(p.commands); i++ {
			res = p.commands[i].Interpret(res, p.vals[i+1])
		}
	}

	return res
}
