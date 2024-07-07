package parser

import (
	"errors"
	"strings"
)

type ActionHandler func(param *Param) error
type Action struct {
	parser *Parser
	noun   string
}

type Parser struct {
	handlers map[string]map[string]ActionHandler
}

func NewParser() *Parser {
	return &Parser{
		handlers: map[string]map[string]ActionHandler{},
	}
}

func (p *Parser) EmptyNoun(handler ActionHandler) {
	p.handlers["NULL"] = map[string]ActionHandler{
		"NULL": handler,
	}
}

func (p *Parser) Noun(noun string) *Action {
	return &Action{
		p,
		noun,
	}
}

func (act *Action) Verb(verb string, handler ActionHandler) *Action {
	if _, ok := act.parser.handlers[act.noun]; !ok {
		act.parser.handlers[act.noun] = map[string]ActionHandler{}
	}

	act.parser.handlers[act.noun][verb] = handler
	return act
}

func (p *Parser) sanitize(args []string) []string {
	newArgs := []string{}
	i := 1
	for i < len(args) {
		if strings.HasPrefix(args[i], "-") {
			i += 2
		} else {
			newArgs = append(newArgs, args[i])
			i++
		}
	}

	return newArgs
}

func (p *Parser) parseEmptyNoun(params *Param) error {
	handler, ok := p.handlers["NULL"]
	if ok {
		return handler["NULL"](params)
	} else {
		return errors.New("parser: unknown command")
	}
}

func (p *Parser) parseNoun(params *Param) error {
	noun, _ := params.Next()
	_, ok := p.handlers[noun]
	if !ok {
		return errors.New("parser: unknown noun")
	}

	verb, ok := params.Next()
	if !ok {
		return errors.New("parser: insufficient arguments")
	}

	handler, ok := p.handlers[noun][verb]
	if !ok {
		return errors.New("parser: unknown verb")
	}

	err := handler(params)
	if err != nil {
		return err
	}

	return nil
}

func (p *Parser) Parse(args []string) error {
	params := NewParam(p.sanitize(args))

	var err error
	if !params.HasNext() {
		err = p.parseEmptyNoun(params)
	} else {
		err = p.parseNoun(params)
	}

	if err != nil {
		return err
	}

	return nil
}
